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
			add(`import '%s/%s';`, path.Join(arr...), strings.TrimSuffix(dep, ".proto")+".pb.dart")
		}
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
			add(`bool secure = true;`)
			add(`String host;`)
			add(`%sClient({this.secure,this.host});`, *service.Name)
			add(`String get unaryPrefix => secure ? "https" : "http";`)
			add(`String get wsPrefix => secure ? "wss" : "ws";`)
			for _, meth := range service.Method {
				if meth.GetClientStreaming() && !meth.GetServerStreaming() {
					continue
				} else if meth.GetClientStreaming() && meth.GetServerStreaming() {
					add(`Stream<%s> %s(Stream<%s> input, {Map<String,dynamic> headers}) async* {`, kind(*meth.OutputType), *meth.Name, kind(*meth.InputType))
					indent++
					{
						add(`var socket = await $io.WebSocket.connect("${this.unaryPrefix}://%s/.%s/%s", headers: headers);`, *f.Package, *service.Name, *meth.Name)
						add(`var combined = $async.StreamGroup.merge([socket, input]);`)
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
								add(`await socket.add(value.writeToBuffer());`)
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
					add(`Future<%s> %s(%s input, {Map<String,String> headers}) async {`, kind(*meth.OutputType), *meth.Name, kind(*meth.InputType))
					indent++
					{
						add(`var response = await $http.post("${this.unaryPrefix}://%s.%s/%s", body: input.writeToBuffer(), headers: {"content-type": "application/octet-stream"}..addAll(headers));`, *f.Package, *service.Name, *meth.Name)
						add(`if (response.statusCode != 200) { throw response; }`)
						add(`return %s.fromBuffer(response.bodyBytes);`, kind(*meth.OutputType))
					}
					indent--
					add(`}`)
				} else if meth.GetServerStreaming() && !meth.GetClientStreaming() {
					add(`Stream<%s> %s(%s input, {Map<String,dynamic> headers}) async* {`, kind(*meth.OutputType), *meth.Name, kind(*meth.InputType))
					indent++
					{
						add(`var socket = await $io.WebSocket.connect("${this.unaryPrefix}://%s/.%s/%s", headers: headers);`, *f.Package, *service.Name, *meth.Name)
						add(`await socket.add(input.writeToBuffer());`)
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
