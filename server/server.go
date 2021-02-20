package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

type Handler func(c echo.Context, req protoreflect.ProtoMessage) (protoreflect.ProtoMessage, error)
type HandlerTransformer func(meth *descriptorpb.MethodDescriptorProto, service *descriptorpb.ServiceDescriptorProto, d *descriptorpb.FileDescriptorProto, h Handler) Handler

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
