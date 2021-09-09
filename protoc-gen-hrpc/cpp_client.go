package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// a/b/c.proto -> a/b/c.(interfix).(suffix)
func convertCxxProto(protoName, interfix, suffix string) string {
	return fmt.Sprintf(`%s.%s.%s`, protoName[:len(protoName)-6], interfix, suffix)
}

func inc(s string) string {
	return fmt.Sprintf(`#include "%s"`, s)
}

func getImports(d *descriptorpb.FileDescriptorProto, mu []*descriptorpb.FileDescriptorProto, isHrpcTypes bool) string {
	set := []string{}
	add := func(s string) {
		for _, item := range set {
			if item == s {
				return
			}
		}
		set = append(set, s)
	}

	add(`#include <QByteArray>`)
	add(`#include <QCoreApplication>`)
	add(`#include <QNetworkAccessManager>`)
	add(`#include <QNetworkReply>`)
	add(`#include <QString>`)
	add(`#include <QFuture>`)
	add(`#include <functional>`)
	add(`#include <variant>`)

	add(`#include <future.h>`)

	add(inc(convertCxxProto(*d.Name, "pb", "h")))
	if len(d.MessageType) > 0 {
		if !isHrpcTypes {
			add(inc(convertCxxProto(*d.Name, "hrpc.types", "h")))
		}
	}

	for _, kind := range d.Service {
		for _, meth := range kind.Method {
			if meth.GetClientStreaming() || meth.GetServerStreaming() {
				add(`#include <QWebSocket>`)
			}
		}
	}

	for _, f := range mu {
		for _, dep := range f.Dependency {
			add(inc(convertCxxProto(dep, "pb", "h")))
			if !isHrpcTypes {
				add(inc(convertCxxProto(dep, "hrpc.types", "h")))
			}
		}
	}

	return strings.Join(set, "\n") + "\n"
}

func typeToCxxNamespaces(s string) string {
	return strings.ReplaceAll(s[1:], ".", "::")
}

func localTypeToCxxNamespaces(s string) string {
	return strings.ReplaceAll(s, ".", "::")
}

type sPair struct {
	fromServ string
	fromClie string
}

func generateClientSockets(d *descriptorpb.FileDescriptorProto) string {
	sb := strings.Builder{}

	add := func(s string) { sb.WriteString(s + "\n") }

	bidiPairs := []sPair{}
	unidiPairs := []sPair{}

	addPair := func(fromS, fromC string, to *[]sPair) {
		for _, item := range *to {
			if item.fromClie == fromC && item.fromServ == fromS {
				continue
			}
		}

		*to = append(*to, sPair{fromServ: fromS, fromClie: fromC})
	}

	sane := func(s string) string { return strings.ReplaceAll(typeToCxxNamespaces(s), "::", "_") }

	for _, service := range d.Service {
		for _, meth := range service.Method {
			if meth.GetClientStreaming() && !meth.GetServerStreaming() {
				continue
			} else if meth.GetClientStreaming() && meth.GetServerStreaming() {
				addPair(meth.GetOutputType(), meth.GetInputType(), &bidiPairs)
			} else if meth.GetServerStreaming() && !meth.GetClientStreaming() {
				addPair(meth.GetOutputType(), meth.GetInputType(), &unidiPairs)
			}
		}
	}

	mixin := func(receiveType, class string) string {
		return fmt.Sprintf(
			`
	Q_OBJECT

	public: Q_SIGNAL void receivedMessage(%s msg);

	public: %s(const QString &origin = QString(), QWebSocketProtocol::Version version = QWebSocketProtocol::VersionLatest, QObject *parent = nullptr) : QWebSocket(origin, version, parent)
	{
		connect(this, &QWebSocket::binaryMessageReceived, [=](const QByteArray& msg) {
			%s incoming;

			if (!incoming.ParseFromArray(msg.constData(), msg.length())) {
				return;
			}

			Q_EMIT receivedMessage(incoming);
		});
	}
`,
			receiveType, class, receiveType,
		)
	}

	for _, pair := range bidiPairs {
		className := fmt.Sprintf(`Receive__%s__Send__%s__Stream`, sane(pair.fromServ), sane(pair.fromClie))

		add(fmt.Sprintf(`
class %s : public QWebSocket {
	%s

	public: [[nodiscard]] bool send(const %s& in) {
		std::string strData;
		if (!in.SerializeToString(&strData)) {
			return false;
		}
		QByteArray data = QByteArray::fromStdString(strData);

		auto count = sendBinaryMessage(data);
		return count == data.length();
	}
};`, className, mixin(typeToCxxNamespaces(pair.fromServ), className), typeToCxxNamespaces(pair.fromClie)))
	}
	for _, pair := range unidiPairs {
		className := fmt.Sprintf("Receive__%s__Stream", sane(pair.fromServ))
		add(fmt.Sprintf(`
class %s : public QWebSocket {
	%s
};`, className, mixin(typeToCxxNamespaces(pair.fromServ), className)))
	}

	return sb.String()
}

