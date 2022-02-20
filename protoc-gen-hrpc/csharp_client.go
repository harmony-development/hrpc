// SPDX-FileCopyrightText: 2021 Harmony Contributors
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/pluginpb"
)

func GenerateCsharpClient(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)

	for _, f := range d.ProtoFile {
		for _, service := range f.Service {
			file := new(pluginpb.CodeGeneratorResponse_File)
			name := *service.Name + ".client.hrpc.cs"
			file.Name = &name

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

			add("using System.Text.RegularExpressions;")
			add("using Hrpc;")

			add("")

			add("namespace Harmony.Client;")
			add("")

			add("public class %sClient\n{", *service.Name)
			indent++

			add("public readonly Dictionary<string, string> DefaultRequestHeaders = new();")
			add("private string _host { get; init; }")
			add("private string _wsHost => new Regex(\"http\").Replace(_host, \"ws\", 1);")
			add("public readonly HttpClient _client = new();")
			add("")

			add("public %sClient(string host)", *service.Name)
			add("{")
			indent++
			add("_host = host;")
			indent--
			add("}")

			for _, meth := range service.Method {
				path := fmt.Sprintf("/%s.%s/%s", *f.Package, *service.Name, *meth.Name)

				if meth.GetClientStreaming() || meth.GetServerStreaming() {
					add("// stream %s", *meth.Name)
					add("public async Task<StreamClient<%s>> %s(%s pm)", fixMethodString(*meth.OutputType), *meth.Name, fixMethodString(*meth.InputType))
					add("{")
					indent++

					add("var stream = new StreamClient<%s>(DefaultRequestHeaders);", fixMethodString(*meth.OutputType))
					add("await stream.Connect(_wsHost + \"%s\", pm);", path)
					add("return stream;")

					indent--
					add("}")
					add("")
				} else {
					add("// unary %s", *meth.Name)
					add("public async Task<%s> %s(%s pm)", fixMethodString(*meth.OutputType), *meth.Name, fixMethodString(*meth.InputType))
					indent++

					add("=> await _client.HrpcUnaryAsync<%s, %s>(_host + \"%s\", pm);", fixMethodString(*meth.InputType), fixMethodString(*meth.OutputType), path)

					indent--
					add("")
				}
			}

			indent--
			add("}")
			r.File = append(r.File, file)

		}
	}
	return
}

func fixMethodString(inp string) string {
	// why
	inp = strings.TrimPrefix(inp, ".")

	tmp := strings.Split(inp, ".")
	for i := range tmp {
		tmp[i] = strings.ToUpper(string(tmp[i][0])) + tmp[i][1:]
	}
	return strings.Join(tmp, ".")
}
