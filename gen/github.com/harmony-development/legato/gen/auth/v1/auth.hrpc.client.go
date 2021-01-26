package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "github.com/gorilla/websocket"
import "net/url"
import "bytes"

import "github.com/golang/protobuf/ptypes/empty"

import "github.com/harmony-development/legato/gen/auth/v1"

type AuthServiceClient struct {
	client    *http.Client
	serverURL string
}

func NewAuthServiceClient(url string) *AuthServiceClient {
	return &AuthServiceClient{
		client:    &http.Client{},
		serverURL: url,
	}
}

func (client *AuthServiceClient) Federate(r *v1.FederateRequest) (*v1.FederateReply, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.auth.v1.AuthService/Federate", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.FederateReply{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) LoginFederated(r *v1.LoginFederatedRequest) (*v1.Session, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.auth.v1.AuthService/LoginFederated", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.Session{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) Key(r *empty.Empty) (*v1.KeyReply, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.auth.v1.AuthService/Key", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.KeyReply{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) BeginAuth(r *empty.Empty) (*v1.BeginAuthResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.auth.v1.AuthService/BeginAuth", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.BeginAuthResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) NextStep(r *v1.NextStepRequest) (*v1.AuthStep, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.auth.v1.AuthService/NextStep", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.AuthStep{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) StepBack(r *v1.StepBackRequest) (*v1.AuthStep, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.auth.v1.AuthService/StepBack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.AuthStep{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *AuthServiceClient) StreamSteps(r *v1.StreamStepsRequest) (chan *v1.AuthStep, error) {
	panic("unimplemented")
}
