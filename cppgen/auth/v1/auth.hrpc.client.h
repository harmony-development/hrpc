#include <QByteArray>
#include <QCoreApplication>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QString>
#include <variant>
#include <QWebSocket>
#include "google/protobuf/empty.pb.h"

class Receive__protocol_auth_v1_AuthStep__Stream : public QWebSocket {
	
	Q_OBJECT

	public: Q_SIGNAL void receivedMessage(protocol::auth::v1::AuthStep msg);

	public: Receive__protocol_auth_v1_AuthStep__Stream(const QString &origin = QString(), QWebSocketProtocol::Version version = QWebSocketProtocol::VersionLatest, QObject *parent = nullptr) : QWebSocket(origin, version, parent)
	{
		connect(this, &QWebSocket::binaryMessageReceived, [=](const QByteArray& msg) {
			protocol::auth::v1::AuthStep incoming;

			if (!incoming.ParseFromArray(msg.constData(), msg.length())) {
				return;
			}

			Q_EMIT receivedMessage(incoming);
		});
	}

};

class AuthServiceServiceClient {
	QString host;
	bool secure;
	QSharedPointer<QNetworkAccessManager> nam;
	QString httpProtocol() const { return secure ? QStringLiteral("https://") : QStringLiteral("http://"); }
	QString wsProtocol() const { return secure ? QStringLiteral("wss://") : QStringLiteral("ws://"); }
	explicit AuthServiceServiceClient(const QString& host, bool secure) : host(host), secure(secure), nam(new QNetworkAccessManager) {}
public:
	template<typename T> using Result = std::variant<T, QString>;
	Result<protocol::auth::v1::FederateReply> Federate(const protocol::auth::v1::FederateRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::auth::v1::Session> LoginFederated(const protocol::auth::v1::LoginFederatedRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::auth::v1::KeyReply> Key(const google::protobuf::Empty& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::auth::v1::BeginAuthResponse> BeginAuth(const google::protobuf::Empty& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::auth::v1::AuthStep> NextStep(const protocol::auth::v1::NextStepRequest& in, QMap<QByteArray,QString> headers = {});
	Result<protocol::auth::v1::AuthStep> StepBack(const protocol::auth::v1::StepBackRequest& in, QMap<QByteArray,QString> headers = {});
// todo client <- server stream
};
