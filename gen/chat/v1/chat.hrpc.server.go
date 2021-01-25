package v1

import "context"
import "net/http"

import "github.com/golang/protobuf/ptypes/empty"

import "github.com/harmony-development/legato/gen/chat/v1"

type ChatServiceServer interface {
	CreateGuild(ctx context.Context, r *v1.CreateGuildRequest) (resp v1.CreateGuildResponse, err error)

	CreateInvite(ctx context.Context, r *v1.CreateInviteRequest) (resp v1.CreateInviteResponse, err error)

	CreateChannel(ctx context.Context, r *v1.CreateChannelRequest) (resp v1.CreateChannelResponse, err error)

	CreateEmotePack(ctx context.Context, r *v1.CreateEmotePackRequest) (resp v1.CreateEmotePackResponse, err error)

	GetGuildList(ctx context.Context, r *v1.GetGuildListRequest) (resp v1.GetGuildListResponse, err error)

	AddGuildToGuildList(ctx context.Context, r *v1.AddGuildToGuildListRequest) (resp v1.AddGuildToGuildListResponse, err error)

	RemoveGuildFromGuildList(ctx context.Context, r *v1.RemoveGuildFromGuildListRequest) (resp v1.RemoveGuildFromGuildListResponse, err error)

	GetGuild(ctx context.Context, r *v1.GetGuildRequest) (resp v1.GetGuildResponse, err error)

	GetGuildInvites(ctx context.Context, r *v1.GetGuildInvitesRequest) (resp v1.GetGuildInvitesResponse, err error)

	GetGuildMembers(ctx context.Context, r *v1.GetGuildMembersRequest) (resp v1.GetGuildMembersResponse, err error)

	GetGuildChannels(ctx context.Context, r *v1.GetGuildChannelsRequest) (resp v1.GetGuildChannelsResponse, err error)

	GetChannelMessages(ctx context.Context, r *v1.GetChannelMessagesRequest) (resp v1.GetChannelMessagesResponse, err error)

	GetMessage(ctx context.Context, r *v1.GetMessageRequest) (resp v1.GetMessageResponse, err error)

	GetEmotePacks(ctx context.Context, r *v1.GetEmotePacksRequest) (resp v1.GetEmotePacksResponse, err error)

	GetEmotePackEmotes(ctx context.Context, r *v1.GetEmotePackEmotesRequest) (resp v1.GetEmotePackEmotesResponse, err error)

	UpdateGuildInformation(ctx context.Context, r *v1.UpdateGuildInformationRequest) (resp empty.Empty, err error)

	UpdateChannelInformation(ctx context.Context, r *v1.UpdateChannelInformationRequest) (resp empty.Empty, err error)

	UpdateChannelOrder(ctx context.Context, r *v1.UpdateChannelOrderRequest) (resp empty.Empty, err error)

	UpdateMessage(ctx context.Context, r *v1.UpdateMessageRequest) (resp empty.Empty, err error)

	AddEmoteToPack(ctx context.Context, r *v1.AddEmoteToPackRequest) (resp empty.Empty, err error)

	DeleteGuild(ctx context.Context, r *v1.DeleteGuildRequest) (resp empty.Empty, err error)

	DeleteInvite(ctx context.Context, r *v1.DeleteInviteRequest) (resp empty.Empty, err error)

	DeleteChannel(ctx context.Context, r *v1.DeleteChannelRequest) (resp empty.Empty, err error)

	DeleteMessage(ctx context.Context, r *v1.DeleteMessageRequest) (resp empty.Empty, err error)

	DeleteEmoteFromPack(ctx context.Context, r *v1.DeleteEmoteFromPackRequest) (resp empty.Empty, err error)

	DeleteEmotePack(ctx context.Context, r *v1.DeleteEmotePackRequest) (resp empty.Empty, err error)

	DequipEmotePack(ctx context.Context, r *v1.DequipEmotePackRequest) (resp empty.Empty, err error)

	JoinGuild(ctx context.Context, r *v1.JoinGuildRequest) (resp v1.JoinGuildResponse, err error)

	LeaveGuild(ctx context.Context, r *v1.LeaveGuildRequest) (resp empty.Empty, err error)

	TriggerAction(ctx context.Context, r *v1.TriggerActionRequest) (resp empty.Empty, err error)

	SendMessage(ctx context.Context, r *v1.SendMessageRequest) (resp v1.SendMessageResponse, err error)

	QueryHasPermission(ctx context.Context, r *v1.QueryPermissionsRequest) (resp v1.QueryPermissionsResponse, err error)

	SetPermissions(ctx context.Context, r *v1.SetPermissionsRequest) (resp empty.Empty, err error)

	GetPermissions(ctx context.Context, r *v1.GetPermissionsRequest) (resp v1.GetPermissionsResponse, err error)

	MoveRole(ctx context.Context, r *v1.MoveRoleRequest) (resp v1.MoveRoleResponse, err error)

	GetGuildRoles(ctx context.Context, r *v1.GetGuildRolesRequest) (resp v1.GetGuildRolesResponse, err error)

	AddGuildRole(ctx context.Context, r *v1.AddGuildRoleRequest) (resp v1.AddGuildRoleResponse, err error)

	ModifyGuildRole(ctx context.Context, r *v1.ModifyGuildRoleRequest) (resp empty.Empty, err error)

	DeleteGuildRole(ctx context.Context, r *v1.DeleteGuildRoleRequest) (resp empty.Empty, err error)

	ManageUserRoles(ctx context.Context, r *v1.ManageUserRolesRequest) (resp empty.Empty, err error)

	GetUserRoles(ctx context.Context, r *v1.GetUserRolesRequest) (resp v1.GetUserRolesResponse, err error)

	GetUser(ctx context.Context, r *v1.GetUserRequest) (resp v1.GetUserResponse, err error)

	GetUserMetadata(ctx context.Context, r *v1.GetUserMetadataRequest) (resp v1.GetUserMetadataResponse, err error)

	ProfileUpdate(ctx context.Context, r *v1.ProfileUpdateRequest) (resp empty.Empty, err error)

	Typing(ctx context.Context, r *v1.TypingRequest) (resp empty.Empty, err error)

	PreviewGuild(ctx context.Context, r *v1.PreviewGuildRequest) (resp v1.PreviewGuildResponse, err error)
}

