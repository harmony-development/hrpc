package server

import (
	"context"
	"net/http"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

type HRPCServiceHandler interface {
	http.Handler

	Routes() []string
}

type HRPCServer struct {
	handlers []HRPCServiceHandler
	serveMux *http.ServeMux
}

type Handler func(c context.Context, req protoreflect.ProtoMessage, headers http.Header) (protoreflect.ProtoMessage, error)
type HandlerTransformer func(meth *descriptorpb.DescriptorProto, d *descriptorpb.FileDescriptorProto, h Handler) Handler

func ChainHandlerTransformers(funs ...HandlerTransformer) HandlerTransformer {
	switch len(funs) {
	case 0:
		return nil
	case 1:
		return funs[0]
	default:
		// {a, b, c}[0] -> a
		fn := funs[0]
		// {a, b, c}[1:] -> {b, c}
		other := funs[1:]
		// b, c
		// fn = b(a(handler))
		// c
		// fn = c(b(a(handler)))
		// or
		// handler |> a |> b |> c
		for _, fun := range other {
			fn = func(meth *descriptorpb.DescriptorProto, d *descriptorpb.FileDescriptorProto, h Handler) Handler {
				return fn(meth, d, fun(meth, d, h))
			}
		}
		return fn
	}
}

func NewHRPCServer(items ...HRPCServiceHandler) *HRPCServer {
	herpes := &HRPCServer{
		handlers: items,
		serveMux: http.NewServeMux(),
	}
	for _, item := range herpes.handlers {
		for _, route := range item.Routes() {
			herpes.serveMux.Handle(route, item)
		}
	}
	return herpes
}

func (h *HRPCServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.serveMux.ServeHTTP(w, r)
}
