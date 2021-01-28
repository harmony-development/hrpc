package jiti

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"
import "github.com/gorilla/websocket"
import "google.golang.org/protobuf/types/descriptorpb"

var Test *descriptorpb.FileDescriptorProto = new(descriptorpb.FileDescriptorProto)

func init() {
	data := []byte("\n\ntest.proto\x12\x04test\"\x16\n\x04Ping\x12\x0e\n\x02mu\x18\x01 \x01(\tR\x02mu\"\x16\n\x04Pong\x12\x0e\n\x02mu\x18\x01 \x01(\tR\x02mu2H\n\x02Mu\x12\x1c\n\x02Mu\x12\n.test.Ping\x1a\n.test.Pong\x12$\n\x06MuMute\x12\n.test.Ping\x1a\n.test.Pong(\x010\x01B\vZ\thrpc/jitiJ\xf1\x02\n\x06\x12\x04\x00\x00\x10\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\b\n\x01\x02\x12\x03\x02\x00\r\n\b\n\x01\b\x12\x03\x04\x00 \n\t\n\x02\b\v\x12\x03\x04\x00 \n\n\n\x02\x04\x00\x12\x04\x06\x00\b\x01\n\n\n\x03\x04\x00\x01\x12\x03\x06\b\f\n\v\n\x04\x04\x00\x02\x00\x12\x03\a\x04\x12\n\f\n\x05\x04\x00\x02\x00\x05\x12\x03\a\x04\n\n\f\n\x05\x04\x00\x02\x00\x01\x12\x03\a\v\r\n\f\n\x05\x04\x00\x02\x00\x03\x12\x03\a\x10\x11\n\n\n\x02\x04\x01\x12\x04\t\x00\v\x01\n\n\n\x03\x04\x01\x01\x12\x03\t\b\f\n\v\n\x04\x04\x01\x02\x00\x12\x03\n\x04\x12\n\f\n\x05\x04\x01\x02\x00\x05\x12\x03\n\x04\n\n\f\n\x05\x04\x01\x02\x00\x01\x12\x03\n\v\r\n\f\n\x05\x04\x01\x02\x00\x03\x12\x03\n\x10\x11\n\n\n\x02\x06\x00\x12\x04\r\x00\x10\x01\n\n\n\x03\x06\x00\x01\x12\x03\r\b\n\n\v\n\x04\x06\x00\x02\x00\x12\x03\x0e\x04 \n\f\n\x05\x06\x00\x02\x00\x01\x12\x03\x0e\b\n\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03\x0e\v\x0f\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03\x0e\x1a\x1e\n\v\n\x04\x06\x00\x02\x01\x12\x03\x0f\x042\n\f\n\x05\x06\x00\x02\x01\x01\x12\x03\x0f\b\x0e\n\f\n\x05\x06\x00\x02\x01\x05\x12\x03\x0f\x0f\x15\n\f\n\x05\x06\x00\x02\x01\x02\x12\x03\x0f\x16\x1a\n\f\n\x05\x06\x00\x02\x01\x06\x12\x03\x0f%+\n\f\n\x05\x06\x00\x02\x01\x03\x12\x03\x0f,0b\x06proto3")

	err := proto.Unmarshal(data, Test)
	if err != nil {
		panic(err)
	}
}

type MuServer interface {
	Mu(ctx context.Context, r *Ping, headers http.Header) (resp *Pong, err error)

	MuMute(ctx context.Context, in chan *Ping, out chan *Pong, headers http.Header)
}

type MuHandler struct {
	Server       MuServer
	ErrorHandler func(err error, w http.ResponseWriter)
	UnaryPre     func(d *descriptorpb.FileDescriptorProto, f func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error)) func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error)
	upgrader     websocket.Upgrader
}

func NewMuHandler(s MuServer, errHandler func(err error, w http.ResponseWriter)) *MuHandler {
	return &MuHandler{
		Server:       s,
		ErrorHandler: errHandler,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *MuHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/test.Mu/Mu":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(Ping)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.Mu(c, req.(*Ping), headers)
			}

			println(h.UnaryPre)
			if h.UnaryPre != nil {
				invoker = h.UnaryPre(Test, invoker)
			}

			resp, err := invoker(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/test.Mu/MuMute":
		{
			var err error

			in := make(chan *Ping)
			err = nil

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			out := make(chan *Pong)

			ws, err := h.upgrader.Upgrade(w, req, nil)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			go func() {

				msgs := make(chan []byte)

				go func() {
					for {
						_, message, err := ws.ReadMessage()
						if err != nil {
							close(msgs)
							break
						}
						msgs <- message
					}
				}()

				defer ws.WriteMessage(websocket.CloseMessage, []byte{})

				for {
					select {

					case data, ok := <-msgs:
						if !ok {
							return
						}

						item := new(Ping)
						err = proto.Unmarshal(data, item)
						if err != nil {
							close(in)
							close(out)
							return
						}

						in <- item

					case msg, ok := <-out:
						if !ok {
							return
						}

						w, err := ws.NextWriter(websocket.BinaryMessage)
						if err != nil {

							close(in)

							close(out)
							return
						}

						response, err := proto.Marshal(msg)
						if err != nil {

							close(in)

							close(out)
							return
						}

						w.Write(response)
						if err := w.Close(); err != nil {

							close(in)

							close(out)
							return
						}
					}
				}
			}()

			h.Server.MuMute(req.Context(), in, out, req.Header)
		}

	}
}
