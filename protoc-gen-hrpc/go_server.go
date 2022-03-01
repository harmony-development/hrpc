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
			add("\"google.golang.org/protobuf/proto\"")
			indent--
			add(")")

			add(`
			type VTProtoMessage interface {
				proto.Message
				MarshalVT() ([]byte, error)
				UnmarshalVT([]byte) error
			}

			func FromProtoChannel[T VTProtoMessage](in chan VTProtoMessage) chan T {
				res := make(chan T)
				go func() {
					for {
						v := <-in
						res <- v.(T)
					}
				}()
			
				return res
			}
			
			func ToProtoChannel[T VTProtoMessage](in chan T) chan VTProtoMessage {
				res := make(chan VTProtoMessage)
				go func() {
					for {
						v := <-in
						res <- v
					}
				}()
			
				return res
			}			
			`)

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

			add("func (h *%sHandler[T]) Routes() map[string]func(T, VTProtoMessage) (VTProtoMessage, error) {", *service.Name)
			indent++
			add("return map[string]func(T, VTProtoMessage) (VTProtoMessage, error){")
			indent++
			for _, meth := range requests {
				add("\"%s.%s/%s\": func(c T, msg VTProtoMessage) (VTProtoMessage, error) {", *f.Package, *service.Name, *meth.Name)
				indent++
				add("return h.impl.%s(c, msg.(*%s))", *meth.Name, rawType(*meth.InputType))
				indent--
				add("},")
			}
			indent--
			add("}")
			indent--
			add("}")

			add("func (h *%sHandler[T]) StreamRoutes() map[string]func(T, chan VTProtoMessage) (chan VTProtoMessage, error) {", *service.Name)
			indent++
			add("return map[string]func(T, chan VTProtoMessage) (chan VTProtoMessage, error){")
			indent++
			for _, meth := range streams {
				add("\"%s.%s/%s\": func(c T, msg chan VTProtoMessage) (chan VTProtoMessage, error) {", *f.Package, *service.Name, *meth.Name)
				indent++
				add("res, err := h.impl.%s(c, FromProtoChannel[*%s](msg))", *meth.Name, rawType(*meth.InputType))
				add("return ToProtoChannel(res), err")
				indent--
				add("},")
			}
			indent--
			add("}")
			indent--
			add("}")

			r.File = append(r.File, file)
		}
	}
	return
}
