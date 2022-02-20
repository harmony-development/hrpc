// SPDX-FileCopyrightText: 2021 Harmony Contributors
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func GenerateCsharpServer(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)

	for _, f := range d.ProtoFile {
		for _, service := range f.Service {
			file := new(pluginpb.CodeGeneratorResponse_File)
			name := *service.Name + ".server.hrpc.cs"
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

			var requests []*descriptorpb.MethodDescriptorProto
			var streams []*descriptorpb.MethodDescriptorProto

			for _, meth := range service.Method {
				if meth.GetClientStreaming() || meth.GetServerStreaming() {
					streams = append(streams, meth)
				} else {
					requests = append(requests, meth)
				}
			}

			add("")

			add("namespace Harmony.Server;\n")

			add("public partial interface IRequestHost\n{")
			indent++

			for _, meth := range requests {
				add("")
				add("public Task<%s> %s(%s pm);", fixMethodString(*meth.OutputType), *meth.Name, fixMethodString(*meth.InputType))
			}

			indent--
			add("}")
			add("")

			add("public partial interface IStreamHost\n{")
			indent++

			for _, meth := range streams {
				add("")
				add("public void %s(%s pm); // stream %s", *meth.Name, fixMethodString(*meth.InputType), fixMethodString(*meth.OutputType))
			}

			indent--
			add("}")

			r.File = append(r.File, file)
		}
	}
	return
}
