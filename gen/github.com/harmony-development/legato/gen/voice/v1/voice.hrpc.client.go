package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "bytes"

import "github.com/harmony-development/legato/gen/voice/v1"

type VoiceServiceClient struct {
	client    *http.Client
	serverURL string
}

func NewVoiceServiceClient(url string) *VoiceServiceClient {
	return &VoiceServiceClient{
		client:    &http.Client{},
		serverURL: url,
	}
}

func (client *VoiceServiceClient) Connect() (in chan *v1.ClientSignal, out chan *v1.Signal, err error) {
	panic("unimplemented")
}
