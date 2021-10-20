// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type PackageData struct {
	Messages []*descriptor.DescriptorProto
	Services []*descriptor.ServiceDescriptorProto
}

type Docs map[string]PackageData

var wellKnown = map[descriptor.FieldDescriptorProto_Type]string{
	descriptor.FieldDescriptorProto_TYPE_DOUBLE:   "double",
	descriptor.FieldDescriptorProto_TYPE_FLOAT:    "float",
	descriptor.FieldDescriptorProto_TYPE_INT64:    "int64",
	descriptor.FieldDescriptorProto_TYPE_UINT64:   "uint64",
	descriptor.FieldDescriptorProto_TYPE_INT32:    "int32",
	descriptor.FieldDescriptorProto_TYPE_FIXED64:  "fixed64",
	descriptor.FieldDescriptorProto_TYPE_FIXED32:  "fixed32",
	descriptor.FieldDescriptorProto_TYPE_BOOL:     "bool",
	descriptor.FieldDescriptorProto_TYPE_STRING:   "string",
	descriptor.FieldDescriptorProto_TYPE_BYTES:    "bytes",
	descriptor.FieldDescriptorProto_TYPE_UINT32:   "uint32",
	descriptor.FieldDescriptorProto_TYPE_SFIXED32: "sfixed32",
	descriptor.FieldDescriptorProto_TYPE_SFIXED64: "sfixed64",
	descriptor.FieldDescriptorProto_TYPE_SINT32:   "sint32",
	descriptor.FieldDescriptorProto_TYPE_SINT64:   "sint64",
	// descriptor.FieldDescriptorProto_TYPE_GROUP
	// descriptor.FieldDescriptorProto_TYPE_MESSAGE
	// descriptor.FieldDescriptorProto_TYPE_ENUM
}

func getAllTypes(in []*descriptorpb.DescriptorProto) (out []*descriptorpb.DescriptorProto) {
	out = append(out, in...)

	for _, kind := range in {
		for _, item := range getAllTypes(kind.NestedType) {
			*item.Name = *kind.Name + "." + *item.Name
			out = append(out, item)
		}
	}

	return
}

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

	docs := Docs{}

	for _, item := range gen.ProtoFile {
		out := docs[*item.Package]
		out.Messages = append(out.Messages, getAllTypes(item.MessageType)...)
		out.Services = append(out.Services, item.Service...)
		docs[*item.Package] = out
	}

	response := pluginpb.CodeGeneratorResponse{}

	for pkg, stuff := range docs {
		resolveLink := func(s string) string {
			transform := func(s string) string {
				return strings.ReplaceAll(strings.ToLower(s), ".", "-")
			}
			if strings.HasPrefix(s, pkg) {
				return fmt.Sprintf("#%s", transform(strings.TrimPrefix(s, pkg)[1:]))
			}

			split := strings.Split(s, ".")
			prefix := strings.Join(split[:len(split)-1], ".")

			return fmt.Sprintf(`{{< ref "%s.md" >}}#%s`, prefix, transform(split[len(split)-1]))
		}

		file := strings.Builder{}

		file.WriteString("---\n")
		file.WriteString(fmt.Sprintf("title: \"Reference: %s\"\n", pkg))
		file.WriteString("---\n")

		file.WriteString("## Message Types \n\n")
		for _, item := range stuff.Messages {
			file.WriteString(fmt.Sprintf("### %s\n", item.GetName()))
			file.WriteString("\nFields\n\n")
			file.WriteString("| Name | Type |\n")
			file.WriteString("| ---- | ---- |\n")
			for _, field := range item.Field {
				file.WriteString("| ")
				switch field.GetType() {
				case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
					file.WriteString(fmt.Sprintf("%s | [%s](%s)", field.GetName(), field.GetTypeName()[1:], resolveLink(field.GetTypeName()[1:])))
				default:
					if v, ok := wellKnown[field.GetType()]; ok {
						file.WriteString(fmt.Sprintf("%s | `%s`", field.GetName(), v))
					} else {
						file.WriteString("UNHANDLED | TYPE")
					}
				}
				file.WriteString(" |\n")
			}
			file.WriteString("\n")
		}

		file.WriteString("## Services \n\n")
		for _, serv := range stuff.Services {
			file.WriteString(fmt.Sprintf("### %s\n\n", serv.GetName()))

			file.WriteString("#### Unary Methods\n\n")

			file.WriteString("| Name | Request | Response |\n")
			file.WriteString("| ---- | ------- | -------- |\n")

			for _, method := range serv.Method {
				if method.GetClientStreaming() || method.GetServerStreaming() {
					continue
				}

				file.WriteString("|")
				file.WriteString(*method.Name)
				file.WriteString("|")
				file.WriteString(fmt.Sprintf("[%s](%s)", method.GetInputType()[1:], resolveLink(method.GetInputType()[1:])))
				file.WriteString("|")
				file.WriteString(fmt.Sprintf("[%s](%s)", method.GetOutputType()[1:], resolveLink(method.GetOutputType()[1:])))
				file.WriteString("|\n")
			}
			file.WriteString("\n")

			file.WriteString("#### Streaming Methods\n\n")

			file.WriteString("| Name | Client Streams | Server Streams |\n")
			file.WriteString("| ---- | -------------- | -------------- |\n")

			for _, method := range serv.Method {
				if !method.GetClientStreaming() || !method.GetServerStreaming() {
					continue
				}

				file.WriteString("|")
				file.WriteString(*method.Name)
				file.WriteString("|")
				file.WriteString(fmt.Sprintf("[%s](%s)", method.GetInputType()[1:], resolveLink(method.GetInputType()[1:])))
				file.WriteString("|")
				file.WriteString(fmt.Sprintf("[%s](%s)", method.GetOutputType()[1:], resolveLink(method.GetOutputType()[1:])))
				file.WriteString("|\n")
			}
		}

		n := new(string)
		*n = pkg + ".md"
		c := new(string)
		*c = file.String()

		response.File = append(response.File, &pluginpb.CodeGeneratorResponse_File{
			Name:    n,
			Content: c,
		})
	}

	msg, err := proto.Marshal(&response)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(msg)
}