func getAllTypes(in []*descriptorpb.DescriptorProto) (out []*descriptorpb.DescriptorProto) {
	out = append(out, in...)

	for _, kind := range in {
		for _, item := range getAllTypes(kind.NestedType) {
			if strings.Contains(item.GetName(), "ExtensionEntry") {
				continue
			}
			*item.Name = *kind.Name + "." + *item.Name
			out = append(out, item)
		}
	}

	return
}

func generateClientTypesHeader(d *descriptorpb.FileDescriptorProto, mu []*descriptorpb.FileDescriptorProto) string {
	sb := strings.Builder{}
	sb.WriteString("#pragma once\n")
	sb.WriteString(getImports(d, mu, true))

	add := func(s string) { sb.WriteString(s + "\n") }

	for _, kind := range getAllTypes(d.MessageType) {
		add(fmt.Sprintf(`Q_DECLARE_METATYPE(%s)`, localTypeToCxxNamespaces(d.GetPackage())+"::"+localTypeToCxxNamespaces(kind.GetName())))
	}

	return sb.String()
}

func generateClientHeader(d *descriptorpb.FileDescriptorProto, mu []*descriptorpb.FileDescriptorProto) string {
	sb := strings.Builder{}
	sb.WriteString("#pragma once\n")
	sb.WriteString(getImports(d, mu, false))

	add := func(s string) { sb.WriteString(s + "\n") }

	add(generateClientSockets(d))

	for _, service := range d.Service {
		add(fmt.Sprintf("class %sServiceClient {", *service.Name))
		add(fmt.Sprintf(`	QString host;`))
		add(fmt.Sprintf(`	bool secure;`))
		add(fmt.Sprintf(`	QString httpProtocol() const { return secure ? QStringLiteral("https://") : QStringLiteral("http://"); }`))
		add(fmt.Sprintf(`	QString wsProtocol() const { return secure ? QStringLiteral("wss://") : QStringLiteral("ws://"); }`))
		add(fmt.Sprintf("\tpublic: explicit %sServiceClient(const QString& host, bool secure) : host(host), secure(secure) {}", *service.Name))
		add(`public:`)
		add(`	QMap<QByteArray,QString> universalHeaders;`)
		add(`	template<typename T> using Result = std::variant<T, QString>;`)
		{
			for _, meth := range service.Method {
				if meth.GetClientStreaming() && !meth.GetServerStreaming() {
					continue
				} else if meth.GetClientStreaming() && meth.GetServerStreaming() {
					sane := func(s string) string { return strings.ReplaceAll(typeToCxxNamespaces(s), "::", "_") }
					className := fmt.Sprintf(`Receive__%s__Send__%s__Stream`, sane(meth.GetOutputType()), sane(meth.GetInputType()))

					add(
						fmt.Sprintf(
							"\t[[ nodiscard ]] %s* %s(QMap<QByteArray,QString> headers = {});",

							className,
							meth.GetName(),
						),
					)
				} else if meth.GetServerStreaming() && !meth.GetClientStreaming() {
					sane := func(s string) string { return strings.ReplaceAll(typeToCxxNamespaces(s), "::", "_") }
					className := fmt.Sprintf(`Receive__%s__Stream`, sane(meth.GetOutputType()))

					add(
						fmt.Sprintf(
							"\t[[ nodiscard ]] %s* %s(const %s& in, QMap<QByteArray,QString> headers = {});",

							className,
							meth.GetName(),
							typeToCxxNamespaces(meth.GetInputType()),
						),
					)
				} else {
					// unary request

					add(
						fmt.Sprintf(
							"\t[[ nodiscard ]] Result<%s> %sSync(const %s& in, QMap<QByteArray,QString> headers = {});",

							typeToCxxNamespaces(meth.GetOutputType()),
							meth.GetName(),
							typeToCxxNamespaces(meth.GetInputType()),
						),
					)
					add(
						fmt.Sprintf(
							"\t[[ nodiscard ]] FutureResult<%s, QString> %s(const %s& in, QMap<QByteArray,QString> headers = {});",

							typeToCxxNamespaces(meth.GetOutputType()),
							meth.GetName(),
							typeToCxxNamespaces(meth.GetInputType()),
						),
					)
				}
			}
		}
		add("};")
	}

	return sb.String()
}

