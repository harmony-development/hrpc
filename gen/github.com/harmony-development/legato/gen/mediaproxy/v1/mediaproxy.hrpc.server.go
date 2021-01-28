package v1

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"
import "github.com/gorilla/websocket"
import "google.golang.org/protobuf/types/descriptorpb"

var Mediaproxyᐳv1ᐳmediaproxy *descriptorpb.FileDescriptorProto = new(descriptorpb.FileDescriptorProto)

func init() {
	data := []byte("\n\x1emediaproxy/v1/mediaproxy.proto\x12\x16protocol.mediaproxy.v1\x1a\x1bharmonytypes/v1/types.proto\"\xaa\x01\n\fSiteMetadata\x12\x1d\n\nsite_title\x18\x01 \x01(\tR\tsiteTitle\x12\x1d\n\npage_title\x18\x02 \x01(\tR\tpageTitle\x12\x12\n\x04kind\x18\x03 \x01(\tR\x04kind\x12 \n\vdescription\x18\x04 \x01(\tR\vdescription\x12\x10\n\x03url\x18\x05 \x01(\tR\x03url\x12\x14\n\x05image\x18\x06 \x01(\tR\x05image\",\n\x18FetchLinkMetadataRequest\x12\x10\n\x03url\x18\x01 \x01(\tR\x03url\"&\n\x12InstantViewRequest\x12\x10\n\x03url\x18\x01 \x01(\tR\x03url\"\x8c\x01\n\x13InstantViewResponse\x12@\n\bmetadata\x18\x01 \x01(\v2$.protocol.mediaproxy.v1.SiteMetadataR\bmetadata\x12\x18\n\acontent\x18\x02 \x01(\tR\acontent\x12\x19\n\bis_valid\x18\x03 \x01(\bR\aisValid\"B\n\x16CanInstantViewResponse\x12(\n\x10can_instant_view\x18\x01 \x01(\bR\x0ecanInstantView2\xeb\x02\n\x11MediaProxyService\x12r\n\x11FetchLinkMetadata\x120.protocol.mediaproxy.v1.FetchLinkMetadataRequest\x1a$.protocol.mediaproxy.v1.SiteMetadata\"\x05\x9aD\x02\b\x01\x12m\n\vInstantView\x12*.protocol.mediaproxy.v1.InstantViewRequest\x1a+.protocol.mediaproxy.v1.InstantViewResponse\"\x05\x9aD\x02\b\x01\x12s\n\x0eCanInstantView\x12*.protocol.mediaproxy.v1.InstantViewRequest\x1a..protocol.mediaproxy.v1.CanInstantViewResponse\"\x05\x9aD\x02\b\x01B9Z7github.com/harmony-development/legato/gen/mediaproxy/v1J\xe5\b\n\x06\x12\x04\x00\x00*\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\b\n\x01\x02\x12\x03\x02\x00\x1f\n\t\n\x02\x03\x00\x12\x03\x03\x00%\n\b\n\x01\b\x12\x03\x05\x00N\n\t\n\x02\b\v\x12\x03\x05\x00N\n\n\n\x02\x04\x00\x12\x04\a\x00\x0e\x01\n\n\n\x03\x04\x00\x01\x12\x03\a\b\x14\n\v\n\x04\x04\x00\x02\x00\x12\x03\b\x04\x1a\n\f\n\x05\x04\x00\x02\x00\x05\x12\x03\b\x04\n\n\f\n\x05\x04\x00\x02\x00\x01\x12\x03\b\v\x15\n\f\n\x05\x04\x00\x02\x00\x03\x12\x03\b\x18\x19\n\v\n\x04\x04\x00\x02\x01\x12\x03\t\x04\x1a\n\f\n\x05\x04\x00\x02\x01\x05\x12\x03\t\x04\n\n\f\n\x05\x04\x00\x02\x01\x01\x12\x03\t\v\x15\n\f\n\x05\x04\x00\x02\x01\x03\x12\x03\t\x18\x19\n\v\n\x04\x04\x00\x02\x02\x12\x03\n\x04\x14\n\f\n\x05\x04\x00\x02\x02\x05\x12\x03\n\x04\n\n\f\n\x05\x04\x00\x02\x02\x01\x12\x03\n\v\x0f\n\f\n\x05\x04\x00\x02\x02\x03\x12\x03\n\x12\x13\n\v\n\x04\x04\x00\x02\x03\x12\x03\v\x04\x1b\n\f\n\x05\x04\x00\x02\x03\x05\x12\x03\v\x04\n\n\f\n\x05\x04\x00\x02\x03\x01\x12\x03\v\v\x16\n\f\n\x05\x04\x00\x02\x03\x03\x12\x03\v\x19\x1a\n\v\n\x04\x04\x00\x02\x04\x12\x03\f\x04\x13\n\f\n\x05\x04\x00\x02\x04\x05\x12\x03\f\x04\n\n\f\n\x05\x04\x00\x02\x04\x01\x12\x03\f\v\x0e\n\f\n\x05\x04\x00\x02\x04\x03\x12\x03\f\x11\x12\n\v\n\x04\x04\x00\x02\x05\x12\x03\r\x04\x15\n\f\n\x05\x04\x00\x02\x05\x05\x12\x03\r\x04\n\n\f\n\x05\x04\x00\x02\x05\x01\x12\x03\r\v\x10\n\f\n\x05\x04\x00\x02\x05\x03\x12\x03\r\x13\x14\n\n\n\x02\x04\x01\x12\x04\x10\x00\x12\x01\n\n\n\x03\x04\x01\x01\x12\x03\x10\b \n\v\n\x04\x04\x01\x02\x00\x12\x03\x11\x04\x13\n\f\n\x05\x04\x01\x02\x00\x05\x12\x03\x11\x04\n\n\f\n\x05\x04\x01\x02\x00\x01\x12\x03\x11\v\x0e\n\f\n\x05\x04\x01\x02\x00\x03\x12\x03\x11\x11\x12\n\n\n\x02\x04\x02\x12\x04\x14\x00\x16\x01\n\n\n\x03\x04\x02\x01\x12\x03\x14\b\x1a\n\v\n\x04\x04\x02\x02\x00\x12\x03\x15\x04\x13\n\f\n\x05\x04\x02\x02\x00\x05\x12\x03\x15\x04\n\n\f\n\x05\x04\x02\x02\x00\x01\x12\x03\x15\v\x0e\n\f\n\x05\x04\x02\x02\x00\x03\x12\x03\x15\x11\x12\n\n\n\x02\x04\x03\x12\x04\x17\x00\x1b\x01\n\n\n\x03\x04\x03\x01\x12\x03\x17\b\x1b\n\v\n\x04\x04\x03\x02\x00\x12\x03\x18\x04\x1e\n\f\n\x05\x04\x03\x02\x00\x06\x12\x03\x18\x04\x10\n\f\n\x05\x04\x03\x02\x00\x01\x12\x03\x18\x11\x19\n\f\n\x05\x04\x03\x02\x00\x03\x12\x03\x18\x1c\x1d\n\v\n\x04\x04\x03\x02\x01\x12\x03\x19\x04\x17\n\f\n\x05\x04\x03\x02\x01\x05\x12\x03\x19\x04\n\n\f\n\x05\x04\x03\x02\x01\x01\x12\x03\x19\v\x12\n\f\n\x05\x04\x03\x02\x01\x03\x12\x03\x19\x15\x16\n\v\n\x04\x04\x03\x02\x02\x12\x03\x1a\x04\x16\n\f\n\x05\x04\x03\x02\x02\x05\x12\x03\x1a\x04\b\n\f\n\x05\x04\x03\x02\x02\x01\x12\x03\x1a\t\x11\n\f\n\x05\x04\x03\x02\x02\x03\x12\x03\x1a\x14\x15\n\n\n\x02\x04\x04\x12\x04\x1c\x00\x1e\x01\n\n\n\x03\x04\x04\x01\x12\x03\x1c\b\x1e\n\v\n\x04\x04\x04\x02\x00\x12\x03\x1d\x04\x1e\n\f\n\x05\x04\x04\x02\x00\x05\x12\x03\x1d\x04\b\n\f\n\x05\x04\x04\x02\x00\x01\x12\x03\x1d\t\x19\n\f\n\x05\x04\x04\x02\x00\x03\x12\x03\x1d\x1c\x1d\n\n\n\x02\x06\x00\x12\x04 \x00*\x01\n\n\n\x03\x06\x00\x01\x12\x03 \b\x19\n\f\n\x04\x06\x00\x02\x00\x12\x04!\x04#\x05\n\f\n\x05\x06\x00\x02\x00\x01\x12\x03!\b\x19\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03!\x1a2\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03!=I\n\f\n\x05\x06\x00\x02\x00\x04\x12\x03\"\bI\n\x0f\n\b\x06\x00\x02\x00\x04\xc3\b\x01\x12\x03\"\bI\n\f\n\x04\x06\x00\x02\x01\x12\x04$\x04&\x05\n\f\n\x05\x06\x00\x02\x01\x01\x12\x03$\b\x13\n\f\n\x05\x06\x00\x02\x01\x02\x12\x03$\x14&\n\f\n\x05\x06\x00\x02\x01\x03\x12\x03$1D\n\f\n\x05\x06\x00\x02\x01\x04\x12\x03%\bI\n\x0f\n\b\x06\x00\x02\x01\x04\xc3\b\x01\x12\x03%\bI\n\f\n\x04\x06\x00\x02\x02\x12\x04'\x04)\x05\n\f\n\x05\x06\x00\x02\x02\x01\x12\x03'\b\x16\n\f\n\x05\x06\x00\x02\x02\x02\x12\x03'\x17)\n\f\n\x05\x06\x00\x02\x02\x03\x12\x03'4J\n\f\n\x05\x06\x00\x02\x02\x04\x12\x03(\bI\n\x0f\n\b\x06\x00\x02\x02\x04\xc3\b\x01\x12\x03(\bIb\x06proto3")

	err := proto.Unmarshal(data, Mediaproxyᐳv1ᐳmediaproxy)
	if err != nil {
		panic(err)
	}
}

