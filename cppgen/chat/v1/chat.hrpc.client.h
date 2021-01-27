#include <QByteArray>
#include <QCoreApplication>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QString>
#include <variant>
#include <QWebSocket>
#include "google/protobuf/empty.pb.h"
#include "harmonytypes/v1/types.pb.h"
#include "chat/v1/profile.pb.h"
#include "chat/v1/guilds.pb.h"
#include "chat/v1/channels.pb.h"
#include "chat/v1/messages.pb.h"
#include "chat/v1/emotes.pb.h"
#include "chat/v1/permissions.pb.h"
#include "chat/v1/streaming.pb.h"
#include "chat/v1/postbox.pb.h"

class Receive__protocol_chat_v1_Event__Send__protocol_chat_v1_StreamEventsRequest__Stream : public QWebSocket {
	
	Q_OBJECT

	public: Q_SIGNAL void receivedMessage(protocol::chat::v1::Event msg);

	public: Receive__protocol_chat_v1_Event__Send__protocol_chat_v1_StreamEventsRequest__Stream(const QString &origin = QString(), QWebSocketProtocol::Version version = QWebSocketProtocol::VersionLatest, QObject *parent = nullptr) : QWebSocket(origin, version, parent)
	{
		connect(this, &QWebSocket::binaryMessageReceived, [=](const QByteArray& msg) {
			protocol::chat::v1::Event incoming;

			if (!incoming.ParseFromArray(msg.constData(), msg.length())) {
				return;
			}

			Q_EMIT receivedMessage(incoming);
		});
	}


	public: bool send(const protocol::chat::v1::StreamEventsRequest& in) {
		QByteArray data = QByteArray::fromStdString(in.SerializeAsString());
		if (data.length() == 0) {
			return false;
		}

		auto count = sendBinaryMessage(data);
		return count == data.length();
	}
};

class Receive__protocol_chat_v1_SyncEvent__Stream : public QWebSocket {
	
	Q_OBJECT

	public: Q_SIGNAL void receivedMessage(protocol::chat::v1::SyncEvent msg);

	public: Receive__protocol_chat_v1_SyncEvent__Stream(const QString &origin = QString(), QWebSocketProtocol::Version version = QWebSocketProtocol::VersionLatest, QObject *parent = nullptr) : QWebSocket(origin, version, parent)
	{
		connect(this, &QWebSocket::binaryMessageReceived, [=](const QByteArray& msg) {
			protocol::chat::v1::SyncEvent incoming;

			if (!incoming.ParseFromArray(msg.constData(), msg.length())) {
				return;
			}

			Q_EMIT receivedMessage(incoming);
		});
	}

};

