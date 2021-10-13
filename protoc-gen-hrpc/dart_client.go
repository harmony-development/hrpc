package main

import (
	"fmt"
	"path"
	"strings"

	"google.golang.org/protobuf/types/pluginpb"
)

func GenerateDartClient(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)
	for _, f := range d.ProtoFile {
		if len(f.Service) == 0 {
			continue
		}

		file := new(pluginpb.CodeGeneratorResponse_File)
		file.Name = new(string)
		*file.Name = path.Join(strings.TrimSuffix(*f.Name, ".proto")) + ".client.hrpc.dart"

		dat := strings.Builder{}
		indent := 0
		add := func(format string, v ...interface{}) {
			dat.WriteString(strings.Repeat("\t", indent))
			dat.WriteString(fmt.Sprintf(format, v...))
			dat.WriteRune('\n')
		}

		defer func() {
			file.Content = new(string)
			*file.Content = dat.String()
		}()

		nestLevels := strings.Count(*f.Name, "/")
		arr := make([]string, nestLevels)
		for idx := range arr {
			arr[idx] = ".."
		}

		for _, dep := range f.Dependency {
			add(`export '%s/%s';`, path.Join(arr...), strings.TrimSuffix(dep, ".proto")+".pb.dart")
			add(`import '%s/%s';`, path.Join(arr...), strings.TrimSuffix(dep, ".proto")+".pb.dart")
		}
		add(`export '%s';`, strings.TrimSuffix(path.Base(*f.Name), ".proto")+".pb.dart")
		add(`import '%s';`, strings.TrimSuffix(path.Base(*f.Name), ".proto")+".pb.dart")
		add(`import 'package:http/http.dart' as $http;`)
		add(`import 'dart:io' as $io;`)
		add(`import 'package:async/async.dart' as $async;`)

		kind := func(in string) string {
			split := strings.Split(in, ".")
			return split[len(split)-1]
		}

		for _, service := range f.Service {
			add(`class %sClient {`, *service.Name)
			indent++
			add(`late Uri server;`)
			add(`late Map<String,String> commonHeaders;`)

			add(`Uri get wsServer => server.hasScheme ? server.replace(scheme: server.scheme == "https" ? "wss" : "ws") : server.replace(scheme: "wss");`)
			for _, meth := range service.Method {
				if meth.GetClientStreaming() && !meth.GetServerStreaming() {
					continue
				} else if meth.GetClientStreaming() && meth.GetServerStreaming() {
					add(`Stream<%s> %s(Stream<%s> input, {Map<String,dynamic> headers = const {}}) async* {`, kind(*meth.OutputType), *meth.Name, kind(*meth.InputType))
					indent++
					{
						add(`var socket = await $io.WebSocket.connect(this.wsServer.replace(path: "/%s.%s/%s").toString(), headers: headers..addAll(this.commonHeaders));`, *f.Package, *service.Name, *meth.Name)
						add(`var combined = $async.StreamGroup.merge<dynamic>([socket, input]);`)
						add(`await for (var value in combined) {`)
						indent++
						{
							add(`if (value is List<int>) {`)
							indent++
							{
								add(`yield %s.fromBuffer(value);`, kind(*meth.OutputType))
							}
							indent--
							add(`} else if (value is %s) {`, kind(*meth.InputType))
							indent++
							{
								add(`socket.add(value.writeToBuffer());`)
							}
							indent--
							add(`}`)
						}
						indent--
						add(`}`)
					}
					indent--
					add(`}`)
				} else if !meth.GetServerStreaming() && !meth.GetClientStreaming() {
					add(`Future<%s> %s(%s input, {Map<String,String> headers = const {}}) async {`, kind(*meth.OutputType), *meth.Name, kind(*meth.InputType))
					indent++
					{
						add(`var response = await $http.post(this.server.replace(path: "/%s.%s/%s"), body: input.writeToBuffer(), headers: {"content-type": "application/hrpc"}..addAll(headers)..addAll(this.commonHeaders));`, *f.Package, *service.Name, *meth.Name)
						add(`if (response.statusCode != 200) { throw response; }`)
						add(`return %s.fromBuffer(response.bodyBytes);`, kind(*meth.OutputType))
					}
					indent--
					add(`}`)
				} else if meth.GetServerStreaming() && !meth.GetClientStreaming() {
					add(`Stream<%s> %s(%s input, {Map<String,dynamic> headers = const {}}) async* {`, kind(*meth.OutputType), *meth.Name, kind(*meth.InputType))
					indent++
					{
						add(`var socket = await $io.WebSocket.connect(this.server.replace(path: "/%s.%s/%s").toString(), headers: headers..addAll(this.commonHeaders));`, *f.Package, *service.Name, *meth.Name)
						add(`socket.add(input.writeToBuffer());`)
						add(`await for (var value in socket) {`)
						indent++
						{
							add(`if (value is List<int>) {`)
							indent++
							{
								add(`yield %s.fromBuffer(value);`, kind(*meth.OutputType))
							}
							indent--
							add(`}`)
						}
						indent--
						add(`}`)
					}
					indent--
					add(`}`)
				}
			}
			indent--
			add(`}`)
		}

		r.File = append(r.File, file)
	}
	return
}
