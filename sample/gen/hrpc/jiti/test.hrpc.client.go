package jiti

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "github.com/gorilla/websocket"
import "net/url"
import "bytes"

type MuClient struct {
	client    *http.Client
	serverURL string
}

func NewMuClient(url string) *MuClient {
	return &MuClient{
		client:    &http.Client{},
		serverURL: url,
	}
}

func (client *MuClient) Mu(r *Ping) (*Pong, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/test.Mu/Mu", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &Pong{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *MuClient) MuMute() (in chan<- *Ping, out <-chan *Pong, err error) {
	u := url.URL{Scheme: "ws", Host: client.serverURL, Path: "/test.Mu/MuMute"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	inC := make(chan *Ping)
	outC := make(chan *Pong)

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
				if !ok {
					return
				}

				thing := new(Pong)
				err = proto.Unmarshal(msg, thing)
				if err != nil {
					return
				}

				outC <- thing
			case send, ok := <-inC:
				if !ok {
					return
				}

				data, err := proto.Marshal(send)
				if err != nil {
					return
				}

				err = c.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					return
				}
			}
		}
	}()

	return inC, outC, nil
}
