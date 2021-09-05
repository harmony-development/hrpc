package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

type GeneratorFunc = func(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile

var generators = map[string]GeneratorFunc{}

func main() {
	var flags flag.FlagSet
	client := flags.Bool("client", true, "generate client code")
	server := flags.Bool("server", true, "generate server code")
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if *client {
				GenerateGoClient(gen, f)
			}
			if *server {
				GenerateGoServer(gen, f)
			}
		}
		return nil
	})
}
