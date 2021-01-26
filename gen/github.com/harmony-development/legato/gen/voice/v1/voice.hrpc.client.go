package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "github.com/gorilla/websocket"
import "net/url"
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
	u := url.URL{Scheme: "ws", Host: client.serverURL, Path: "/protocol.voice.v1.VoiceService/Connect"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	in = make(chan *v1.ClientSignal)
	out = make(chan *v1.Signal)

	go func() {
		defer c.Close()

		msgs := make(chan []byte)

		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					close(msgs)
					break
				}
				msgs <- message
			}
		}()

		for {
			select {
			case msg, ok := <-msgs:
				thing := new(v1.Signal)
				proto.Unmarshal
			}
		}
	}()

	return in, out, nil
}
