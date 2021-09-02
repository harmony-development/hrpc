package server

import (
	"context"
	"fmt"

	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type HRPCServiceHandler interface {
	Name() string
	Routes() map[string]RawHandler
}

type (
	RawHandler         func(c context.Context, r *fasthttp.Request) ([]byte, error)
	Handler            func(c context.Context, req proto.Message) (proto.Message, error)
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
	return func(c context.Context, r *fasthttp.Request) ([]byte, error) {
		newMessage := proto.Clone(messageType)
		if err := proto.Unmarshal(r.Body(), newMessage); err != nil {
			return nil, err
		}
		result, err := unaryHandler(c, newMessage)
		if err != nil {
			return nil, err
		}
		var response []byte
		switch string(r.Header.Peek("Content-Type")) {
		case "application/hrpc":
			response, err = proto.Marshal(result)
		default:
			response, err = protojson.Marshal(result)
		}
		if err != nil {
			return nil, err
		}
		return response, nil
	}
}
