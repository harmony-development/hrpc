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

func GenerateGoServer(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)

	for _, f := range d.ProtoFile {
		for _, service := range f.Service {
			file := new(pluginpb.CodeGeneratorResponse_File)
			file_path := strings.TrimSuffix(*f.Name, ".proto")
			name := file_path + ".hrpc.go"
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

			add("package %s", strings.ReplaceAll(file_path[:strings.LastIndex(file_path, "/")], "/", ""))

			add("import (")
			indent++
			add("\"context\"")
			add("goserver \"github.com/harmony-development/hrpc/pkg/go-server\"")
			indent--
			add(")")

			var requests []*descriptorpb.MethodDescriptorProto
			var streams []*descriptorpb.MethodDescriptorProto

			for _, meth := range service.Method {
				if meth.GetClientStreaming() || meth.GetServerStreaming() {
					streams = append(streams, meth)
				} else {
					requests = append(requests, meth)
				}
			}

			add("type %s[T context.Context] interface {", *service.Name)
			indent++

			for _, meth := range requests {
				add("%s(T, *%s) (*%s, error)", *meth.Name, rawType(*meth.InputType), rawType(*meth.OutputType))
			}
			for _, meth := range streams {
				add("%s(T, chan *%s) (chan *%s, error)", *meth.Name, rawType(*meth.InputType), rawType(*meth.OutputType))
			}

			indent--
			add("}")

			add("type %sHandler[T context.Context] struct {", *service.Name)
			indent++
			add("impl %s[T]", *service.Name)
			indent--
			add("}")

			methodData := func(i string, o string) string {
				return fmt.Sprintf("goserver.MethodData{Input: &%s{}, Output: &%s{}}", i, o)
			}

			add("func (h *%sHandler[T]) Routes() map[string]goserver.UnaryMethodData[T] {", *service.Name)
			indent++
			add("return map[string]goserver.UnaryMethodData[T]{")
			indent++
			for _, meth := range requests {
				reqPath := fmt.Sprintf("%s.%s/%s", *f.Package, *service.Name, *meth.Name)
				input := rawType(*meth.InputType)
				output := rawType(*meth.OutputType)
				add(`"%s": {`, reqPath)
				add("MethodData: %s,", methodData(input, output))
				add(`Handler: func(c T, msg goserver.VTProtoMessage) (goserver.VTProtoMessage, error) {
						return h.impl.%s(c, msg.(*%s))
					},
				},
				`, *meth.Name, input)
			}
			indent--
			add("}")
			indent--
			add("}")

			add("func (h *%sHandler[T]) StreamRoutes() map[string]goserver.StreamMethodData[T] {", *service.Name)
			indent++
			add("return map[string]goserver.StreamMethodData[T]{")
			indent++
			for _, meth := range streams {
				reqPath := fmt.Sprintf("%s.%s/%s", *f.Package, *service.Name, *meth.Name)
				input := rawType(*meth.InputType)
				output := rawType(*meth.OutputType)

				add(`"%s": {`, reqPath)
				indent++
				add("MethodData: %s,", methodData(input, output))
				add(`Handler: func(c T, msg chan goserver.VTProtoMessage) (chan goserver.VTProtoMessage, error) {
					res, err := h.impl.%s(c, goserver.FromProtoChannel[*%s](msg))
					return goserver.ToProtoChannel(res), err
				},`, *meth.Name, input)
				indent--
				add("},")
			}
			indent--
			add("}")
			add("}")

			r.File = append(r.File, file)
		}
	}
	return
}
