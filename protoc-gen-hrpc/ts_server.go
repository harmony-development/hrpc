// SPDX-FileCopyrightText: 2021 Harmony Contributors
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func rawType(s string) string {
	return s[strings.LastIndex(s, ".")+1:]
}

func GenerateTSServer(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)

	for _, f := range d.ProtoFile {
		for _, service := range f.Service {
			file := new(pluginpb.CodeGeneratorResponse_File)
			file_path := strings.TrimSuffix(*f.Name, ".proto")
			name := file_path + ".iface.ts"
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

			importPath := "./" + file_path[strings.LastIndex(file_path, "/")+1:]
			importTypes := []string{}
			for _, method := range service.Method {
				importTypes = append(importTypes, rawType(*method.InputType))
				importTypes = append(importTypes, rawType(*method.OutputType))
			}
			add("import {")
			indent++
			for _, importType := range importTypes {
				add("%s,", importType)
			}
			indent--
			add("} from '%s';", importPath)

			for _, meth := range service.Method {
				if meth.GetClientStreaming() || meth.GetServerStreaming() {
					streams = append(streams, meth)
				} else {
					requests = append(requests, meth)
				}
			}

			add("export interface %s<C> {", *service.Name)
			indent++

			for _, meth := range requests {
				add("")
				add("%s(ctx: C, request: %s): Promise<%s>", strcase.ToLowerCamel(*meth.Name), rawType(*meth.InputType), rawType(*meth.OutputType))
			}

			for _, meth := range streams {
				add("")
				add("%s(ctx: C, request: AsyncIterable<%s>): AsyncIterable<%s>", strcase.ToLowerCamel(*meth.Name), rawType(*meth.InputType), rawType(*meth.OutputType))
			}
			indent--
			add("}")

			r.File = append(r.File, file)
		}
	}
	return
}
