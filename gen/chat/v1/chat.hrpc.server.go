package v1

import "context"
import "net/http"
import "io/ioutil"
import "google.golang.org/protobuf/proto"
import "github.com/gorilla/websocket"
import "google.golang.org/protobuf/types/descriptorpb"

import "github.com/golang/protobuf/ptypes/empty"

var Chatᐳv1ᐳchat *descriptorpb.FileDescriptorProto = new(descriptorpb.FileDescriptorProto)

func init() {
	data := []byte("\n\x12chat/v1/chat.proto\x12\x10protocol.chat.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1bharmonytypes/v1/types.proto\x1a\x15chat/v1/profile.proto\x1a\x14chat/v1/guilds.proto\x1a\x16chat/v1/channels.proto\x1a\x16chat/v1/messages.proto\x1a\x14chat/v1/emotes.proto\x1a\x19chat/v1/permissions.proto\x1a\x17chat/v1/streaming.proto\x1a\x15chat/v1/postbox.proto2\xc0)\n\vChatService\x12f\n\vCreateGuild\x12$.protocol.chat.v1.CreateGuildRequest\x1a%.protocol.chat.v1.CreateGuildResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x1a\x00\x12~\n\fCreateInvite\x12%.protocol.chat.v1.CreateInviteRequest\x1a&.protocol.chat.v1.CreateInviteResponse\"\x1f\x9aD\x02\b\x01\x9aD\x17\x1a\x15invites.manage.create\x12\x82\x01\n\rCreateChannel\x12&.protocol.chat.v1.CreateChannelRequest\x1a'.protocol.chat.v1.CreateChannelResponse\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16channels.manage.create\x12m\n\x0fCreateEmotePack\x12(.protocol.chat.v1.CreateEmotePackRequest\x1a).protocol.chat.v1.CreateEmotePackResponse\"\x05\x9aD\x02\b\x01\x12d\n\fGetGuildList\x12%.protocol.chat.v1.GetGuildListRequest\x1a&.protocol.chat.v1.GetGuildListResponse\"\x05\x9aD\x02\b\x01\x12~\n\x13AddGuildToGuildList\x12,.protocol.chat.v1.AddGuildToGuildListRequest\x1a-.protocol.chat.v1.AddGuildToGuildListResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x10\x01\x12\x8d\x01\n\x18RemoveGuildFromGuildList\x121.protocol.chat.v1.RemoveGuildFromGuildListRequest\x1a2.protocol.chat.v1.RemoveGuildFromGuildListResponse\"\n\x9aD\x02\b\x01\x9aD\x02\x10\x01\x12X\n\bGetGuild\x12!.protocol.chat.v1.GetGuildRequest\x1a\".protocol.chat.v1.GetGuildResponse\"\x05\x9aD\x02\b\x01\x12~\n\x0fGetGuildInvites\x12(.protocol.chat.v1.GetGuildInvitesRequest\x1a).protocol.chat.v1.GetGuildInvitesResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\finvites.view\x12m\n\x0fGetGuildMembers\x12(.protocol.chat.v1.GetGuildMembersRequest\x1a).protocol.chat.v1.GetGuildMembersResponse\"\x05\x9aD\x02\b\x01\x12k\n\x10GetGuildChannels\x12).protocol.chat.v1.GetGuildChannelsRequest\x1a*.protocol.chat.v1.GetGuildChannelsResponse\"\x00\x12\x88\x01\n\x12GetChannelMessages\x12+.protocol.chat.v1.GetChannelMessagesRequest\x1a,.protocol.chat.v1.GetChannelMessagesResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.view\x12p\n\nGetMessage\x12#.protocol.chat.v1.GetMessageRequest\x1a$.protocol.chat.v1.GetMessageResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.view\x12g\n\rGetEmotePacks\x12&.protocol.chat.v1.GetEmotePacksRequest\x1a'.protocol.chat.v1.GetEmotePacksResponse\"\x05\x9aD\x02\b\x01\x12v\n\x12GetEmotePackEmotes\x12+.protocol.chat.v1.GetEmotePackEmotesRequest\x1a,.protocol.chat.v1.GetEmotePackEmotesResponse\"\x05\x9aD\x02\b\x01\x12\x8c\x01\n\x16UpdateGuildInformation\x12/.protocol.chat.v1.UpdateGuildInformationRequest\x1a\x16.google.protobuf.Empty\")\x9aD\x02\b\x01\x9aD!\x1a\x1fguild.manage.change-information\x12\x93\x01\n\x18UpdateChannelInformation\x121.protocol.chat.v1.UpdateChannelInformationRequest\x1a\x16.google.protobuf.Empty\",\x9aD\x02\b\x01\x9aD$\x1a\"channels.manage.change-information\x12y\n\x12UpdateChannelOrder\x12+.protocol.chat.v1.UpdateChannelOrderRequest\x1a\x16.google.protobuf.Empty\"\x1e\x9aD\x02\b\x01\x9aD\x16\x1a\x14channels.manage.move\x12h\n\rUpdateMessage\x12&.protocol.chat.v1.UpdateMessageRequest\x1a\x16.google.protobuf.Empty\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send\x12X\n\x0eAddEmoteToPack\x12'.protocol.chat.v1.AddEmoteToPackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12j\n\vDeleteGuild\x12$.protocol.chat.v1.DeleteGuildRequest\x1a\x16.google.protobuf.Empty\"\x1d\x9aD\x02\b\x01\x9aD\x15\x1a\x13guild.manage.delete\x12n\n\fDeleteInvite\x12%.protocol.chat.v1.DeleteInviteRequest\x1a\x16.google.protobuf.Empty\"\x1f\x9aD\x02\b\x01\x9aD\x17\x1a\x15invites.manage.delete\x12q\n\rDeleteChannel\x12&.protocol.chat.v1.DeleteChannelRequest\x1a\x16.google.protobuf.Empty\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16channels.manage.delete\x12V\n\rDeleteMessage\x12&.protocol.chat.v1.DeleteMessageRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12b\n\x13DeleteEmoteFromPack\x12,.protocol.chat.v1.DeleteEmoteFromPackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12Z\n\x0fDeleteEmotePack\x12(.protocol.chat.v1.DeleteEmotePackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12Z\n\x0fDequipEmotePack\x12(.protocol.chat.v1.DequipEmotePackRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12[\n\tJoinGuild\x12\".protocol.chat.v1.JoinGuildRequest\x1a#.protocol.chat.v1.JoinGuildResponse\"\x05\x9aD\x02\b\x01\x12P\n\nLeaveGuild\x12#.protocol.chat.v1.LeaveGuildRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12j\n\rTriggerAction\x12&.protocol.chat.v1.TriggerActionRequest\x1a\x16.google.protobuf.Empty\"\x19\x9aD\x02\b\x01\x9aD\x11\x1a\x0factions.trigger\x12s\n\vSendMessage\x12$.protocol.chat.v1.SendMessageRequest\x1a%.protocol.chat.v1.SendMessageResponse\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send\x12\x88\x01\n\x12QueryHasPermission\x12).protocol.chat.v1.QueryPermissionsRequest\x1a*.protocol.chat.v1.QueryPermissionsResponse\"\x1b\x9aD\x02\b\x01\x9aD\x13\x1a\x11permissions.query\x12s\n\x0eSetPermissions\x12'.protocol.chat.v1.SetPermissionsRequest\x1a\x16.google.protobuf.Empty\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16permissions.manage.set\x12\x85\x01\n\x0eGetPermissions\x12'.protocol.chat.v1.GetPermissionsRequest\x1a(.protocol.chat.v1.GetPermissionsResponse\" \x9aD\x02\b\x01\x9aD\x18\x1a\x16permissions.manage.get\x12i\n\bMoveRole\x12!.protocol.chat.v1.MoveRoleRequest\x1a\".protocol.chat.v1.MoveRoleResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12u\n\rGetGuildRoles\x12&.protocol.chat.v1.GetGuildRolesRequest\x1a'.protocol.chat.v1.GetGuildRolesResponse\"\x13\x9aD\x02\b\x01\x9aD\v\x1a\troles.get\x12u\n\fAddGuildRole\x12%.protocol.chat.v1.AddGuildRoleRequest\x1a&.protocol.chat.v1.AddGuildRoleResponse\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12k\n\x0fModifyGuildRole\x12(.protocol.chat.v1.ModifyGuildRoleRequest\x1a\x16.google.protobuf.Empty\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12k\n\x0fDeleteGuildRole\x12(.protocol.chat.v1.DeleteGuildRoleRequest\x1a\x16.google.protobuf.Empty\"\x16\x9aD\x02\b\x01\x9aD\x0e\x1a\froles.manage\x12p\n\x0fManageUserRoles\x12(.protocol.chat.v1.ManageUserRolesRequest\x1a\x16.google.protobuf.Empty\"\x1b\x9aD\x02\b\x01\x9aD\x13\x1a\x11roles.user.manage\x12d\n\fGetUserRoles\x12%.protocol.chat.v1.GetUserRolesRequest\x1a&.protocol.chat.v1.GetUserRolesResponse\"\x05\x9aD\x02\b\x01\x12Y\n\fStreamEvents\x12%.protocol.chat.v1.StreamEventsRequest\x1a\x17.protocol.chat.v1.Event\"\x05\x9aD\x02\b\x01(\x010\x01\x12F\n\x04Sync\x12\x1d.protocol.chat.v1.SyncRequest\x1a\x1b.protocol.chat.v1.SyncEvent\"\x000\x01\x12U\n\aGetUser\x12 .protocol.chat.v1.GetUserRequest\x1a!.protocol.chat.v1.GetUserResponse\"\x05\x9aD\x02\b\x01\x12m\n\x0fGetUserMetadata\x12(.protocol.chat.v1.GetUserMetadataRequest\x1a).protocol.chat.v1.GetUserMetadataResponse\"\x05\x9aD\x02\b\x01\x12V\n\rProfileUpdate\x12&.protocol.chat.v1.ProfileUpdateRequest\x1a\x16.google.protobuf.Empty\"\x05\x9aD\x02\b\x01\x12Z\n\x06Typing\x12\x1f.protocol.chat.v1.TypingRequest\x1a\x16.google.protobuf.Empty\"\x17\x9aD\x02\b\x01\x9aD\x0f\x1a\rmessages.send\x12d\n\fPreviewGuild\x12%.protocol.chat.v1.PreviewGuildRequest\x1a&.protocol.chat.v1.PreviewGuildResponse\"\x05\x9aD\x02\b\x00B3Z1github.com/harmony-development/legato/gen/chat/v1J\xa8-\n\a\x12\x05\x00\x00\xdf\x01\x01\n\b\n\x01\f\x12\x03\x00\x00\x12\n\t\n\x02\x03\x00\x12\x03\x02\x00%\n.\n\x02\x03\x01\x12\x03\x04\x00%\x1a# import \"validate/validate.proto\";\n\n\t\n\x02\x03\x02\x12\x03\x05\x00\x1f\n\t\n\x02\x03\x03\x12\x03\x06\x00\x1e\n\t\n\x02\x03\x04\x12\x03\a\x00 \n\t\n\x02\x03\x05\x12\x03\b\x00 \n\t\n\x02\x03\x06\x12\x03\t\x00\x1e\n\t\n\x02\x03\a\x12\x03\n\x00#\n\t\n\x02\x03\b\x12\x03\v\x00!\n\t\n\x02\x03\t\x12\x03\f\x00\x1f\n\b\n\x01\x02\x12\x03\x0e\x00\x19\n\b\n\x01\b\x12\x03\x10\x00H\n\t\n\x02\b\v\x12\x03\x10\x00H\n\v\n\x02\x06\x00\x12\x05\x12\x00\xdf\x01\x01\n\n\n\x03\x06\x00\x01\x12\x03\x12\b\x13\n\f\n\x04\x06\x00\x02\x00\x12\x04\x13\x02\x16\x03\n\f\n\x05\x06\x00\x02\x00\x01\x12\x03\x13\x06\x11\n\f\n\x05\x06\x00\x02\x00\x02\x12\x03\x13\x12$\n\f\n\x05\x06\x00\x02\x00\x03\x12\x03\x13.A\n\f\n\x05\x06\x00\x02\x00\x04\x12\x03\x14\x04E\n\x0f\n\b\x06\x00\x02\x00\x04\xc3\b\x01\x12\x03\x14\x04E\n\f\n\x05\x06\x00\x02\x00\x04\x12\x03\x15\x04D\n\x0f\n\b\x06\x00\x02\x00\x04\xc3\b\x03\x12\x03\x15\x04D\n\f\n\x04\x06\x00\x02\x01\x12\x04\x18\x02\x1b\x03\n\f\n\x05\x06\x00\x02\x01\x01\x12\x03\x18\x06\x12\n\f\n\x05\x06\x00\x02\x01\x02\x12\x03\x18\x13&\n\f\n\x05\x06\x00\x02\x01\x03\x12\x03\x180D\n\f\n\x05\x06\x00\x02\x01\x04\x12\x03\x19\x04E\n\x0f\n\b\x06\x00\x02\x01\x04\xc3\b\x01\x12\x03\x19\x04E\n\f\n\x05\x06\x00\x02\x01\x04\x12\x03\x1a\x04Y\n\x0f\n\b\x06\x00\x02\x01\x04\xc3\b\x03\x12\x03\x1a\x04Y\n\f\n\x04\x06\x00\x02\x02\x12\x04\x1d\x02 \x03\n\f\n\x05\x06\x00\x02\x02\x01\x12\x03\x1d\x06\x13\n\f\n\x05\x06\x00\x02\x02\x02\x12\x03\x1d\x14(\n\f\n\x05\x06\x00\x02\x02\x03\x12\x03\x1d2G\n\f\n\x05\x06\x00\x02\x02\x04\x12\x03\x1e\x04E\n\x0f\n\b\x06\x00\x02\x02\x04\xc3\b\x01\x12\x03\x1e\x04E\n\f\n\x05\x06\x00\x02\x02\x04\x12\x03\x1f\x04Z\n\x0f\n\b\x06\x00\x02\x02\x04\xc3\b\x03\x12\x03\x1f\x04Z\n\f\n\x04\x06\x00\x02\x03\x12\x04!\x02#\x03\n\f\n\x05\x06\x00\x02\x03\x01\x12\x03!\x06\x15\n\f\n\x05\x06\x00\x02\x03\x02\x12\x03!\x16,\n\f\n\x05\x06\x00\x02\x03\x03\x12\x03!6M\n\f\n\x05\x06\x00\x02\x03\x04\x12\x03\"\x04E\n\x0f\n\b\x06\x00\x02\x03\x04\xc3\b\x01\x12\x03\"\x04E\n\f\n\x04\x06\x00\x02\x04\x12\x04%\x02'\x03\n\f\n\x05\x06\x00\x02\x04\x01\x12\x03%\x06\x12\n\f\n\x05\x06\x00\x02\x04\x02\x12\x03%\x13&\n\f\n\x05\x06\x00\x02\x04\x03\x12\x03%0D\n\f\n\x05\x06\x00\x02\x04\x04\x12\x03&\x04E\n\x0f\n\b\x06\x00\x02\x04\x04\xc3\b\x01\x12\x03&\x04E\n\f\n\x04\x06\x00\x02\x05\x12\x04(\x02+\x03\n\f\n\x05\x06\x00\x02\x05\x01\x12\x03(\x06\x19\n\f\n\x05\x06\x00\x02\x05\x02\x12\x03(\x1a4\n\f\n\x05\x06\x00\x02\x05\x03\x12\x03(>Y\n\f\n\x05\x06\x00\x02\x05\x04\x12\x03)\x04E\n\x0f\n\b\x06\x00\x02\x05\x04\xc3\b\x01\x12\x03)\x04E\n\f\n\x05\x06\x00\x02\x05\x04\x12\x03*\x04<\n\x0f\n\b\x06\x00\x02\x05\x04\xc3\b\x02\x12\x03*\x04<\n\f\n\x04\x06\x00\x02\x06\x12\x04,\x02/\x03\n\f\n\x05\x06\x00\x02\x06\x01\x12\x03,\x06\x1e\n\f\n\x05\x06\x00\x02\x06\x02\x12\x03,\x1f>\n\f\n\x05\x06\x00\x02\x06\x03\x12\x03,Hh\n\f\n\x05\x06\x00\x02\x06\x04\x12\x03-\x04E\n\x0f\n\b\x06\x00\x02\x06\x04\xc3\b\x01\x12\x03-\x04E\n\f\n\x05\x06\x00\x02\x06\x04\x12\x03.\x04<\n\x0f\n\b\x06\x00\x02\x06\x04\xc3\b\x02\x12\x03.\x04<\n\f\n\x04\x06\x00\x02\a\x12\x041\x023\x03\n\f\n\x05\x06\x00\x02\a\x01\x12\x031\x06\x0e\n\f\n\x05\x06\x00\x02\a\x02\x12\x031\x0f\x1e\n\f\n\x05\x06\x00\x02\a\x03\x12\x031(8\n\f\n\x05\x06\x00\x02\a\x04\x12\x032\x04E\n\x0f\n\b\x06\x00\x02\a\x04\xc3\b\x01\x12\x032\x04E\n<\n\x04\x06\x00\x02\b\x12\x045\x028\x03\x1a. This requires the \"invites.view\" permission.\n\n\f\n\x05\x06\x00\x02\b\x01\x12\x035\x06\x15\n\f\n\x05\x06\x00\x02\b\x02\x12\x035\x16,\n\f\n\x05\x06\x00\x02\b\x03\x12\x0356M\n\f\n\x05\x06\x00\x02\b\x04\x12\x036\x04E\n\x0f\n\b\x06\x00\x02\b\x04\xc3\b\x01\x12\x036\x04E\n\f\n\x05\x06\x00\x02\b\x04\x12\x037\x04P\n\x0f\n\b\x06\x00\x02\b\x04\xc3\b\x03\x12\x037\x04P\n\f\n\x04\x06\x00\x02\t\x12\x049\x02;\x03\n\f\n\x05\x06\x00\x02\t\x01\x12\x039\x06\x15\n\f\n\x05\x06\x00\x02\t\x02\x12\x039\x16,\n\f\n\x05\x06\x00\x02\t\x03\x12\x0396M\n\f\n\x05\x06\x00\x02\t\x04\x12\x03:\x04E\n\x0f\n\b\x06\x00\x02\t\x04\xc3\b\x01\x12\x03:\x04E\na\n\x04\x06\x00\x02\n\x12\x03>\x02T\x1aT You will only be informed of channels you have the \"messages.view\" permission for.\n\n\f\n\x05\x06\x00\x02\n\x01\x12\x03>\x06\x16\n\f\n\x05\x06\x00\x02\n\x02\x12\x03>\x17.\n\f\n\x05\x06\x00\x02\n\x03\x12\x03>8P\n\f\n\x04\x06\x00\x02\v\x12\x04@\x02C\x03\n\f\n\x05\x06\x00\x02\v\x01\x12\x03@\x06\x18\n\f\n\x05\x06\x00\x02\v\x02\x12\x03@\x192\n\f\n\x05\x06\x00\x02\v\x03\x12\x03@<V\n\f\n\x05\x06\x00\x02\v\x04\x12\x03A\x04E\n\x0f\n\b\x06\x00\x02\v\x04\xc3\b\x01\x12\x03A\x04E\n\f\n\x05\x06\x00\x02\v\x04\x12\x03B\x04Q\n\x0f\n\b\x06\x00\x02\v\x04\xc3\b\x03\x12\x03B\x04Q\n\f\n\x04\x06\x00\x02\f\x12\x04E\x02H\x03\n\f\n\x05\x06\x00\x02\f\x01\x12\x03E\x06\x10\n\f\n\x05\x06\x00\x02\f\x02\x12\x03E\x11\"\n\f\n\x05\x06\x00\x02\f\x03\x12\x03E,>\n\f\n\x05\x06\x00\x02\f\x04\x12\x03F\x04E\n\x0f\n\b\x06\x00\x02\f\x04\xc3\b\x01\x12\x03F\x04E\n\f\n\x05\x06\x00\x02\f\x04\x12\x03G\x04Q\n\x0f\n\b\x06\x00\x02\f\x04\xc3\b\x03\x12\x03G\x04Q\n\f\n\x04\x06\x00\x02\r\x12\x04I\x02K\x03\n\f\n\x05\x06\x00\x02\r\x01\x12\x03I\x06\x13\n\f\n\x05\x06\x00\x02\r\x02\x12\x03I\x14(\n\f\n\x05\x06\x00\x02\r\x03\x12\x03I2G\n\f\n\x05\x06\x00\x02\r\x04\x12\x03J\x04E\n\x0f\n\b\x06\x00\x02\r\x04\xc3\b\x01\x12\x03J\x04E\n\f\n\x04\x06\x00\x02\x0e\x12\x04L\x02N\x03\n\f\n\x05\x06\x00\x02\x0e\x01\x12\x03L\x06\x18\n\f\n\x05\x06\x00\x02\x0e\x02\x12\x03L\x192\n\f\n\x05\x06\x00\x02\x0e\x03\x12\x03L<V\n\f\n\x05\x06\x00\x02\x0e\x04\x12\x03M\x04E\n\x0f\n\b\x06\x00\x02\x0e\x04\xc3\b\x01\x12\x03M\x04E\n\f\n\x04\x06\x00\x02\x0f\x12\x04P\x02S\x03\n\f\n\x05\x06\x00\x02\x0f\x01\x12\x03P\x06\x1c\n\f\n\x05\x06\x00\x02\x0f\x02\x12\x03P\x1d:\n\f\n\x05\x06\x00\x02\x0f\x03\x12\x03PDY\n\f\n\x05\x06\x00\x02\x0f\x04\x12\x03Q\x04E\n\x0f\n\b\x06\x00\x02\x0f\x04\xc3\b\x01\x12\x03Q\x04E\n\f\n\x05\x06\x00\x02\x0f\x04\x12\x03R\x04c\n\x0f\n\b\x06\x00\x02\x0f\x04\xc3\b\x03\x12\x03R\x04c\n\f\n\x04\x06\x00\x02\x10\x12\x04T\x02W\x03\n\f\n\x05\x06\x00\x02\x10\x01\x12\x03T\x06\x1e\n\f\n\x05\x06\x00\x02\x10\x02\x12\x03T\x1f>\n\f\n\x05\x06\x00\x02\x10\x03\x12\x03TH]\n\f\n\x05\x06\x00\x02\x10\x04\x12\x03U\x04E\n\x0f\n\b\x06\x00\x02\x10\x04\xc3\b\x01\x12\x03U\x04E\n\f\n\x05\x06\x00\x02\x10\x04\x12\x03V\x04f\n\x0f\n\b\x06\x00\x02\x10\x04\xc3\b\x03\x12\x03V\x04f\n\f\n\x04\x06\x00\x02\x11\x12\x04X\x02[\x03\n\f\n\x05\x06\x00\x02\x11\x01\x12\x03X\x06\x18\n\f\n\x05\x06\x00\x02\x11\x02\x12\x03X\x192\n\f\n\x05\x06\x00\x02\x11\x03\x12\x03X<Q\n\f\n\x05\x06\x00\x02\x11\x04\x12\x03Y\x04E\n\x0f\n\b\x06\x00\x02\x11\x04\xc3\b\x01\x12\x03Y\x04E\n\f\n\x05\x06\x00\x02\x11\x04\x12\x03Z\x04X\n\x0f\n\b\x06\x00\x02\x11\x04\xc3\b\x03\x12\x03Z\x04X\n\f\n\x04\x06\x00\x02\x12\x12\x04\\\x02_\x03\n\f\n\x05\x06\x00\x02\x12\x01\x12\x03\\\x06\x13\n\f\n\x05\x06\x00\x02\x12\x02\x12\x03\\\x14(\n\f\n\x05\x06\x00\x02\x12\x03\x12\x03\\2G\n\f\n\x05\x06\x00\x02\x12\x04\x12\x03]\x04E\n\x0f\n\b\x06\x00\x02\x12\x04\xc3\b\x01\x12\x03]\x04E\n\f\n\x05\x06\x00\x02\x12\x04\x12\x03^\x04Q\n\x0f\n\b\x06\x00\x02\x12\x04\xc3\b\x03\x12\x03^\x04Q\n\f\n\x04\x06\x00\x02\x13\x12\x04`\x02b\x03\n\f\n\x05\x06\x00\x02\x13\x01\x12\x03`\x06\x14\n\f\n\x05\x06\x00\x02\x13\x02\x12\x03`\x15*\n\f\n\x05\x06\x00\x02\x13\x03\x12\x03`4I\n\f\n\x05\x06\x00\x02\x13\x04\x12\x03a\x04E\n\x0f\n\b\x06\x00\x02\x13\x04\xc3\b\x01\x12\x03a\x04E\n\f\n\x04\x06\x00\x02\x14\x12\x04d\x02g\x03\n\f\n\x05\x06\x00\x02\x14\x01\x12\x03d\x06\x11\n\f\n\x05\x06\x00\x02\x14\x02\x12\x03d\x12$\n\f\n\x05\x06\x00\x02\x14\x03\x12\x03d.C\n\f\n\x05\x06\x00\x02\x14\x04\x12\x03e\x04E\n\x0f\n\b\x06\x00\x02\x14\x04\xc3\b\x01\x12\x03e\x04E\n\f\n\x05\x06\x00\x02\x14\x04\x12\x03f\x04W\n\x0f\n\b\x06\x00\x02\x14\x04\xc3\b\x03\x12\x03f\x04W\n\f\n\x04\x06\x00\x02\x15\x12\x04h\x02k\x03\n\f\n\x05\x06\x00\x02\x15\x01\x12\x03h\x06\x12\n\f\n\x05\x06\x00\x02\x15\x02\x12\x03h\x13&\n\f\n\x05\x06\x00\x02\x15\x03\x12\x03h0E\n\f\n\x05\x06\x00\x02\x15\x04\x12\x03i\x04E\n\x0f\n\b\x06\x00\x02\x15\x04\xc3\b\x01\x12\x03i\x04E\n\f\n\x05\x06\x00\x02\x15\x04\x12\x03j\x04Y\n\x0f\n\b\x06\x00\x02\x15\x04\xc3\b\x03\x12\x03j\x04Y\n\f\n\x04\x06\x00\x02\x16\x12\x04l\x02o\x03\n\f\n\x05\x06\x00\x02\x16\x01\x12\x03l\x06\x13\n\f\n\x05\x06\x00\x02\x16\x02\x12\x03l\x14(\n\f\n\x05\x06\x00\x02\x16\x03\x12\x03l2G\n\f\n\x05\x06\x00\x02\x16\x04\x12\x03m\x04E\n\x0f\n\b\x06\x00\x02\x16\x04\xc3\b\x01\x12\x03m\x04E\n\f\n\x05\x06\x00\x02\x16\x04\x12\x03n\x04Z\n\x0f\n\b\x06\x00\x02\x16\x04\xc3\b\x03\x12\x03n\x04Z\ni\n\x04\x06\x00\x02\x17\x12\x04r\x02t\x03\x1a[ This requires the \"messages.manage.delete\" permission if you are not the\n message author.\n\n\f\n\x05\x06\x00\x02\x17\x01\x12\x03r\x06\x13\n\f\n\x05\x06\x00\x02\x17\x02\x12\x03r\x14(\n\f\n\x05\x06\x00\x02\x17\x03\x12\x03r2G\n\f\n\x05\x06\x00\x02\x17\x04\x12\x03s\x04E\n\x0f\n\b\x06\x00\x02\x17\x04\xc3\b\x01\x12\x03s\x04E\n\f\n\x04\x06\x00\x02\x18\x12\x04u\x02w\x03\n\f\n\x05\x06\x00\x02\x18\x01\x12\x03u\x06\x19\n\f\n\x05\x06\x00\x02\x18\x02\x12\x03u\x1a4\n\f\n\x05\x06\x00\x02\x18\x03\x12\x03u>S\n\f\n\x05\x06\x00\x02\x18\x04\x12\x03v\x04E\n\x0f\n\b\x06\x00\x02\x18\x04\xc3\b\x01\x12\x03v\x04E\n\f\n\x04\x06\x00\x02\x19\x12\x04x\x02z\x03\n\f\n\x05\x06\x00\x02\x19\x01\x12\x03x\x06\x15\n\f\n\x05\x06\x00\x02\x19\x02\x12\x03x\x16,\n\f\n\x05\x06\x00\x02\x19\x03\x12\x03x6K\n\f\n\x05\x06\x00\x02\x19\x04\x12\x03y\x04E\n\x0f\n\b\x06\x00\x02\x19\x04\xc3\b\x01\x12\x03y\x04E\n\f\n\x04\x06\x00\x02\x1a\x12\x04{\x02}\x03\n\f\n\x05\x06\x00\x02\x1a\x01\x12\x03{\x06\x15\n\f\n\x05\x06\x00\x02\x1a\x02\x12\x03{\x16,\n\f\n\x05\x06\x00\x02\x1a\x03\x12\x03{6K\n\f\n\x05\x06\x00\x02\x1a\x04\x12\x03|\x04E\n\x0f\n\b\x06\x00\x02\x1a\x04\xc3\b\x01\x12\x03|\x04E\n\r\n\x04\x06\x00\x02\x1b\x12\x05\u007f\x02\x81\x01\x03\n\f\n\x05\x06\x00\x02\x1b\x01\x12\x03\u007f\x06\x0f\n\f\n\x05\x06\x00\x02\x1b\x02\x12\x03\u007f\x10 \n\f\n\x05\x06\x00\x02\x1b\x03\x12\x03\u007f*;\n\r\n\x05\x06\x00\x02\x1b\x04\x12\x04\x80\x01\x04E\n\x10\n\b\x06\x00\x02\x1b\x04\xc3\b\x01\x12\x04\x80\x01\x04E\n\x0e\n\x04\x06\x00\x02\x1c\x12\x06\x82\x01\x02\x84\x01\x03\n\r\n\x05\x06\x00\x02\x1c\x01\x12\x04\x82\x01\x06\x10\n\r\n\x05\x06\x00\x02\x1c\x02\x12\x04\x82\x01\x11\"\n\r\n\x05\x06\x00\x02\x1c\x03\x12\x04\x82\x01,A\n\r\n\x05\x06\x00\x02\x1c\x04\x12\x04\x83\x01\x04E\n\x10\n\b\x06\x00\x02\x1c\x04\xc3\b\x01\x12\x04\x83\x01\x04E\n\x0e\n\x04\x06\x00\x02\x1d\x12\x06\x86\x01\x02\x89\x01\x03\n\r\n\x05\x06\x00\x02\x1d\x01\x12\x04\x86\x01\x06\x13\n\r\n\x05\x06\x00\x02\x1d\x02\x12\x04\x86\x01\x14(\n\r\n\x05\x06\x00\x02\x1d\x03\x12\x04\x86\x012G\n\r\n\x05\x06\x00\x02\x1d\x04\x12\x04\x87\x01\x04E\n\x10\n\b\x06\x00\x02\x1d\x04\xc3\b\x01\x12\x04\x87\x01\x04E\n\r\n\x05\x06\x00\x02\x1d\x04\x12\x04\x88\x01\x04S\n\x10\n\b\x06\x00\x02\x1d\x04\xc3\b\x03\x12\x04\x88\x01\x04S\n\x0e\n\x04\x06\x00\x02\x1e\x12\x06\x8b\x01\x02\x8e\x01\x03\n\r\n\x05\x06\x00\x02\x1e\x01\x12\x04\x8b\x01\x06\x11\n\r\n\x05\x06\x00\x02\x1e\x02\x12\x04\x8b\x01\x12$\n\r\n\x05\x06\x00\x02\x1e\x03\x12\x04\x8b\x01.A\n\r\n\x05\x06\x00\x02\x1e\x04\x12\x04\x8c\x01\x04E\n\x10\n\b\x06\x00\x02\x1e\x04\xc3\b\x01\x12\x04\x8c\x01\x04E\n\r\n\x05\x06\x00\x02\x1e\x04\x12\x04\x8d\x01\x04Q\n\x10\n\b\x06\x00\x02\x1e\x04\xc3\b\x03\x12\x04\x8d\x01\x04Q\n\x0e\n\x04\x06\x00\x02\x1f\x12\x06\x90\x01\x02\x93\x01\x03\n\r\n\x05\x06\x00\x02\x1f\x01\x12\x04\x90\x01\x06\x18\n\r\n\x05\x06\x00\x02\x1f\x02\x12\x04\x90\x01\x190\n\r\n\x05\x06\x00\x02\x1f\x03\x12\x04\x90\x01:R\n\r\n\x05\x06\x00\x02\x1f\x04\x12\x04\x91\x01\x04E\n\x10\n\b\x06\x00\x02\x1f\x04\xc3\b\x01\x12\x04\x91\x01\x04E\n\r\n\x05\x06\x00\x02\x1f\x04\x12\x04\x92\x01\x04U\n\x10\n\b\x06\x00\x02\x1f\x04\xc3\b\x03\x12\x04\x92\x01\x04U\n\x0e\n\x04\x06\x00\x02 \x12\x06\x95\x01\x02\x98\x01\x03\n\r\n\x05\x06\x00\x02 \x01\x12\x04\x95\x01\x06\x14\n\r\n\x05\x06\x00\x02 \x02\x12\x04\x95\x01\x15*\n\r\n\x05\x06\x00\x02 \x03\x12\x04\x95\x014I\n\r\n\x05\x06\x00\x02 \x04\x12\x04\x96\x01\x04E\n\x10\n\b\x06\x00\x02 \x04\xc3\b\x01\x12\x04\x96\x01\x04E\n\r\n\x05\x06\x00\x02 \x04\x12\x04\x97\x01\x04Z\n\x10\n\b\x06\x00\x02 \x04\xc3\b\x03\x12\x04\x97\x01\x04Z\n\x0e\n\x04\x06\x00\x02!\x12\x06\x9a\x01\x02\x9d\x01\x03\n\r\n\x05\x06\x00\x02!\x01\x12\x04\x9a\x01\x06\x14\n\r\n\x05\x06\x00\x02!\x02\x12\x04\x9a\x01\x15*\n\r\n\x05\x06\x00\x02!\x03\x12\x04\x9a\x014J\n\r\n\x05\x06\x00\x02!\x04\x12\x04\x9b\x01\x04E\n\x10\n\b\x06\x00\x02!\x04\xc3\b\x01\x12\x04\x9b\x01\x04E\n\r\n\x05\x06\x00\x02!\x04\x12\x04\x9c\x01\x04Z\n\x10\n\b\x06\x00\x02!\x04\xc3\b\x03\x12\x04\x9c\x01\x04Z\n\x0e\n\x04\x06\x00\x02\"\x12\x06\x9f\x01\x02\xa2\x01\x03\n\r\n\x05\x06\x00\x02\"\x01\x12\x04\x9f\x01\x06\x0e\n\r\n\x05\x06\x00\x02\"\x02\x12\x04\x9f\x01\x0f\x1e\n\r\n\x05\x06\x00\x02\"\x03\x12\x04\x9f\x01(8\n\r\n\x05\x06\x00\x02\"\x04\x12\x04\xa0\x01\x04E\n\x10\n\b\x06\x00\x02\"\x04\xc3\b\x01\x12\x04\xa0\x01\x04E\n\r\n\x05\x06\x00\x02\"\x04\x12\x04\xa1\x01\x04P\n\x10\n\b\x06\x00\x02\"\x04\xc3\b\x03\x12\x04\xa1\x01\x04P\n\x0e\n\x04\x06\x00\x02#\x12\x06\xa4\x01\x02\xa7\x01\x03\n\r\n\x05\x06\x00\x02#\x01\x12\x04\xa4\x01\x06\x13\n\r\n\x05\x06\x00\x02#\x02\x12\x04\xa4\x01\x14(\n\r\n\x05\x06\x00\x02#\x03\x12\x04\xa4\x012G\n\r\n\x05\x06\x00\x02#\x04\x12\x04\xa5\x01\x04E\n\x10\n\b\x06\x00\x02#\x04\xc3\b\x01\x12\x04\xa5\x01\x04E\n\r\n\x05\x06\x00\x02#\x04\x12\x04\xa6\x01\x04M\n\x10\n\b\x06\x00\x02#\x04\xc3\b\x03\x12\x04\xa6\x01\x04M\n\x0e\n\x04\x06\x00\x02$\x12\x06\xa9\x01\x02\xac\x01\x03\n\r\n\x05\x06\x00\x02$\x01\x12\x04\xa9\x01\x06\x12\n\r\n\x05\x06\x00\x02$\x02\x12\x04\xa9\x01\x13&\n\r\n\x05\x06\x00\x02$\x03\x12\x04\xa9\x010D\n\r\n\x05\x06\x00\x02$\x04\x12\x04\xaa\x01\x04E\n\x10\n\b\x06\x00\x02$\x04\xc3\b\x01\x12\x04\xaa\x01\x04E\n\r\n\x05\x06\x00\x02$\x04\x12\x04\xab\x01\x04P\n\x10\n\b\x06\x00\x02$\x04\xc3\b\x03\x12\x04\xab\x01\x04P\n\x0e\n\x04\x06\x00\x02%\x12\x06\xae\x01\x02\xb1\x01\x03\n\r\n\x05\x06\x00\x02%\x01\x12\x04\xae\x01\x06\x15\n\r\n\x05\x06\x00\x02%\x02\x12\x04\xae\x01\x16,\n\r\n\x05\x06\x00\x02%\x03\x12\x04\xae\x016K\n\r\n\x05\x06\x00\x02%\x04\x12\x04\xaf\x01\x04E\n\x10\n\b\x06\x00\x02%\x04\xc3\b\x01\x12\x04\xaf\x01\x04E\n\r\n\x05\x06\x00\x02%\x04\x12\x04\xb0\x01\x04P\n\x10\n\b\x06\x00\x02%\x04\xc3\b\x03\x12\x04\xb0\x01\x04P\n\x0e\n\x04\x06\x00\x02&\x12\x06\xb3\x01\x02\xb6\x01\x03\n\r\n\x05\x06\x00\x02&\x01\x12\x04\xb3\x01\x06\x15\n\r\n\x05\x06\x00\x02&\x02\x12\x04\xb3\x01\x16,\n\r\n\x05\x06\x00\x02&\x03\x12\x04\xb3\x016K\n\r\n\x05\x06\x00\x02&\x04\x12\x04\xb4\x01\x04E\n\x10\n\b\x06\x00\x02&\x04\xc3\b\x01\x12\x04\xb4\x01\x04E\n\r\n\x05\x06\x00\x02&\x04\x12\x04\xb5\x01\x04P\n\x10\n\b\x06\x00\x02&\x04\xc3\b\x03\x12\x04\xb5\x01\x04P\n\x0e\n\x04\x06\x00\x02'\x12\x06\xb8\x01\x02\xbb\x01\x03\n\r\n\x05\x06\x00\x02'\x01\x12\x04\xb8\x01\x06\x15\n\r\n\x05\x06\x00\x02'\x02\x12\x04\xb8\x01\x16,\n\r\n\x05\x06\x00\x02'\x03\x12\x04\xb8\x016K\n\r\n\x05\x06\x00\x02'\x04\x12\x04\xb9\x01\x04E\n\x10\n\b\x06\x00\x02'\x04\xc3\b\x01\x12\x04\xb9\x01\x04E\n\r\n\x05\x06\x00\x02'\x04\x12\x04\xba\x01\x04U\n\x10\n\b\x06\x00\x02'\x04\xc3\b\x03\x12\x04\xba\x01\x04U\n\x0e\n\x04\x06\x00\x02(\x12\x06\xbd\x01\x02\xc3\x01\x03\n\r\n\x05\x06\x00\x02(\x01\x12\x04\xbd\x01\x06\x12\n\r\n\x05\x06\x00\x02(\x02\x12\x04\xbd\x01\x13&\n\r\n\x05\x06\x00\x02(\x03\x12\x04\xbd\x010D\n\r\n\x05\x06\x00\x02(\x04\x12\x04\xbe\x01\x04E\n\xd0\x01\n\b\x06\x00\x02(\x04\xc3\b\x01\x12\x04\xbe\x01\x04E\"\xbd\x01 This permissions node is only required if you are trying to get the roles\n of someone other than yourself.\n\n option (harmonytypes.v1.metadata).requires_permission_node = \"roles.user.get\";\n\n\x0e\n\x04\x06\x00\x02)\x12\x06\xc5\x01\x02\xc7\x01\x03\n\r\n\x05\x06\x00\x02)\x01\x12\x04\xc5\x01\x06\x12\n\r\n\x05\x06\x00\x02)\x05\x12\x04\xc5\x01\x13\x19\n\r\n\x05\x06\x00\x02)\x02\x12\x04\xc5\x01\x1a-\n\r\n\x05\x06\x00\x02)\x06\x12\x04\xc5\x017=\n\r\n\x05\x06\x00\x02)\x03\x12\x04\xc5\x01>C\n\r\n\x05\x06\x00\x02)\x04\x12\x04\xc6\x01\x04E\n\x10\n\b\x06\x00\x02)\x04\xc3\b\x01\x12\x04\xc6\x01\x04E\n\f\n\x04\x06\x00\x02*\x12\x04\xc9\x01\x024\n\r\n\x05\x06\x00\x02*\x01\x12\x04\xc9\x01\x06\n\n\r\n\x05\x06\x00\x02*\x02\x12\x04\xc9\x01\v\x16\n\r\n\x05\x06\x00\x02*\x06\x12\x04\xc9\x01 &\n\r\n\x05\x06\x00\x02*\x03\x12\x04\xc9\x01'0\n\x0e\n\x04\x06\x00\x02+\x12\x06\xcb\x01\x02\xcd\x01\x03\n\r\n\x05\x06\x00\x02+\x01\x12\x04\xcb\x01\x06\r\n\r\n\x05\x06\x00\x02+\x02\x12\x04\xcb\x01\x0e\x1c\n\r\n\x05\x06\x00\x02+\x03\x12\x04\xcb\x01&5\n\r\n\x05\x06\x00\x02+\x04\x12\x04\xcc\x01\x04E\n\x10\n\b\x06\x00\x02+\x04\xc3\b\x01\x12\x04\xcc\x01\x04E\n\x0e\n\x04\x06\x00\x02,\x12\x06\xcf\x01\x02\xd1\x01\x03\n\r\n\x05\x06\x00\x02,\x01\x12\x04\xcf\x01\x06\x15\n\r\n\x05\x06\x00\x02,\x02\x12\x04\xcf\x01\x16,\n\r\n\x05\x06\x00\x02,\x03\x12\x04\xcf\x016M\n\r\n\x05\x06\x00\x02,\x04\x12\x04\xd0\x01\x04E\n\x10\n\b\x06\x00\x02,\x04\xc3\b\x01\x12\x04\xd0\x01\x04E\n\x0e\n\x04\x06\x00\x02-\x12\x06\xd3\x01\x02\xd5\x01\x03\n\r\n\x05\x06\x00\x02-\x01\x12\x04\xd3\x01\x06\x13\n\r\n\x05\x06\x00\x02-\x02\x12\x04\xd3\x01\x14(\n\r\n\x05\x06\x00\x02-\x03\x12\x04\xd3\x012G\n\r\n\x05\x06\x00\x02-\x04\x12\x04\xd4\x01\x04E\n\x10\n\b\x06\x00\x02-\x04\xc3\b\x01\x12\x04\xd4\x01\x04E\n\x0e\n\x04\x06\x00\x02.\x12\x06\xd7\x01\x02\xda\x01\x03\n\r\n\x05\x06\x00\x02.\x01\x12\x04\xd7\x01\x06\f\n\r\n\x05\x06\x00\x02.\x02\x12\x04\xd7\x01\r\x1a\n\r\n\x05\x06\x00\x02.\x03\x12\x04\xd7\x01$9\n\r\n\x05\x06\x00\x02.\x04\x12\x04\xd8\x01\x04E\n\x10\n\b\x06\x00\x02.\x04\xc3\b\x01\x12\x04\xd8\x01\x04E\n\r\n\x05\x06\x00\x02.\x04\x12\x04\xd9\x01\x04Q\n\x10\n\b\x06\x00\x02.\x04\xc3\b\x03\x12\x04\xd9\x01\x04Q\n\x0e\n\x04\x06\x00\x02/\x12\x06\xdc\x01\x02\xde\x01\x03\n\r\n\x05\x06\x00\x02/\x01\x12\x04\xdc\x01\x06\x12\n\r\n\x05\x06\x00\x02/\x02\x12\x04\xdc\x01\x13&\n\r\n\x05\x06\x00\x02/\x03\x12\x04\xdc\x011E\n\r\n\x05\x06\x00\x02/\x04\x12\x04\xdd\x01\x04F\n\x10\n\b\x06\x00\x02/\x04\xc3\b\x01\x12\x04\xdd\x01\x04Fb\x06proto3")

	err := proto.Unmarshal(data, Chatᐳv1ᐳchat)
	if err != nil {
		panic(err)
	}
}

