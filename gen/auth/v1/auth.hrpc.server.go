package v1

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"
import "github.com/gorilla/websocket"

import "github.com/golang/protobuf/ptypes/empty"

import "github.com/harmony-development/legato/gen/auth/v1"

type AuthServiceServer interface {
	Federate(ctx context.Context, r *v1.FederateRequest, headers http.Header) (resp *v1.FederateReply, err error)

	LoginFederated(ctx context.Context, r *v1.LoginFederatedRequest, headers http.Header) (resp *v1.Session, err error)

	Key(ctx context.Context, r *empty.Empty, headers http.Header) (resp *v1.KeyReply, err error)

	BeginAuth(ctx context.Context, r *empty.Empty, headers http.Header) (resp *v1.BeginAuthResponse, err error)

	NextStep(ctx context.Context, r *v1.NextStepRequest, headers http.Header) (resp *v1.AuthStep, err error)

	StepBack(ctx context.Context, r *v1.StepBackRequest, headers http.Header) (resp *v1.AuthStep, err error)

	StreamSteps(ctx context.Context, r *v1.StreamStepsRequest, out chan *v1.AuthStep, headers http.Header)
}

type AuthServiceHandler struct {
	Server       AuthServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	upgrader     websocket.Upgrader
}

func NewAuthServiceHandler(s AuthServiceServer, errHandler func(err error, w http.ResponseWriter)) *AuthServiceHandler {
	return &AuthServiceHandler{
		Server:       s,
		ErrorHandler: errHandler,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *AuthServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.auth.v1.AuthService/Federate":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(v1.FederateRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.Federate(req.Context(), nil, req.Header)

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

	case "/protocol.auth.v1.AuthService/LoginFederated":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(v1.LoginFederatedRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.LoginFederated(req.Context(), nil, req.Header)

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

	case "/protocol.auth.v1.AuthService/Key":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(empty.Empty)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.Key(req.Context(), nil, req.Header)

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

	case "/protocol.auth.v1.AuthService/BeginAuth":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(empty.Empty)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.BeginAuth(req.Context(), nil, req.Header)

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

	case "/protocol.auth.v1.AuthService/NextStep":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(v1.NextStepRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.NextStep(req.Context(), nil, req.Header)

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

	case "/protocol.auth.v1.AuthService/StepBack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(v1.StepBackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.StepBack(req.Context(), nil, req.Header)

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

	case "/protocol.auth.v1.AuthService/StreamSteps":
		{
			var err error

			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			in := new(v1.StreamStepsRequest)
			err = proto.Unmarshal(body, in)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			out := make(chan *v1.AuthStep)

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

							close(out)
							return
						}

						response, err := proto.Marshal(msg)
						if err != nil {

							close(out)
							return
						}

						w.Write(response)
						if err := w.Close(); err != nil {

							close(out)
							return
						}
					}
				}
			}()

			h.Server.StreamSteps(req.Context(), in, out, req.Header)
		}

	}
}
