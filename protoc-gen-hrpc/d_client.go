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

func GenerateDClient(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)
	for _, f := range d.ProtoFile {
		if len(f.Service) == 0 {
			continue
		}

		file := new(pluginpb.CodeGeneratorResponse_File)
		file.Name = new(string)
		*file.Name = path.Join(strings.Split(*f.Package, ".")...) + "/" + path.Base(strings.TrimSuffix(*f.Name, ".proto")) + "Hrpc.d"

		dat := strings.Builder{}
		add := func(format string, v ...interface{}) {
			dat.WriteString(fmt.Sprintf(format, v...))
			dat.WriteRune('\n')
		}
		final := func(in []string) string {
			return in[len(in)-1]
		}
		splitter := func(s string) []string {
			return strings.Split(s, ".")
		}

		defer func() {
			file.Content = new(string)
			*file.Content = dat.String()
		}()

		module := *f.Package + "." + path.Base(strings.TrimSuffix(*f.Name, ".proto"))
		add("module %sHrpc;", module)
		add("import %s;", module)

		for _, dep := range f.Dependency {
			add(`import %s;`, strings.ReplaceAll(strings.TrimSuffix(dep, ".proto"), "/", "."))
		}

		for _, service := range f.Service {
			add(`mixin Service!(`)
			add(`	"%s", "%s",`+"\n", *service.Name, *f.Package)
			for _, meth := range service.Method {
				if meth.GetClientStreaming() && !meth.GetServerStreaming() {
					continue
				} else if meth.GetClientStreaming() && meth.GetServerStreaming() {

				} else if !meth.GetServerStreaming() && !meth.GetClientStreaming() {
					add(`	"Unary", %s, %s, "%s",`, final(splitter(*meth.InputType)), final(splitter(*meth.OutputType)), *meth.Name)
				}
			}
			add(`);`)
		}

		r.File = append(r.File, file)
	}
	return
}
