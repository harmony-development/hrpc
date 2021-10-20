// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: MPL-2.0

package server

import (
	"context"
	"fmt"
	"time"

	"github.com/fasthttp/websocket"
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
	StreamHandler      func(c context.Context, req chan proto.Message) (chan proto.Message, error)
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

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn, msgType proto.Message, in chan proto.Message) {
	defer func() {
		conn.Close()
	}()
	for {
		_, rawMessage, err := conn.ReadMessage()
		if err != nil {
			break
		}
		msg, err := UnmarshalHRPC(rawMessage, "application/hrpc", msgType)
		if err != nil {
			break
		}
		in <- msg
	}
}

func writer(conn *websocket.Conn, out chan proto.Message) {
	ticker := time.NewTicker(time.Second * 10)
	defer func() {
		conn.Close()
	}()
	for {
		select {
		case message := <-out:
			w, err := conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			raw, err := MarshalHRPC(message, "application/hrpc")
			if err != nil {
				return
			}
			if _, err := w.Write(raw); err != nil {
				return
			}
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// NewStreamingHandler creates a new raw HTTP handler that
func NewStreamingHandler(messageType proto.Message, stream StreamHandler) RawHandler {
	return func(c context.Context, r *fasthttp.Request) ([]byte, error) {
		ctx := c.(*fasthttp.RequestCtx)
		err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
			inChan := make(chan proto.Message)
			go reader(conn, messageType, inChan)
			outChan, err := stream(c, inChan)
			if err != nil {
				return
			}
			go writer(conn, outChan)
		})
		return nil, err
	}
}
