package v1

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"

	v1 "github.com/harmony-development/legato/gen/voice/v1"
)

type VoiceServiceServer interface {
	Connect(ctx context.Context, in chan *v1.ClientSignal, out chan *v1.Signal, headers http.Header)
}

type VoiceServiceHandler struct {
	Server       VoiceServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	upgrader     websocket.Upgrader
}

func NewVoiceServiceHandler(s VoiceServiceServer, errHandler func(err error, w http.ResponseWriter)) *VoiceServiceHandler {
	return &VoiceServiceHandler{
		Server:       s,
		ErrorHandler: errHandler,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *VoiceServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.voice.v1.VoiceService/Connect":
		{
			var err error

			in := make(chan *v1.ClientSignal)
			err = nil

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			out := make(chan *v1.Signal)

			ws, err := h.upgrader.Upgrade(w, req, nil)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			go func() {
				for {
					select {
					case msg, ok := <-out:
						if !ok {
							ws.WriteMessage(websocket.CloseMessage, []byte{})
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

			h.Server.Connect(req.Context(), in, out, req.Header)
		}

	}
}