type ChatServiceHandler struct {
	Server ChatServiceServer
}

func (h *ChatServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.chat.v1.ChatService/CreateGuild":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/CreateInvite":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/CreateChannel":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/CreateEmotePack":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetGuildList":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/AddGuildToGuildList":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/RemoveGuildFromGuildList":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetGuild":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetGuildInvites":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetGuildMembers":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetGuildChannels":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetChannelMessages":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetMessage":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetEmotePacks":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetEmotePackEmotes":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/UpdateGuildInformation":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/UpdateChannelInformation":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/UpdateChannelOrder":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/UpdateMessage":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/AddEmoteToPack":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/DeleteGuild":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/DeleteInvite":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/DeleteChannel":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/DeleteMessage":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/DeleteEmoteFromPack":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/DeleteEmotePack":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/DequipEmotePack":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/JoinGuild":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/LeaveGuild":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/TriggerAction":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/SendMessage":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/QueryHasPermission":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/SetPermissions":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetPermissions":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/MoveRole":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetGuildRoles":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/AddGuildRole":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/ModifyGuildRole":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/DeleteGuildRole":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/ManageUserRoles":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetUserRoles":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/StreamEvents":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/Sync":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetUser":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/GetUserMetadata":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/ProfileUpdate":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/Typing":
		{
			panic("unimplemented")
		}

	case "/protocol.chat.v1.ChatService/PreviewGuild":
		{
			panic("unimplemented")
		}

	}
}
