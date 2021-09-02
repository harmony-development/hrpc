package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	errorPackage   = protogen.GoImportPath("errors")
	contextPackage = protogen.GoImportPath("context")
	httpPackage    = protogen.GoImportPath("fasthttp")
	protoPackage   = protogen.GoImportPath("google.golang.org/protobuf/proto")
	serverPackage  = protogen.GoImportPath("github.com/harmony-development/hrpc/server")
)

func GenerateGoServer(gen *protogen.Plugin, file *protogen.File) {
	if len(file.Services) == 0 {
		return
	}

	filename := file.GeneratedFilenamePrefix + "_hrpc.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	g.P("// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	genContent(g, file)
}

func genContent(g *protogen.GeneratedFile, file *protogen.File) {
	for _, service := range file.Services {
		genService(g, service)
	}
}

func genService(g *protogen.GeneratedFile, service *protogen.Service) {
	serverType := service.GoName + "Server"
	g.P("type ", serverType, " interface {")
	for _, method := range service.Methods {
		g.P(method.Comments.Leading, serverSignature(g, method))
	}
	g.P("}")
	g.P()
	genDefaultImpl(g, service)
	genHandlerStruct(g, service, serverType)
}

func genDefaultImpl(g *protogen.GeneratedFile, service *protogen.Service) {
	errorNewType := g.QualifiedGoIdent(errorPackage.Ident("New"))
	dummyType := "Default" + service.GoName
	g.P("type ", dummyType, " struct {}")
	for _, method := range service.Methods {
		g.P("func (", dummyType, ") ", serverSignature(g, method), " {")
		g.P("return nil, ", errorNewType, `("unimplemented")`)
		g.P("}")
	}
}

func genHandlerStruct(g *protogen.GeneratedFile, service *protogen.Service, serverType string) {
	handlerType := service.GoName + "Handler"
	handlerFuncType := g.QualifiedGoIdent(serverPackage.Ident("RawHandler"))
	handlerMapType := fmt.Sprintf("map[string]%s", handlerFuncType)

	g.P(fmt.Sprintf("type %s struct {", handlerType))
	g.P("Server " + serverType)
	g.P("}")
	g.P("func New", handlerType, "(server ", serverType, ") *", handlerType, " {")
	g.P("return &", handlerType, "{ Server: server }")
	g.P("}")
	g.P(fmt.Sprintf("func (h *%s) Name() string {", handlerType))
	g.P(`return "`, service.GoName, `"`)
	g.P("}")
	g.P(fmt.Sprintf("func (h *%s) Routes() %s {", handlerType, handlerMapType))
	g.P("return ", handlerMapType, "{")
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		g.P(
			fmt.Sprintf(
				`"/%s/": %s,`,
				method.Desc.FullName(),
				genRawHandler(g, service, method),
			),
		)
	}
	g.P("}")
	g.P("}")
	g.P()
}

func genRawHandler(g *protogen.GeneratedFile, service *protogen.Service, m *protogen.Method) string {
	newUnaryHandlerType := g.QualifiedGoIdent(serverPackage.Ident("NewUnaryHandler"))
	messageType := g.QualifiedGoIdent(protoPackage.Ident("Message"))
	return fmt.Sprintf(
		"%s(&%s{}, func(c context.Context, req %s) (%s, error) { return h.Server.%s(c, req.(*%s)) })",
		newUnaryHandlerType,
		g.QualifiedGoIdent(m.Input.GoIdent),
		messageType,
		messageType,
		m.GoName,
		g.QualifiedGoIdent(m.Input.GoIdent),
	)
}

func serverSignature(g *protogen.GeneratedFile, m *protogen.Method) string {
	var inputArgs []string
	var ret string
	inputType := g.QualifiedGoIdent(m.Input.GoIdent)
	outputType := g.QualifiedGoIdent(m.Output.GoIdent)
	inputArgs = append(inputArgs, g.QualifiedGoIdent(contextPackage.Ident("Context")))
	if m.Desc.IsStreamingClient() {
		inputArgs = append(inputArgs, "chan *"+inputType)
	} else {
		inputArgs = append(inputArgs, "*"+inputType)
	}
	if m.Desc.IsStreamingServer() {
		ret = "chan *" + outputType
	} else {
		ret = "*" + outputType
	}
	return fmt.Sprintf("%s(%s) (%s, error)", m.GoName, strings.Join(inputArgs, ", "), ret)
}
