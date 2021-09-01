package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type HRPCServiceHandler interface {
	Routes() map[string]echo.HandlerFunc
	SetUnaryPre(h HandlerTransformer)
}

type HRPCServer struct {
	handlers []HRPCServiceHandler
	serveMux *http.ServeMux
}

type (
	RawHandler         func(c context.Context, w http.ResponseWriter, r *http.Request) error
	Handler            func(c context.Context, req protoreflect.ProtoMessage) (protoreflect.ProtoMessage, error)
	HandlerTransformer func(meth *descriptorpb.MethodDescriptorProto, service *descriptorpb.ServiceDescriptorProto, d *descriptorpb.FileDescriptorProto, h Handler) Handler
)

func ChainHandlerTransformers(funs ...HandlerTransformer) HandlerTransformer {
	switch len(funs) {
	case 0:
		return nil
	case 1:
		return funs[0]
	default:
		return func(meth *descriptorpb.MethodDescriptorProto, service *descriptorpb.ServiceDescriptorProto, d *descriptorpb.FileDescriptorProto, h Handler) Handler {
			in := h
			for i := len(funs) - 1; i >= 0; i-- {
				item := funs[i]
				in = item(meth, service, d, in)
			}
			return in
		}
	}
}

func NewHRPCServer(e *echo.Echo, items ...HRPCServiceHandler) *HRPCServer {
	hentaiRPCServer := &HRPCServer{
		handlers: items,
		serveMux: http.NewServeMux(),
	}
	for _, item := range hentaiRPCServer.handlers {
		for handler, route := range item.Routes() {
			e.Any(handler, route)
		}
	}
	return hentaiRPCServer
}

func (h *HRPCServer) SetUnaryPre(han HandlerTransformer) {
	for _, item := range h.handlers {
		item.SetUnaryPre(han)
	}
}

func ScanProto(src interface{}, m protoreflect.ProtoMessage) error {
	if src == nil {
		return nil
	}
	if b, ok := src.([]byte); ok {
		if err := proto.Unmarshal(b, m); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("unexpected type %T", src)
}

// NewUnaryHandler creates a new raw HTTP handler that deserializes a given message type
func NewUnaryHandler(messageType proto.Message, unaryHandler Handler) RawHandler {
	return func(c context.Context, w http.ResponseWriter, r *http.Request) error {
		newMessage := proto.Clone(messageType)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		if err := proto.Unmarshal(body, newMessage); err != nil {
			return err
		}
		result, err := unaryHandler(c, newMessage)
		if err != nil {
			return err
		}
		var response []byte
		switch r.Header.Get("Content-Type") {
		case "application/hrpc":
			response, err = proto.Marshal(result)
		default:
			response, err = protojson.Marshal(result)
		}
		if err != nil {
			return err
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return nil
	}
}
