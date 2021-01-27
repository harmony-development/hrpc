#include "chat/v1/chat.hrpc.client.h"
auto ChatServiceServiceClient::CreateGuild(const protocol::chat::v1::CreateGuildRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::CreateGuildResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::CreateGuildResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::CreateInvite(const protocol::chat::v1::CreateInviteRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::CreateInviteResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::CreateInviteResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::CreateChannel(const protocol::chat::v1::CreateChannelRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::CreateChannelResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::CreateChannelResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::CreateEmotePack(const protocol::chat::v1::CreateEmotePackRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::CreateEmotePackResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::CreateEmotePackResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetGuildList(const protocol::chat::v1::GetGuildListRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetGuildListResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetGuildListResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::AddGuildToGuildList(const protocol::chat::v1::AddGuildToGuildListRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::AddGuildToGuildListResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::AddGuildToGuildListResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::RemoveGuildFromGuildList(const protocol::chat::v1::RemoveGuildFromGuildListRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::RemoveGuildFromGuildListResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::RemoveGuildFromGuildListResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetGuild(const protocol::chat::v1::GetGuildRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetGuildResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetGuildResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetGuildInvites(const protocol::chat::v1::GetGuildInvitesRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetGuildInvitesResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetGuildInvitesResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetGuildMembers(const protocol::chat::v1::GetGuildMembersRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetGuildMembersResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetGuildMembersResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetGuildChannels(const protocol::chat::v1::GetGuildChannelsRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetGuildChannelsResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetGuildChannelsResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetChannelMessages(const protocol::chat::v1::GetChannelMessagesRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetChannelMessagesResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetChannelMessagesResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetMessage(const protocol::chat::v1::GetMessageRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetMessageResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetMessageResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetEmotePacks(const protocol::chat::v1::GetEmotePacksRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetEmotePacksResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetEmotePacksResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetEmotePackEmotes(const protocol::chat::v1::GetEmotePackEmotesRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetEmotePackEmotesResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetEmotePackEmotesResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::UpdateGuildInformation(const protocol::chat::v1::UpdateGuildInformationRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::UpdateChannelInformation(const protocol::chat::v1::UpdateChannelInformationRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::UpdateChannelOrder(const protocol::chat::v1::UpdateChannelOrderRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::UpdateMessage(const protocol::chat::v1::UpdateMessageRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::AddEmoteToPack(const protocol::chat::v1::AddEmoteToPackRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::DeleteGuild(const protocol::chat::v1::DeleteGuildRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::DeleteInvite(const protocol::chat::v1::DeleteInviteRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::DeleteChannel(const protocol::chat::v1::DeleteChannelRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::DeleteMessage(const protocol::chat::v1::DeleteMessageRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::DeleteEmoteFromPack(const protocol::chat::v1::DeleteEmoteFromPackRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::DeleteEmotePack(const protocol::chat::v1::DeleteEmotePackRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::DequipEmotePack(const protocol::chat::v1::DequipEmotePackRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::JoinGuild(const protocol::chat::v1::JoinGuildRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::JoinGuildResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::JoinGuildResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::LeaveGuild(const protocol::chat::v1::LeaveGuildRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::TriggerAction(const protocol::chat::v1::TriggerActionRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::SendMessage(const protocol::chat::v1::SendMessageRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::SendMessageResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::SendMessageResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::QueryHasPermission(const protocol::chat::v1::QueryPermissionsRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::QueryPermissionsResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::QueryPermissionsResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::SetPermissions(const protocol::chat::v1::SetPermissionsRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetPermissions(const protocol::chat::v1::GetPermissionsRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetPermissionsResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetPermissionsResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::MoveRole(const protocol::chat::v1::MoveRoleRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::MoveRoleResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::MoveRoleResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetGuildRoles(const protocol::chat::v1::GetGuildRolesRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetGuildRolesResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetGuildRolesResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::AddGuildRole(const protocol::chat::v1::AddGuildRoleRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::AddGuildRoleResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::AddGuildRoleResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::ModifyGuildRole(const protocol::chat::v1::ModifyGuildRoleRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::DeleteGuildRole(const protocol::chat::v1::DeleteGuildRoleRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::ManageUserRoles(const protocol::chat::v1::ManageUserRolesRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetUserRoles(const protocol::chat::v1::GetUserRolesRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetUserRolesResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetUserRolesResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::StreamEvents() -> Receive__protocol_chat_v1_Event__Send__protocol_chat_v1_StreamEventsRequest__Stream*
{
	auto sock = new Receive__protocol_chat_v1_Event__Send__protocol_chat_v1_StreamEventsRequest__Stream();
	sock->open(QUrl(wsProtocol()+host));
	return sock;
}
// todo client <- server stream
auto ChatServiceServiceClient::GetUser(const protocol::chat::v1::GetUserRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetUserResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetUserResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::GetUserMetadata(const protocol::chat::v1::GetUserMetadataRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::GetUserMetadataResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::GetUserMetadataResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::ProfileUpdate(const protocol::chat::v1::ProfileUpdateRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::Typing(const protocol::chat::v1::TypingRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<google::protobuf::Empty>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	google::protobuf::Empty ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto ChatServiceServiceClient::PreviewGuild(const protocol::chat::v1::PreviewGuildRequest& in, QMap<QByteArray,QString> headers) -> ChatServiceServiceClient::Result<protocol::chat::v1::PreviewGuildResponse>
{
	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());

	if (data.length() == 0) {
		return {QStringLiteral("failed to serialize protobuf")};
	}

	QUrl serviceURL = QUrl(httpProtocol()+host);

	QNetworkRequest req(serviceURL);
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}

	auto val = nam->post(req, data);

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	if (val->error() == QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%1): %2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	protocol::chat::v1::PreviewGuildResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
