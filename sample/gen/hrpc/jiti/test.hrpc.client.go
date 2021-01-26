package jiti

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/protobuf/proto"
)

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
	resp, err := client.client.Post(fmt.Sprintf("%s/test.Mu/Mu", client.serverURL), "application/octet-stream", bytes.NewReader(input))
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
