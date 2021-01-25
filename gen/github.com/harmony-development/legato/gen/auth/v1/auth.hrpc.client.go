package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
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

// Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)

func (client *AuthServiceClient) Federate(r *v1.FederateRequest) (*v1.FederateRequest, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("%s/protocol.auth.v1.AuthService/Federate", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.FederateRequest{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

// Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)

func (client *AuthServiceClient) LoginFederated(r *v1.LoginFederatedRequest) (*v1.LoginFederatedRequest, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("%s/protocol.auth.v1.AuthService/LoginFederated", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.LoginFederatedRequest{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

// Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)

func (client *AuthServiceClient) Key(r *empty.Empty) (*v1.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("%s/protocol.auth.v1.AuthService/Key", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

// Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)

func (client *AuthServiceClient) BeginAuth(r *empty.Empty) (*v1.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("%s/protocol.auth.v1.AuthService/BeginAuth", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

// Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)

func (client *AuthServiceClient) NextStep(r *v1.NextStepRequest) (*v1.NextStepRequest, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("%s/protocol.auth.v1.AuthService/NextStep", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.NextStepRequest{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

// Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)

func (client *AuthServiceClient) StepBack(r *v1.StepBackRequest) (*v1.StepBackRequest, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("%s/protocol.auth.v1.AuthService/StepBack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.StepBackRequest{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}