type MediaProxyServiceServer interface {
	FetchLinkMetadata(ctx context.Context, r *FetchLinkMetadataRequest, headers http.Header) (resp *SiteMetadata, err error)

	InstantView(ctx context.Context, r *InstantViewRequest, headers http.Header) (resp *InstantViewResponse, err error)

	CanInstantView(ctx context.Context, r *InstantViewRequest, headers http.Header) (resp *CanInstantViewResponse, err error)
}

type MediaProxyServiceHandler struct {
	Server       MediaProxyServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	UnaryPre     func(d *descriptorpb.FileDescriptorProto, f func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error)) func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error)
	upgrader     websocket.Upgrader
}

func NewMediaProxyServiceHandler(s MediaProxyServiceServer, errHandler func(err error, w http.ResponseWriter)) *MediaProxyServiceHandler {
	return &MediaProxyServiceHandler{
		Server:       s,
		ErrorHandler: errHandler,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
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

			requestProto := new(FetchLinkMetadataRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.FetchLinkMetadata(c, req.(*FetchLinkMetadataRequest), headers)
			}

			if h.UnaryPre != nil {
				invoker = h.UnaryPre(Mediaproxyᐳv1ᐳmediaproxy, invoker)
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

	case "/protocol.mediaproxy.v1.MediaProxyService/InstantView":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(InstantViewRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.InstantView(c, req.(*InstantViewRequest), headers)
			}

			if h.UnaryPre != nil {
				invoker = h.UnaryPre(Mediaproxyᐳv1ᐳmediaproxy, invoker)
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

	case "/protocol.mediaproxy.v1.MediaProxyService/CanInstantView":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(InstantViewRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			invoker := func(c context.Context, req proto.Message, headers http.Header) (proto.Message, error) {
				return h.Server.CanInstantView(c, req.(*InstantViewRequest), headers)
			}

			if h.UnaryPre != nil {
				invoker = h.UnaryPre(Mediaproxyᐳv1ᐳmediaproxy, invoker)
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

	}
}
