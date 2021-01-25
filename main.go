package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"

	. "github.com/dave/jennifer/jen"
)

func main() {
	gen := pluginpb.CodeGeneratorRequest{}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = proto.Unmarshal(data, &gen)
	if err != nil {
		panic(err)
	}

	out := pluginpb.CodeGeneratorResponse{}

	for _, item := range gen.ProtoFile {
		outfile := new(pluginpb.CodeGeneratorResponse_File)
		outfile.Name = new(string)
		base := strings.TrimSuffix(filepath.Base(*item.Name), ".proto")
		*outfile.Name = fmt.Sprintf("%s/%s.hrpc.go", *item.Options.GoPackage, base)

		if len(item.Service) == 0 {
			continue
		}

		f := NewFile(filepath.Base(*item.Options.GoPackage))

		for _, serv := range item.Service {

			sname := fmt.Sprintf("%sClient", *serv.Name)

			f.Type().Id(sname).Struct(
				Id("client *").Qual("net/http", "Client"),
				Id("serverURL").Id("string"),
			)

			f.Func().Id(fmt.Sprintf("New%s", sname)).Params(Id("url").String()).Id(fmt.Sprintf("*%s", sname)).Block(
				Return(Op("&").Id(sname).Values(Dict{
					Id("serverURL"): Id("url"),
					Id("client"):    Op("&").Qual("net/http", "Client").Block(),
				})),
			)

			for _, method := range serv.Method {
				if (method.ClientStreaming != nil && *method.ClientStreaming) || (method.ServerStreaming != nil && *method.ServerStreaming) {
					continue
				}

				name := func(in string) Code {
					wellKnown := map[string]struct {
						pkg string
						id  string
					}{
						".google.protobuf.Empty": {
							pkg: "github.com/golang/protobuf/ptypes/empty",
							id:  "Empty",
						},
					}

					if val, ok := wellKnown[in]; ok {
						return Qual(val.pkg, val.id)
					}

					split := strings.Split(*method.InputType, ".")
					final := split[len(split)-1]

					return Qual(*item.Options.GoPackage, final)
				}

				errHandler := func(msg string) Code {
					return If(Id("err").Op("!=").Nil()).Block(
						Return(Nil(), Qual("fmt", "Errorf").Call(Lit(msg+": %w"), Id("err"))),
					)
				}

				f.Func().
					Params(Id("client").Id(fmt.Sprintf("* %sClient", *serv.Name))).
					Id(*method.Name).
					Params(Id("r").Op("*").Add(name(*method.InputType))).
					Params(Op("*").Add(name(*method.OutputType)), Id("error")).
					Block(
						List(Id("input"), Id("err")).Op(":=").Qual("google.golang.org/protobuf/proto", "Marshal").Call(Id("r")),
						errHandler("could not martial request"),
						List(Id("resp"), Id("err")).Op(":=").Id("client.client.Post").Call(
							Qual("fmt", "Sprintf").Call(Lit("%s/"+fmt.Sprintf("%s/%s", *item.Package+"."+*serv.Name, *method.Name)), Id("client.serverURL")),
							Lit("application/octet-stream"),
							Qual("bytes", "NewReader").Call(Id("input")),
						),
						errHandler("error posting request"),
						Defer().Id("resp.Body").Op(".").Id("Close").Call(),
						List(Id("data"), Id("err")).Op(":=").Qual("io/ioutil", "ReadAll").Call(Id("resp.Body")),
						errHandler("error reading response"),
						Id("output").Op(":=").Op("&").Add(name(*method.OutputType)).Block(),
						Id("err").Op("=").Qual("google.golang.org/protobuf/proto", "Unmarshal").Call(Id("data"), Id("output")),
						errHandler("error unmarshalling response"),
						Return(Id("output"), Nil()),
					)
			}
		}

		data := new(string)
		*data = f.GoString()
		outfile.Content = data

		out.File = append(out.File, outfile)
	}

	msg, err := proto.Marshal(&out)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(msg)
}
