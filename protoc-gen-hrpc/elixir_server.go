// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"path"
	"strings"

	"google.golang.org/protobuf/types/pluginpb"
)

func GenerateElixirServer(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)

	for _, f := range d.ProtoFile {
		if len(f.Service) == 0 {
			continue
		}

		file := new(pluginpb.CodeGeneratorResponse_File)
		file.Name = new(string)
		*file.Name = path.Join(strings.TrimSuffix(*f.Name, ".proto")) + ".hrpc.ex"

		dat := strings.Builder{}
		indent := 0
		add := func(format string, a ...interface{}) {
			dat.WriteString(strings.Repeat("\t", indent))
			dat.WriteString(fmt.Sprintf(format, a...))
			dat.WriteRune('\n')
		}
		addI := func(format string, a ...interface{}) {
			add(format, a...)
			indent++
		}
		addD := func(format string, a ...interface{}) {
			indent--
			add(format, a...)
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

		elixirName := func(kind string) string {
			split := strings.Split(kind[1:], ".")
			for idx := range split {
				split[idx] = strings.Title(split[idx])
			}
			return strings.Join(split, ".")
		}
		_ = elixirName

		for _, service := range f.Service {
			addI(`defmodule %s.%sService do`, elixirName(" "+f.GetPackage()), service.GetName())

			addI(`def endpoints do`)

			addI(`[`)

			for _, rpc := range service.GetMethod() {
				route := fmt.Sprintf(`/%s.%s/%s`, f.GetPackage(), service.GetName(), rpc.GetName())

				add(`{"%s", "%s", %t, %t, %s, %s},`, rpc.GetName(), route, rpc.GetClientStreaming(), rpc.GetServerStreaming(), elixirName(rpc.GetInputType()), elixirName(rpc.GetOutputType()))
			}

			addD(`]`)

			addD(`end`)

			addD(`end`)
		}

		r.File = append(r.File, file)
	}
	return
}