type ChatServiceServer interface {
	CreateGuild(ctx context.Context, r *CreateGuildRequest, headers http.Header) (resp *CreateGuildResponse, err error)

	CreateInvite(ctx context.Context, r *CreateInviteRequest, headers http.Header) (resp *CreateInviteResponse, err error)

	CreateChannel(ctx context.Context, r *CreateChannelRequest, headers http.Header) (resp *CreateChannelResponse, err error)

	CreateEmotePack(ctx context.Context, r *CreateEmotePackRequest, headers http.Header) (resp *CreateEmotePackResponse, err error)

	GetGuildList(ctx context.Context, r *GetGuildListRequest, headers http.Header) (resp *GetGuildListResponse, err error)

	AddGuildToGuildList(ctx context.Context, r *AddGuildToGuildListRequest, headers http.Header) (resp *AddGuildToGuildListResponse, err error)

	RemoveGuildFromGuildList(ctx context.Context, r *RemoveGuildFromGuildListRequest, headers http.Header) (resp *RemoveGuildFromGuildListResponse, err error)

	GetGuild(ctx context.Context, r *GetGuildRequest, headers http.Header) (resp *GetGuildResponse, err error)

	GetGuildInvites(ctx context.Context, r *GetGuildInvitesRequest, headers http.Header) (resp *GetGuildInvitesResponse, err error)

	GetGuildMembers(ctx context.Context, r *GetGuildMembersRequest, headers http.Header) (resp *GetGuildMembersResponse, err error)

	GetGuildChannels(ctx context.Context, r *GetGuildChannelsRequest, headers http.Header) (resp *GetGuildChannelsResponse, err error)

	GetChannelMessages(ctx context.Context, r *GetChannelMessagesRequest, headers http.Header) (resp *GetChannelMessagesResponse, err error)

	GetMessage(ctx context.Context, r *GetMessageRequest, headers http.Header) (resp *GetMessageResponse, err error)

	GetEmotePacks(ctx context.Context, r *GetEmotePacksRequest, headers http.Header) (resp *GetEmotePacksResponse, err error)

	GetEmotePackEmotes(ctx context.Context, r *GetEmotePackEmotesRequest, headers http.Header) (resp *GetEmotePackEmotesResponse, err error)

	UpdateGuildInformation(ctx context.Context, r *UpdateGuildInformationRequest, headers http.Header) (resp *empty.Empty, err error)

	UpdateChannelInformation(ctx context.Context, r *UpdateChannelInformationRequest, headers http.Header) (resp *empty.Empty, err error)

	UpdateChannelOrder(ctx context.Context, r *UpdateChannelOrderRequest, headers http.Header) (resp *empty.Empty, err error)

	UpdateMessage(ctx context.Context, r *UpdateMessageRequest, headers http.Header) (resp *empty.Empty, err error)

	AddEmoteToPack(ctx context.Context, r *AddEmoteToPackRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteGuild(ctx context.Context, r *DeleteGuildRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteInvite(ctx context.Context, r *DeleteInviteRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteChannel(ctx context.Context, r *DeleteChannelRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteMessage(ctx context.Context, r *DeleteMessageRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteEmoteFromPack(ctx context.Context, r *DeleteEmoteFromPackRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteEmotePack(ctx context.Context, r *DeleteEmotePackRequest, headers http.Header) (resp *empty.Empty, err error)

	DequipEmotePack(ctx context.Context, r *DequipEmotePackRequest, headers http.Header) (resp *empty.Empty, err error)

	JoinGuild(ctx context.Context, r *JoinGuildRequest, headers http.Header) (resp *JoinGuildResponse, err error)

	LeaveGuild(ctx context.Context, r *LeaveGuildRequest, headers http.Header) (resp *empty.Empty, err error)

	TriggerAction(ctx context.Context, r *TriggerActionRequest, headers http.Header) (resp *empty.Empty, err error)

	SendMessage(ctx context.Context, r *SendMessageRequest, headers http.Header) (resp *SendMessageResponse, err error)

	QueryHasPermission(ctx context.Context, r *QueryPermissionsRequest, headers http.Header) (resp *QueryPermissionsResponse, err error)

	SetPermissions(ctx context.Context, r *SetPermissionsRequest, headers http.Header) (resp *empty.Empty, err error)

	GetPermissions(ctx context.Context, r *GetPermissionsRequest, headers http.Header) (resp *GetPermissionsResponse, err error)

	MoveRole(ctx context.Context, r *MoveRoleRequest, headers http.Header) (resp *MoveRoleResponse, err error)

	GetGuildRoles(ctx context.Context, r *GetGuildRolesRequest, headers http.Header) (resp *GetGuildRolesResponse, err error)

	AddGuildRole(ctx context.Context, r *AddGuildRoleRequest, headers http.Header) (resp *AddGuildRoleResponse, err error)

	ModifyGuildRole(ctx context.Context, r *ModifyGuildRoleRequest, headers http.Header) (resp *empty.Empty, err error)

	DeleteGuildRole(ctx context.Context, r *DeleteGuildRoleRequest, headers http.Header) (resp *empty.Empty, err error)

	ManageUserRoles(ctx context.Context, r *ManageUserRolesRequest, headers http.Header) (resp *empty.Empty, err error)

	GetUserRoles(ctx context.Context, r *GetUserRolesRequest, headers http.Header) (resp *GetUserRolesResponse, err error)

	StreamEvents(ctx context.Context, in chan *StreamEventsRequest, out chan *Event, headers http.Header)

	Sync(ctx context.Context, r *SyncRequest, out chan *SyncEvent, headers http.Header)

	GetUser(ctx context.Context, r *GetUserRequest, headers http.Header) (resp *GetUserResponse, err error)

	GetUserMetadata(ctx context.Context, r *GetUserMetadataRequest, headers http.Header) (resp *GetUserMetadataResponse, err error)

	ProfileUpdate(ctx context.Context, r *ProfileUpdateRequest, headers http.Header) (resp *empty.Empty, err error)

	Typing(ctx context.Context, r *TypingRequest, headers http.Header) (resp *empty.Empty, err error)

	PreviewGuild(ctx context.Context, r *PreviewGuildRequest, headers http.Header) (resp *PreviewGuildResponse, err error)
}

type ChatServiceHandler struct {
	Server       ChatServiceServer
	ErrorHandler func(err error, w http.ResponseWriter)
	upgrader     websocket.Upgrader
}

func NewChatServiceHandler(s ChatServiceServer, errHandler func(err error, w http.ResponseWriter)) *ChatServiceHandler {
	return &ChatServiceHandler{
		Server:       s,
		ErrorHandler: errHandler,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (h *ChatServiceHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {

	case "/protocol.chat.v1.ChatService/CreateGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(CreateGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.CreateGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/CreateInvite":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(CreateInviteRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.CreateInvite(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/CreateChannel":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(CreateChannelRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.CreateChannel(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/CreateEmotePack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(CreateEmotePackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.CreateEmotePack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildList":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildListRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildList(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/AddGuildToGuildList":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(AddGuildToGuildListRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.AddGuildToGuildList(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/RemoveGuildFromGuildList":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(RemoveGuildFromGuildListRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.RemoveGuildFromGuildList(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildInvites":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildInvitesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildInvites(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildMembers":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildMembersRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildMembers(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildChannels":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildChannelsRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildChannels(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetChannelMessages":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetChannelMessagesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetChannelMessages(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetMessage":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetMessageRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetMessage(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetEmotePacks":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetEmotePacksRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetEmotePacks(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetEmotePackEmotes":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetEmotePackEmotesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetEmotePackEmotes(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/UpdateGuildInformation":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(UpdateGuildInformationRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.UpdateGuildInformation(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/UpdateChannelInformation":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(UpdateChannelInformationRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.UpdateChannelInformation(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/UpdateChannelOrder":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(UpdateChannelOrderRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.UpdateChannelOrder(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/UpdateMessage":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(UpdateMessageRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.UpdateMessage(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/AddEmoteToPack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(AddEmoteToPackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.AddEmoteToPack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteInvite":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteInviteRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteInvite(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteChannel":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteChannelRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteChannel(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteMessage":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteMessageRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteMessage(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteEmoteFromPack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteEmoteFromPackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteEmoteFromPack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteEmotePack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteEmotePackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteEmotePack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DequipEmotePack":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DequipEmotePackRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DequipEmotePack(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/JoinGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(JoinGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.JoinGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/LeaveGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(LeaveGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.LeaveGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/TriggerAction":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(TriggerActionRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.TriggerAction(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/SendMessage":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(SendMessageRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.SendMessage(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/QueryHasPermission":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(QueryPermissionsRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.QueryHasPermission(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/SetPermissions":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(SetPermissionsRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.SetPermissions(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetPermissions":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetPermissionsRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetPermissions(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/MoveRole":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(MoveRoleRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.MoveRole(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetGuildRoles":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetGuildRolesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetGuildRoles(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/AddGuildRole":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(AddGuildRoleRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.AddGuildRole(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/ModifyGuildRole":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(ModifyGuildRoleRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.ModifyGuildRole(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/DeleteGuildRole":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(DeleteGuildRoleRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.DeleteGuildRole(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/ManageUserRoles":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(ManageUserRolesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.ManageUserRoles(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetUserRoles":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetUserRolesRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetUserRoles(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/StreamEvents":
		{
			var err error

			in := make(chan *StreamEventsRequest)
			err = nil

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			out := make(chan *Event)

			ws, err := h.upgrader.Upgrade(w, req, nil)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			go func() {

				msgs := make(chan []byte)

				go func() {
					for {
						_, message, err := ws.ReadMessage()
						if err != nil {
							close(msgs)
							break
						}
						msgs <- message
					}
				}()

				defer ws.WriteMessage(websocket.CloseMessage, []byte{})

				for {
					select {

					case data, ok := <-msgs:
						if !ok {
							return
						}

						item := new(StreamEventsRequest)
						err = proto.Unmarshal(data, item)
						if err != nil {
							close(in)
							close(out)
							return
						}

						in <- item

					case msg, ok := <-out:
						if !ok {
							return
						}

						w, err := ws.NextWriter(websocket.BinaryMessage)
						if err != nil {

							close(in)

							close(out)
							return
						}

						response, err := proto.Marshal(msg)
						if err != nil {

							close(in)

							close(out)
							return
						}

						w.Write(response)
						if err := w.Close(); err != nil {

							close(in)

							close(out)
							return
						}
					}
				}
			}()

			h.Server.StreamEvents(req.Context(), in, out, req.Header)
		}

	case "/protocol.chat.v1.ChatService/Sync":
		{
			var err error

			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			in := new(SyncRequest)
			err = proto.Unmarshal(body, in)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			out := make(chan *SyncEvent)

			ws, err := h.upgrader.Upgrade(w, req, nil)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			go func() {

				defer ws.WriteMessage(websocket.CloseMessage, []byte{})

				for {
					select {

					case msg, ok := <-out:
						if !ok {
							return
						}

						w, err := ws.NextWriter(websocket.BinaryMessage)
						if err != nil {

							close(out)
							return
						}

						response, err := proto.Marshal(msg)
						if err != nil {

							close(out)
							return
						}

						w.Write(response)
						if err := w.Close(); err != nil {

							close(out)
							return
						}
					}
				}
			}()

			h.Server.Sync(req.Context(), in, out, req.Header)
		}

	case "/protocol.chat.v1.ChatService/GetUser":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetUserRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetUser(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/GetUserMetadata":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(GetUserMetadataRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.GetUserMetadata(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/ProfileUpdate":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(ProfileUpdateRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.ProfileUpdate(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/Typing":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(TypingRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.Typing(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	case "/protocol.chat.v1.ChatService/PreviewGuild":
		{
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			requestProto := new(PreviewGuildRequest)
			err = proto.Unmarshal(body, requestProto)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			resp, err := h.Server.PreviewGuild(req.Context(), requestProto, req.Header)

			response, err := proto.Marshal(resp)
			if err != nil {
				h.ErrorHandler(err, w)
				return
			}

			w.Header().Add("Content-Type", "application/octet-stream")
			_, err = w.Write(response)

			if err != nil {
				h.ErrorHandler(err, w)
				return
			}
		}

	}
}
