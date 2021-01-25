package v1

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"

import "github.com/harmony-development/legato/gen/mediaproxy/v1"

type MediaProxyServiceServer interface {
	FetchLinkMetadata(ctx context.Context, r *v1.FetchLinkMetadataRequest, headers http.Header) (resp *v1.SiteMetadata, err error)

	InstantView(ctx context.Context, r *v1.InstantViewRequest, headers http.Header) (resp *v1.InstantViewResponse, err error)

	CanInstantView(ctx context.Context, r *v1.InstantViewRequest, headers http.Header) (resp *v1.CanInstantViewResponse, err error)
}

type MediaProxyServiceHandler struct {
	Server       MediaProxyServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
}

func (h *MediaProxyServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.mediaproxy.v1.MediaProxyService/FetchLinkMetadata":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(v1.FetchLinkMetadataRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.FetchLinkMetadata(req.Context(), nil, req.Header)

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

	case "/protocol.mediaproxy.v1.MediaProxyService/InstantView":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(v1.InstantViewRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.InstantView(req.Context(), nil, req.Header)

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

	case "/protocol.mediaproxy.v1.MediaProxyService/CanInstantView":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(v1.InstantViewRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.CanInstantView(req.Context(), nil, req.Header)

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

	}
}
