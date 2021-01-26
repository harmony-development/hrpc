package v1

import "net/http"
import "google.golang.org/protobuf/proto"
import "io/ioutil"
import "fmt"
import "github.com/gorilla/websocket"
import "net/url"
import "bytes"

import "github.com/golang/protobuf/ptypes/empty"

import "github.com/harmony-development/legato/gen/chat/v1"

type ChatServiceClient struct {
	client    *http.Client
	serverURL string
}

func NewChatServiceClient(url string) *ChatServiceClient {
	return &ChatServiceClient{
		client:    &http.Client{},
		serverURL: url,
	}
}

func (client *ChatServiceClient) CreateGuild(r *v1.CreateGuildRequest) (*v1.CreateGuildResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/CreateGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.CreateGuildResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) CreateInvite(r *v1.CreateInviteRequest) (*v1.CreateInviteResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/CreateInvite", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.CreateInviteResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) CreateChannel(r *v1.CreateChannelRequest) (*v1.CreateChannelResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/CreateChannel", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.CreateChannelResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) CreateEmotePack(r *v1.CreateEmotePackRequest) (*v1.CreateEmotePackResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/CreateEmotePack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.CreateEmotePackResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildList(r *v1.GetGuildListRequest) (*v1.GetGuildListResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildList", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetGuildListResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) AddGuildToGuildList(r *v1.AddGuildToGuildListRequest) (*v1.AddGuildToGuildListResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/AddGuildToGuildList", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.AddGuildToGuildListResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) RemoveGuildFromGuildList(r *v1.RemoveGuildFromGuildListRequest) (*v1.RemoveGuildFromGuildListResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/RemoveGuildFromGuildList", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.RemoveGuildFromGuildListResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuild(r *v1.GetGuildRequest) (*v1.GetGuildResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetGuildResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildInvites(r *v1.GetGuildInvitesRequest) (*v1.GetGuildInvitesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildInvites", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetGuildInvitesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildMembers(r *v1.GetGuildMembersRequest) (*v1.GetGuildMembersResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildMembers", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetGuildMembersResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildChannels(r *v1.GetGuildChannelsRequest) (*v1.GetGuildChannelsResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildChannels", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetGuildChannelsResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetChannelMessages(r *v1.GetChannelMessagesRequest) (*v1.GetChannelMessagesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetChannelMessages", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetChannelMessagesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetMessage(r *v1.GetMessageRequest) (*v1.GetMessageResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetMessage", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetMessageResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetEmotePacks(r *v1.GetEmotePacksRequest) (*v1.GetEmotePacksResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetEmotePacks", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetEmotePacksResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetEmotePackEmotes(r *v1.GetEmotePackEmotesRequest) (*v1.GetEmotePackEmotesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetEmotePackEmotes", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetEmotePackEmotesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) UpdateGuildInformation(r *v1.UpdateGuildInformationRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/UpdateGuildInformation", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) UpdateChannelInformation(r *v1.UpdateChannelInformationRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/UpdateChannelInformation", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) UpdateChannelOrder(r *v1.UpdateChannelOrderRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/UpdateChannelOrder", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) UpdateMessage(r *v1.UpdateMessageRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/UpdateMessage", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) AddEmoteToPack(r *v1.AddEmoteToPackRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/AddEmoteToPack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteGuild(r *v1.DeleteGuildRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteInvite(r *v1.DeleteInviteRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteInvite", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteChannel(r *v1.DeleteChannelRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteChannel", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteMessage(r *v1.DeleteMessageRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteMessage", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteEmoteFromPack(r *v1.DeleteEmoteFromPackRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteEmoteFromPack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteEmotePack(r *v1.DeleteEmotePackRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteEmotePack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DequipEmotePack(r *v1.DequipEmotePackRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DequipEmotePack", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) JoinGuild(r *v1.JoinGuildRequest) (*v1.JoinGuildResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/JoinGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.JoinGuildResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) LeaveGuild(r *v1.LeaveGuildRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/LeaveGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) TriggerAction(r *v1.TriggerActionRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/TriggerAction", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) SendMessage(r *v1.SendMessageRequest) (*v1.SendMessageResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/SendMessage", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.SendMessageResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) QueryHasPermission(r *v1.QueryPermissionsRequest) (*v1.QueryPermissionsResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/QueryHasPermission", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.QueryPermissionsResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) SetPermissions(r *v1.SetPermissionsRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/SetPermissions", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetPermissions(r *v1.GetPermissionsRequest) (*v1.GetPermissionsResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetPermissions", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetPermissionsResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) MoveRole(r *v1.MoveRoleRequest) (*v1.MoveRoleResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/MoveRole", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.MoveRoleResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetGuildRoles(r *v1.GetGuildRolesRequest) (*v1.GetGuildRolesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetGuildRoles", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetGuildRolesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) AddGuildRole(r *v1.AddGuildRoleRequest) (*v1.AddGuildRoleResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/AddGuildRole", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.AddGuildRoleResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) ModifyGuildRole(r *v1.ModifyGuildRoleRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/ModifyGuildRole", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) DeleteGuildRole(r *v1.DeleteGuildRoleRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/DeleteGuildRole", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) ManageUserRoles(r *v1.ManageUserRolesRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/ManageUserRoles", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetUserRoles(r *v1.GetUserRolesRequest) (*v1.GetUserRolesResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetUserRoles", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetUserRolesResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) StreamEvents() (in chan *v1.StreamEventsRequest, out chan *v1.Event, err error) {
	u := url.URL{Scheme: "ws", Host: client.serverURL, Path: "/protocol.chat.v1.ChatService/StreamEvents"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	in = make(chan *v1.StreamEventsRequest)
	out = make(chan *v1.Event)

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
				thing := new(v1.Event)
				proto.Unmarshal
			}
		}
	}()

	return in, out, nil
}

func (client *ChatServiceClient) Sync(r *v1.SyncRequest) (chan *v1.SyncEvent, error) {
	panic("unimplemented")
}

func (client *ChatServiceClient) GetUser(r *v1.GetUserRequest) (*v1.GetUserResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetUser", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetUserResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) GetUserMetadata(r *v1.GetUserMetadataRequest) (*v1.GetUserMetadataResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/GetUserMetadata", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.GetUserMetadataResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) ProfileUpdate(r *v1.ProfileUpdateRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/ProfileUpdate", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) Typing(r *v1.TypingRequest) (*empty.Empty, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/Typing", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &empty.Empty{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}

func (client *ChatServiceClient) PreviewGuild(r *v1.PreviewGuildRequest) (*v1.PreviewGuildResponse, error) {
	input, err := proto.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("could not martial request: %w", err)
	}
	resp, err := client.client.Post(fmt.Sprintf("http://%s/protocol.chat.v1.ChatService/PreviewGuild", client.serverURL), "application/octet-stream", bytes.NewReader(input))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	output := &v1.PreviewGuildResponse{}
	err = proto.Unmarshal(data, output)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return output, nil
}
