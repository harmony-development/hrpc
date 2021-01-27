#include <QByteArray>
#include <QCoreApplication>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QString>
#include <variant>
// #include <QtWebSockets>
#include "google/protobuf/empty.pb.h"
class VoiceServiceServiceClient {
	QString host;
	bool secure;
	QSharedPointer<QNetworkAccessManager> nam;
	QString httpProtocol() const { return secure ? QStringLiteral("https://") : QStringLiteral("http://"); }
	explicit VoiceServiceServiceClient(const QString& host, bool secure) : host(host), secure(secure), nam(new QNetworkAccessManager) {}
public:
	template<typename T> using Result = std::variant<T, QString>;
// todo client <-> server stream
};