class ChatServiceServiceClient {
	QString host;
	bool secure;
	QSharedPointer<QNetworkAccessManager> nam;
	QString httpProtocol() const { return secure ? QStringLiteral("https://") : QStringLiteral("http://"); }
	QString wsProtocol() const { return secure ? QStringLiteral("wss://") : QStringLiteral("ws://"); }
	explicit ChatServiceServiceClient(const QString& host, bool secure) : host(host), secure(secure), nam(new QNetworkAccessManager) {}
public:
	template<typename T> using Result = std::variant<T, QString>;
	Result<protocol::chat::v1::CreateGuildResponse> CreateGuild(const protocol::chat::v1::CreateGuildRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::CreateInviteResponse> CreateInvite(const protocol::chat::v1::CreateInviteRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::CreateChannelResponse> CreateChannel(const protocol::chat::v1::CreateChannelRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::CreateEmotePackResponse> CreateEmotePack(const protocol::chat::v1::CreateEmotePackRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetGuildListResponse> GetGuildList(const protocol::chat::v1::GetGuildListRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::AddGuildToGuildListResponse> AddGuildToGuildList(const protocol::chat::v1::AddGuildToGuildListRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::RemoveGuildFromGuildListResponse> RemoveGuildFromGuildList(const protocol::chat::v1::RemoveGuildFromGuildListRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetGuildResponse> GetGuild(const protocol::chat::v1::GetGuildRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetGuildInvitesResponse> GetGuildInvites(const protocol::chat::v1::GetGuildInvitesRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetGuildMembersResponse> GetGuildMembers(const protocol::chat::v1::GetGuildMembersRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetGuildChannelsResponse> GetGuildChannels(const protocol::chat::v1::GetGuildChannelsRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetChannelMessagesResponse> GetChannelMessages(const protocol::chat::v1::GetChannelMessagesRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetMessageResponse> GetMessage(const protocol::chat::v1::GetMessageRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetEmotePacksResponse> GetEmotePacks(const protocol::chat::v1::GetEmotePacksRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetEmotePackEmotesResponse> GetEmotePackEmotes(const protocol::chat::v1::GetEmotePackEmotesRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> UpdateGuildInformation(const protocol::chat::v1::UpdateGuildInformationRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> UpdateChannelInformation(const protocol::chat::v1::UpdateChannelInformationRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> UpdateChannelOrder(const protocol::chat::v1::UpdateChannelOrderRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> UpdateMessage(const protocol::chat::v1::UpdateMessageRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> AddEmoteToPack(const protocol::chat::v1::AddEmoteToPackRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> DeleteGuild(const protocol::chat::v1::DeleteGuildRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> DeleteInvite(const protocol::chat::v1::DeleteInviteRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> DeleteChannel(const protocol::chat::v1::DeleteChannelRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> DeleteMessage(const protocol::chat::v1::DeleteMessageRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> DeleteEmoteFromPack(const protocol::chat::v1::DeleteEmoteFromPackRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> DeleteEmotePack(const protocol::chat::v1::DeleteEmotePackRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> DequipEmotePack(const protocol::chat::v1::DequipEmotePackRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::JoinGuildResponse> JoinGuild(const protocol::chat::v1::JoinGuildRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> LeaveGuild(const protocol::chat::v1::LeaveGuildRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> TriggerAction(const protocol::chat::v1::TriggerActionRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::SendMessageResponse> SendMessage(const protocol::chat::v1::SendMessageRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::QueryPermissionsResponse> QueryHasPermission(const protocol::chat::v1::QueryPermissionsRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> SetPermissions(const protocol::chat::v1::SetPermissionsRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetPermissionsResponse> GetPermissions(const protocol::chat::v1::GetPermissionsRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::MoveRoleResponse> MoveRole(const protocol::chat::v1::MoveRoleRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetGuildRolesResponse> GetGuildRoles(const protocol::chat::v1::GetGuildRolesRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::AddGuildRoleResponse> AddGuildRole(const protocol::chat::v1::AddGuildRoleRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> ModifyGuildRole(const protocol::chat::v1::ModifyGuildRoleRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> DeleteGuildRole(const protocol::chat::v1::DeleteGuildRoleRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> ManageUserRoles(const protocol::chat::v1::ManageUserRolesRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetUserRolesResponse> GetUserRoles(const protocol::chat::v1::GetUserRolesRequest& in, QMap<QByteArray,QString> headers = {});
	Receive__protocol_chat_v1_Event__Send__protocol_chat_v1_StreamEventsRequest__Stream* StreamEvents();
// todo client <- server stream
	Result<protocol::chat::v1::GetUserResponse> GetUser(const protocol::chat::v1::GetUserRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::GetUserMetadataResponse> GetUserMetadata(const protocol::chat::v1::GetUserMetadataRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> ProfileUpdate(const protocol::chat::v1::ProfileUpdateRequest& in, QMap<QByteArray,QString> headers = {});
	Result<google::protobuf::Empty> Typing(const protocol::chat::v1::TypingRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::chat::v1::PreviewGuildResponse> PreviewGuild(const protocol::chat::v1::PreviewGuildRequest& in, QMap<QByteArray,QString> headers = {});
};
