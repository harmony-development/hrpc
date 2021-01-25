package v1

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"

type VoiceServiceServer interface {
}

type VoiceServiceHandler struct {
	Server       VoiceServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
}

func (h *VoiceServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	}
}
