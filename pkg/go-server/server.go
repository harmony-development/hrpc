package goserver

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type VTProtoMessage interface {
	proto.Message
	MarshalVT() ([]byte, error)
	UnmarshalVT([]byte) error
}

type MethodData struct {
	Input  VTProtoMessage
	Output VTProtoMessage
}

type UnaryMethodData[C context.Context] struct {
	MethodData
	Handler func(C, VTProtoMessage) (VTProtoMessage, error)
}

type StreamMethodData[C context.Context] struct {
	MethodData
	Handler func(C, chan VTProtoMessage) (chan VTProtoMessage, error)
}

type MethodHandler[T context.Context] interface {
	Routes() map[string]UnaryMethodData[T]
	StreamRoutes() map[string]StreamMethodData[T]
}

func FromProtoChannel[T VTProtoMessage](in chan VTProtoMessage) chan T {
	res := make(chan T)
	go func() {
		for {
			v := <-in
			res <- v.(T)
		}
	}()

	return res
}

func ToProtoChannel[T VTProtoMessage](in chan T) chan VTProtoMessage {
	res := make(chan VTProtoMessage)
	go func() {
		for {
			v := <-in
			res <- v
		}
	}()

	return res
}
