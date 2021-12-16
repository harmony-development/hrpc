// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package chat

import (
	bytes "bytes"
	context "context"
	errors "errors"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	ioutil "io/ioutil"
	http "net/http"
	httptest "net/http/httptest"
)

type ChatServiceClient interface {
	SendMessage(context.Context, *Message) (*emptypb.Empty, error)
	StreamMessages(context.Context, chan *emptypb.Empty) (chan *Message, error)
}

type HTTPChatServiceClient struct {
	Client  http.Client
	BaseURL string
}

func (client *HTTPChatServiceClient) SendMessage(req *Message) (*emptypb.Empty, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	resp, err := client.Client.Post(client.BaseURL+"/chat.ChatService.SendMessage/", "application/hrpc", reader)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &emptypb.Empty{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPChatServiceClient) StreamMessages(req *emptypb.Empty) (chan *Message, error) {
	return nil, errors.New("unimplemented")
}

type HTTPTestChatServiceClient struct {
	Client interface {
		Test(*http.Request, ...int) (*http.Response, error)
	}
}

func (client *HTTPTestChatServiceClient) SendMessage(req *Message) (*emptypb.Empty, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/chat.ChatService.SendMessage/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &emptypb.Empty{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestChatServiceClient) StreamMessages(req *emptypb.Empty) (chan *Message, error) {
	return nil, errors.New("unimplemented")
}