package jiti

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"
import "github.com/gorilla/websocket"

type MuServer interface {
	Mu(ctx context.Context, r *Ping, headers http.Header) (resp *Pong, err error)

	MuMute(ctx context.Context, in chan *Ping, out chan *Pong, headers http.Header)
}

type MuHandler struct {
	Server       MuServer
	ErrorHandler func(err error, w http.ResponseWriter)
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

			resp, err := h.Server.Mu(req.Context(), requestProto, req.Header)

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
