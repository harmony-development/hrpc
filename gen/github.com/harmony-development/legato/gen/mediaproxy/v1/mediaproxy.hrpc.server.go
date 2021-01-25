package v1

import "context"
import "net/http"

import "github.com/harmony-development/legato/gen/mediaproxy/v1"

type MediaProxyServiceServer interface {
	FetchLinkMetadata(ctx context.Context, r *v1.FetchLinkMetadataRequest) (resp v1.FetchLinkMetadataRequest, err error)

	InstantView(ctx context.Context, r *v1.InstantViewRequest) (resp v1.InstantViewRequest, err error)

	CanInstantView(ctx context.Context, r *v1.InstantViewRequest) (resp v1.InstantViewRequest, err error)
}

type MediaProxyServiceHandler struct {
	Server MediaProxyServiceServer
}

func (h *MediaProxyServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.mediaproxy.v1.MediaProxyService/FetchLinkMetadata":
		{
			panic("unimplemented")
		}

	case "/protocol.mediaproxy.v1.MediaProxyService/InstantView":
		{
			panic("unimplemented")
		}

	case "/protocol.mediaproxy.v1.MediaProxyService/CanInstantView":
		{
			panic("unimplemented")
		}

	}
}