func generateClientImpl(d *descriptorpb.FileDescriptorProto) string {
	sb := strings.Builder{}

	add := func(s string) { sb.WriteString(s + "\n") }

	add(inc(convertCxxProto(*d.Name, "hrpc.client", "h")))
	add(inc("QThreadStorage"))

	add(`namespace {`)
	add(`QThreadStorage<QNetworkAccessManager*> globalNam;`)
	add(`void initialiseGlobalNam(bool secure, const QString& host) {
	if (globalNam.hasLocalData()) {
		return;
	}

	auto split = host.split(":");
	auto hname = split[0];
	auto port = split[1].toInt();
	
	globalNam.setLocalData(new QNetworkAccessManager);
	if (secure) {
		globalNam.localData()->connectToHostEncrypted(hname, port);
	} else {
		globalNam.localData()->connectToHost(hname, port);
	}
}`)
	add(`}`)

	for _, service := range d.Service {
		for _, meth := range service.Method {
			path := fmt.Sprintf("/%s.%s/%s", *d.Package, *service.Name, *meth.Name)

			if meth.GetClientStreaming() && !meth.GetServerStreaming() {
				continue
			} else if meth.GetClientStreaming() && meth.GetServerStreaming() {
				sane := func(s string) string { return strings.ReplaceAll(typeToCxxNamespaces(s), "::", "_") }
				className := fmt.Sprintf(`Receive__%s__Send__%s__Stream`, sane(meth.GetOutputType()), sane(meth.GetInputType()))

				add(
					fmt.Sprintf(
						"auto %sServiceClient::%s(QMap<QByteArray,QString> headers) -> %s*",

						service.GetName(),
						meth.GetName(),
						className,
					),
				)
				add(`{`)
				add(`auto url = QUrl(wsProtocol()+host); url.setPath(QStringLiteral("` + path + `")); auto req = QNetworkRequest(url);`)
				add(`
					for (const auto& item : universalHeaders.keys()) {
						req.setRawHeader(item, universalHeaders[item].toLocal8Bit());
					}
					for (const auto& item : headers.keys()) {
						req.setRawHeader(item, headers[item].toLocal8Bit());
					}
				`)
				add(fmt.Sprintf(`	auto sock = new %s();`, className))
				add(`	sock->open(req);`)
				add(`	return sock;`)
				add(`}`)
			} else if meth.GetServerStreaming() && !meth.GetClientStreaming() {
				sane := func(s string) string { return strings.ReplaceAll(typeToCxxNamespaces(s), "::", "_") }
				className := fmt.Sprintf(`Receive__%s__Stream`, sane(meth.GetOutputType()))

				add(
					fmt.Sprintf(
						"auto %sServiceClient::%s(const %s& in, QMap<QByteArray,QString> headers) -> %s*",

						service.GetName(),
						meth.GetName(),
						typeToCxxNamespaces(meth.GetInputType()),
						className,
					),
				)
				add(`{`)
				add(`auto url = QUrl(wsProtocol()+host); url.setPath(QStringLiteral("` + path + `")); auto req = QNetworkRequest(url);`)
				add(`
					for (const auto& item : universalHeaders.keys()) {
						req.setRawHeader(item, universalHeaders[item].toLocal8Bit());
					}
					for (const auto& item : headers.keys()) {
						req.setRawHeader(item, headers[item].toLocal8Bit());
					}
				`)
				add(fmt.Sprintf(`	auto sock = new %s();`, className))
				add(`	std::string strData;`)
				add(`	if (!in.SerializeToString(&strData)) { return nullptr; }`)
				add(`	QByteArray data = QByteArray::fromStdString(strData);`)
				add(`	sock->open(req);`)
				add(`	QObject::connect(sock, &QWebSocket::connected, [=]() { sock->sendBinaryMessage(data); });`)
				add(`	return sock;`)
				add(`}`)
			} else {
				// unary request

				addBody := func() {
					add(fmt.Sprintf(
						`

	initialiseGlobalNam(secure, host);

	QUrl serviceURL = QUrl(httpProtocol()+host);
	serviceURL.setPath(QStringLiteral("` + path + `"));

	QNetworkRequest req(serviceURL);
	for (const auto& item : universalHeaders.keys()) {
		req.setRawHeader(item, universalHeaders[item].toLocal8Bit());
	}
	for (const auto& item : headers.keys()) {
		req.setRawHeader(item, headers[item].toLocal8Bit());
	}
	req.setRawHeader("content-type", "application/hrpc");
	req.setAttribute(QNetworkRequest::Http2AllowedAttribute, true);

	auto nam = globalNam.localData();
	auto val = nam->post(req, data);
`,
					))
				}

				add(
					fmt.Sprintf(
						"auto %sServiceClient::%sSync(const %s& in, QMap<QByteArray,QString> headers) -> %sServiceClient::Result<%s>",

						service.GetName(),
						meth.GetName(),
						typeToCxxNamespaces(meth.GetInputType()),
						service.GetName(),
						typeToCxxNamespaces(meth.GetOutputType()),
					),
				)

				add(`
{
	std::string strData;
	if (!in.SerializeToString(&strData)) { return {QStringLiteral("failed to serialize protobuf")}; }
	QByteArray data = QByteArray::fromStdString(strData);
`)

				addBody()
				add(
					fmt.Sprintf(`

	while (!val->isFinished()) {
		QCoreApplication::processEvents();
	}

	auto response = val->readAll();

	if (val->error() != QNetworkReply::NoError) {
		return {QStringLiteral("network failure(%%1): %%2\n%%3").arg(val->error()).arg(val->errorString()).arg(QString::fromLocal8Bit(response))};
	}

	%s ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};
`, typeToCxxNamespaces(meth.GetOutputType())),
				)
				add(`}`)

				add(
					fmt.Sprintf(
						"FutureResult<%s, QString> %sServiceClient::%s(const %s& in, QMap<QByteArray,QString> headers)",

						typeToCxxNamespaces(meth.GetOutputType()),
						service.GetName(),
						meth.GetName(),
						typeToCxxNamespaces(meth.GetInputType()),
					),
				)

				add(`
{
	FutureResult<` + typeToCxxNamespaces(meth.GetOutputType()) + `, QString> res;

	std::string strData;
	if (!in.SerializeToString(&strData)) { res.fail({QStringLiteral("failed to serialize protobuf")}); return res; }
	QByteArray data = QByteArray::fromStdString(strData);
`)
				addBody()
				add(
					fmt.Sprintf(`

	QObject::connect(val, &QNetworkReply::finished, [val, res]() mutable {
		auto response = val->readAll();

		if (val->error() != QNetworkReply::NoError) {
			val->deleteLater();
			res.fail({QStringLiteral("request failure(%%1): %%2\n%%3").arg(val->error()).arg(val->errorString()).arg(QString::fromLocal8Bit(response))});
			return;
		}

		%s ret;
		if (!ret.ParseFromArray(response.constData(), response.length())) {
			val->deleteLater();
			res.fail({QStringLiteral("error parsing response into protobuf")});
			return;
		}
		
		val->deleteLater();
		res.succeed({ret});
		return;
	});

	return res;
`, typeToCxxNamespaces(meth.GetOutputType())),
				)

				add(`}`)
			}
		}
	}

	return sb.String()
}

