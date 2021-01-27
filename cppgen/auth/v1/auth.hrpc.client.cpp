#include "auth/v1/auth.hrpc.client.h"
auto AuthServiceServiceClient::Federate(const protocol::auth::v1::FederateRequest& in, QMap<QByteArray,QString> headers) -> AuthServiceServiceClient::Result<protocol::auth::v1::FederateReply>
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

	protocol::auth::v1::FederateReply ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto AuthServiceServiceClient::LoginFederated(const protocol::auth::v1::LoginFederatedRequest& in, QMap<QByteArray,QString> headers) -> AuthServiceServiceClient::Result<protocol::auth::v1::Session>
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

	protocol::auth::v1::Session ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto AuthServiceServiceClient::Key(const google::protobuf::Empty& in, QMap<QByteArray,QString> headers) -> AuthServiceServiceClient::Result<protocol::auth::v1::KeyReply>
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

	protocol::auth::v1::KeyReply ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto AuthServiceServiceClient::BeginAuth(const google::protobuf::Empty& in, QMap<QByteArray,QString> headers) -> AuthServiceServiceClient::Result<protocol::auth::v1::BeginAuthResponse>
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

	protocol::auth::v1::BeginAuthResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto AuthServiceServiceClient::NextStep(const protocol::auth::v1::NextStepRequest& in, QMap<QByteArray,QString> headers) -> AuthServiceServiceClient::Result<protocol::auth::v1::AuthStep>
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

	protocol::auth::v1::AuthStep ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto AuthServiceServiceClient::StepBack(const protocol::auth::v1::StepBackRequest& in, QMap<QByteArray,QString> headers) -> AuthServiceServiceClient::Result<protocol::auth::v1::AuthStep>
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

	protocol::auth::v1::AuthStep ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
// todo client <- server stream
