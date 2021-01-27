#include <QByteArray>
#include <QCoreApplication>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QString>
#include <variant>
#include "harmonytypes/v1/types.pb.h"

class MediaProxyServiceServiceClient {
	QString host;
	bool secure;
	QSharedPointer<QNetworkAccessManager> nam;
	QString httpProtocol() const { return secure ? QStringLiteral("https://") : QStringLiteral("http://"); }
	QString wsProtocol() const { return secure ? QStringLiteral("wss://") : QStringLiteral("ws://"); }
	explicit MediaProxyServiceServiceClient(const QString& host, bool secure) : host(host), secure(secure), nam(new QNetworkAccessManager) {}
public:
	template<typename T> using Result = std::variant<T, QString>;
	Result<protocol::mediaproxy::v1::SiteMetadata> FetchLinkMetadata(const protocol::mediaproxy::v1::FetchLinkMetadataRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::mediaproxy::v1::InstantViewResponse> InstantView(const protocol::mediaproxy::v1::InstantViewRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::mediaproxy::v1::CanInstantViewResponse> CanInstantView(const protocol::mediaproxy::v1::InstantViewRequest& in, QMap<QByteArray,QString> headers = {});
};
