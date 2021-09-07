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

func MarshalHRPC(content proto.Message, contentType string) ([]byte, error) {
	var response []byte
	var err error
	switch contentType {
	case "application/hrpc-json":
		response, err = protojson.Marshal(content)
	default:
		response, err = proto.Marshal(content)
	}
	return response, err
}

func UnmarshalHRPC(content []byte, contentType string, messageType proto.Message) (proto.Message, error) {
	newMessage := proto.Clone(messageType)
	var err error
	switch contentType {
	case "application/hrpc-json":
		if err := protojson.Unmarshal(content, newMessage); err != nil {
			return nil, err
		}
	default:
		if err := proto.Unmarshal(content, newMessage); err != nil {
			return nil, err
		}
	}
	return newMessage, err
}

// NewUnaryHandler creates a new raw HTTP handler that deserializes a given message type
func NewUnaryHandler(messageType proto.Message, unaryHandler Handler) RawHandler {
	return func(c context.Context, r *fasthttp.Request) ([]byte, error) {
		contentType := string(r.Header.Peek("Content-Type"))
		msg, err := UnmarshalHRPC(r.Body(), contentType, messageType)
		if err != nil {
			return nil, err
		}
		result, err := unaryHandler(c, msg)
		if err != nil {
			return nil, err
		}
		return MarshalHRPC(result, contentType)
	}
}
