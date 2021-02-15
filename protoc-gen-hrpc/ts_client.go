package main

import (
	"fmt"
	"path"
	"strings"

	"google.golang.org/protobuf/types/pluginpb"
)

func GenerateTSClient(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)
	for _, f := range d.ProtoFile {
		if len(f.Service) == 0 {
			continue
		}

		file := new(pluginpb.CodeGeneratorResponse_File)
		file.Name = new(string)
		*file.Name = path.Join(strings.Split(*f.Package, ".")...) + "/" + path.Base(strings.TrimSuffix(*f.Name, ".proto")) + ".ts"

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
		finalRemoved := func(in []string) string {
			return strings.Join(in[:len(in)-1], ".")
		}

		defer func() {
			file.Content = new(string)
			*file.Content = dat.String()
		}()

		// module := *f.Package + "." + path.Base(strings.TrimSuffix(*f.Name, ".proto"))
		namespace := strings.TrimSuffix(*f.Package, ".proto")
		add("import gen from '%s';", "./output")

		add("import svc = gen.%s;", namespace)

		for _, service := range f.Service {
			add(`class %s {`, *service.Name)
			add(`
			host: string;
			session?: string;

			constructor(host: string) {
				this.host = host;
			}
			`)
			add(`
			unary(endpoint: string, body: Uint8Array) {
				return fetch(
					` + "`" + `${this.host}${endpoint}` + "`" + `,
					{
						method: "POST",
						body,
						headers: {
							"Content-Type": "application/octet-stream",
							Authorization: this.session || "",
						}
					}
				)
			}`)
			add(`
			stream(endpoint: string) {
				return new WebSocket(` + "`" + `${this.host}${endpoint}` + "`" + `, ["access_token", this.session || ""])
			}`)
			for _, meth := range service.Method {
				path := fmt.Sprintf("/%s.%s/%s", *f.Package, *service.Name, *meth.Name)
				if meth.GetClientStreaming() || meth.GetServerStreaming() {
					add(`%s() {`, *meth.Name)
					add(`return this.stream("%s")`, path)
					add("}")
				} else {
					inputPackage := finalRemoved(splitter((*meth.InputType)[1:]))
					inputType := final(splitter((*meth.InputType)[1:]))
					outputPackage := finalRemoved(splitter((*meth.OutputType)[1:]))
					outputType := final(splitter((*meth.OutputType)[1:]))

					inputJSType := ""
					IinputJSType := ""
					if inputPackage == *f.Package {
						IinputJSType = fmt.Sprintf("svc.I%s", inputType)
						inputJSType = fmt.Sprintf("svc.%s", inputType)
					} else {
						IinputJSType = fmt.Sprintf("gen.%s.I%s", inputPackage, inputType)
						inputJSType = fmt.Sprintf("gen.%s.%s", inputPackage, inputType)
					}

					outputJSType := ""
					if inputPackage == *f.Package {
						outputJSType = fmt.Sprintf("svc.%s", outputType)
					} else {
						outputJSType = fmt.Sprintf("gen.%s.%s", outputPackage, outputType)
					}

					add(`async %s (req: %s) {`, *meth.Name, IinputJSType)
					add(`const buffer = %s.encode(
						%s.create(req)
					).finish();`, inputJSType, inputJSType)
					add(`const resp = await this.unary('%s', buffer)`, path)
					add(`return %s.decode(new Uint8Array(await resp.arrayBuffer()));`, outputJSType)
					add("}")
				}
			}
			add(`}`)
		}

		r.File = append(r.File, file)
	}
	return
}
