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

		kind := func(in string) string {
			split := strings.Split(in, ".")
			return split[len(split)-1]
		}

		for _, service := range f.Service {
			add(`class %sClient {`, *service.Name)
			indent++
			for _, meth := range service.Method {
				if meth.GetClientStreaming() && !meth.GetServerStreaming() {
					continue
				} else if meth.GetClientStreaming() && meth.GetServerStreaming() {

				} else if !meth.GetServerStreaming() && !meth.GetClientStreaming() {
					add(`Future<%s> %s(%s input) async {`, kind(*meth.OutputType), *meth.Name, kind(*meth.InputType))
					indent++
					{
						add(`var response = await $http.post();`)
					}
					indent--
					add(`}`)
				}
			}
			add(`}`)
		}

		r.File = append(r.File, file)
	}
	return
}
