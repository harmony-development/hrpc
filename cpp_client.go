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

func getImports(d *descriptorpb.FileDescriptorProto) string {
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
	add(`#include <variant>`)

	for _, kind := range d.Service {
		for _, meth := range kind.Method {
			if meth.GetClientStreaming() || meth.GetServerStreaming() {
				add(`// #include <QtWebSockets>`)
			}
		}
	}

	for _, dep := range d.Dependency {
		add(inc(convertCxxProto(dep, "pb", "h")))
	}

	return strings.Join(set, "\n") + "\n"
}

func typeToCxxNamespaces(s string) string {
	return strings.ReplaceAll(s[1:], ".", "::")
}

func generateClientHeader(d *descriptorpb.FileDescriptorProto) string {
	sb := strings.Builder{}
	sb.WriteString(getImports(d))

	add := func(s string) { sb.WriteString(s + "\n") }

	for _, service := range d.Service {
		add(fmt.Sprintf("class %sServiceClient {", *service.Name))
		add(fmt.Sprintf(`	QString host;`))
		add(fmt.Sprintf(`	bool secure;`))
		add(fmt.Sprintf(`	QSharedPointer<QNetworkAccessManager> nam;`))
		add(fmt.Sprintf(`	QString httpProtocol() const { return secure ? QStringLiteral("https://") : QStringLiteral("http://"); }`))
		add(fmt.Sprintf("\texplicit %sServiceClient(const QString& host, bool secure) : host(host), secure(secure), nam(new QNetworkAccessManager) {}", *service.Name))
		add(`public:`)
		add(`	template<typename T> using Result = std::variant<T, QString>;`)
		{
			for _, meth := range service.Method {
				if meth.GetClientStreaming() && !meth.GetServerStreaming() {
					continue
				} else if meth.GetClientStreaming() && meth.GetServerStreaming() {
					add(`// todo client <-> server stream`)
				} else if meth.GetServerStreaming() && !meth.GetClientStreaming() {
					add(`// todo client <- server stream`)
				} else {
					// unary request

					add(
						fmt.Sprintf(
							"\tResult<%s> %s(const %s& in, QMap<QByteArray,QString> headers = {});",

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

	for _, service := range d.Service {
		for _, meth := range service.Method {
			if meth.GetClientStreaming() && !meth.GetServerStreaming() {
				continue
			} else if meth.GetClientStreaming() && meth.GetServerStreaming() {
				add(`// todo client <-> server stream`)
			} else if meth.GetServerStreaming() && !meth.GetClientStreaming() {
				add(`// todo client <- server stream`)
			} else {
				// unary request

				add(
					fmt.Sprintf(
						"auto %sServiceClient::%s(const %s& in, QMap<QByteArray,QString> headers) -> %sServiceClient::Result<%s>",

						service.GetName(),
						meth.GetName(),
						typeToCxxNamespaces(meth.GetInputType()),
						service.GetName(),
						typeToCxxNamespaces(meth.GetOutputType()),
					),
				)
				add(`{`)

				add(`	QByteArray data = QByteArray::fromStdString(in.SerializeAsString());`)
				add(
					fmt.Sprintf(`
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
		return {QStringLiteral("network failure(%%1): %%2").arg(val->error()).arg(val->errorString())};
	}

	auto response = val->readAll();

	%s ret;
	if (!ret.ParseFromArray(response.constData(), response.length())) {
		return {QStringLiteral("error parsing response into protobuf")};
	}

	return {ret};
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
		if len(file.Service) == 0 {
			continue
		}
		{
			headerFile := new(pluginpb.CodeGeneratorResponse_File)
			headerFile.Name = new(string)
			*headerFile.Name = convertCxxProto(*file.Name, "hrpc.client", "h")

			headerFile.Content = new(string)
			*headerFile.Content = generateClientHeader(file)

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
