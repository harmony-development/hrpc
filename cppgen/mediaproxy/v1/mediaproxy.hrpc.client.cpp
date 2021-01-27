#include "mediaproxy/v1/mediaproxy.hrpc.client.h"
auto MediaProxyServiceServiceClient::FetchLinkMetadata(const protocol::mediaproxy::v1::FetchLinkMetadataRequest& in, QMap<QByteArray,QString> headers) -> MediaProxyServiceServiceClient::Result<protocol::mediaproxy::v1::SiteMetadata>
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

	protocol::mediaproxy::v1::SiteMetadata ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto MediaProxyServiceServiceClient::InstantView(const protocol::mediaproxy::v1::InstantViewRequest& in, QMap<QByteArray,QString> headers) -> MediaProxyServiceServiceClient::Result<protocol::mediaproxy::v1::InstantViewResponse>
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

	protocol::mediaproxy::v1::InstantViewResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
auto MediaProxyServiceServiceClient::CanInstantView(const protocol::mediaproxy::v1::InstantViewRequest& in, QMap<QByteArray,QString> headers) -> MediaProxyServiceServiceClient::Result<protocol::mediaproxy::v1::CanInstantViewResponse>
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

	protocol::mediaproxy::v1::CanInstantViewResponse ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};

}
