package main

import (
	"path"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"

	. "github.com/dave/jennifer/jen"
)

func resolvedGoType(item *descriptorpb.FileDescriptorProto, method *descriptorpb.MethodDescriptorProto, in string) QualPair {
	wellKnown := map[string]QualPair{
		".google.protobuf.Empty": {
			Package: "github.com/golang/protobuf/ptypes/empty",
			ID:      "empty.Empty",
		},
	}

	if val, ok := wellKnown[in]; ok {
		return val
	}

	split := strings.Split(in, ".")
	final := split[len(split)-1]

	if strings.Join(split[:len(split)-1], ".") == "."+*item.Package {
		return QualPair{
			Package: "",
			ID:      final,
		}
	}

	return QualPair{
		Package: *item.Options.GoPackage,
		ID:      path.Base(*item.Options.GoPackage) + "." + final,
	}
}

func goFilename(item *descriptorpb.FileDescriptorProto, interfix string) string {
	name := *item.Name
	if ext := path.Ext(name); ext == ".proto" || ext == ".protodevel" {
		name = name[:len(name)-len(ext)]
	}
	name += ".hrpc." + interfix + ".go"

	if importPath, _, ok := goPackageOption(item); ok && importPath != "" {
		_, name = path.Split(name)
		name = path.Join(importPath, name)
		return name
	}

	return name
}

func GenerateGoClient(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)
	for _, fi := range d.ProtoFile {
		if len(fi.Service) == 0 {
			continue
		}

		file := new(pluginpb.CodeGeneratorResponse_File)
		file.Name = new(string)
		*file.Name = goFilename(fi, "client")

		f := NewFile(filepath.Base(*fi.Options.GoPackage))

		for _, serv := range fi.Service {
			f.Type().Id(*serv.Name+"Client").Struct(
				Id("client").Id("*").Qual("net/http", "Client"),
				Id("serverURL").String(),
				Id("Header").Qual("net/http", "Header"),
				Id("HTTPProto").String(),
				Id("WSProto").String(),
			)

			iffer := func(id, y, n string) *Statement {
				return If(Id("secure")).Block(
					Id(id).Op("=").Lit(y),
				).Else().Block(
					Id(id).Op("=").Lit(n),
				)
			}

			f.Func().Id("New"+*serv.Name).Params(
				Id("url").String(),
				Id("secure").Bool(),
			).Id("*"+*serv.Name+"Client").Block(
				Var().Id("httpproto").String(),
				iffer("httpproto", "https", "http"),
				Var().Id("wsproto").String(),
				iffer("wsproto", "wss", "ws"),
				Return(Op("&").Id(*serv.Name+"Client").Values(Dict{
					Id("client"):    Op("&").Qual("net/http", "Client").Block(),
					Id("serverURL"): Id("url"),

					Id("Header"):    Qual("net/http", "Header").Block(),
					Id("HTTPProto"): Id("httpproto"),
					Id("WSProto"):   Id("wsproto"),
				})),
			)

			for _, fun := range serv.Method {
				e := func() *Statement {
					return If(Id("err").Op("!=").Nil()).Block(
						Return(Nil(), Qual("fmt", "Errorf").Call(Lit("error reading response: %w"), Id("err"))),
					)
				}

				// bidi
				if *fun.ClientStreaming && *fun.ServerStreaming {

					// uni serv
				} else if *fun.ServerStreaming {

					// unary request
				} else {

				}
			}
		}

		file.Content = new(string)
		*file.Content = f.GoString()
		r.File = append(r.File, file)
	}
	return
}