// GenerateQtCxxClient generates a C++ Qt client
func GenerateQtCxxClient(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)
	for _, file := range d.ProtoFile {
		{
			metatypesFile := new(pluginpb.CodeGeneratorResponse_File)
			metatypesFile.Name = new(string)
			*metatypesFile.Name = convertCxxProto(*file.Name, "hrpc.types", "h")

			metatypesFile.Content = new(string)
			*metatypesFile.Content = generateClientTypesHeader(file, d.ProtoFile)

			r.File = append(r.File, metatypesFile)
		}
		if len(file.Service) == 0 {
			continue
		}
		{
			headerFile := new(pluginpb.CodeGeneratorResponse_File)
			headerFile.Name = new(string)
			*headerFile.Name = convertCxxProto(*file.Name, "hrpc.client", "h")

			headerFile.Content = new(string)
			*headerFile.Content = generateClientHeader(file, d.ProtoFile)

			r.File = append(r.File, headerFile)
		}
		{
			implFile := new(pluginpb.CodeGeneratorResponse_File)
			implFile.Name = new(string)
			*implFile.Name = convertCxxProto(*file.Name, "hrpc.client", "cpp")

			implFile.Content = new(string)
			*implFile.Content = generateClientImpl(file)

			r.File = append(r.File, implFile)
		}
	}
	return
}
