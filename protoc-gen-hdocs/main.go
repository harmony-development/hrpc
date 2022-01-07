// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	packageCommentPath          = 2
	messageCommentPath          = 4
	enumCommentPath             = 5
	serviceCommentPath          = 6
	extensionCommentPath        = 7
	syntaxCommentPath           = 12
	messageFieldCommentPath     = 2
	messageMessageCommentPath   = 3
	messageEnumCommentPath      = 4
	messageExtensionCommentPath = 6
	enumValueCommentPath        = 2
	serviceMethodCommentPath    = 2
)

type CommentData struct {
	LeadingDetached []string

	Leading  string
	Trailing string
}

func removeExcessSpace(str string) string {
	return strings.TrimSpace(strings.ReplaceAll(str, "\n ", "\n"))
}

func commentDataFrom(loc *descriptor.SourceCodeInfo_Location) CommentData {
	detached := make([]string, len(loc.GetLeadingDetachedComments()))
	for i, c := range loc.GetLeadingDetachedComments() {
		detached[i] = removeExcessSpace(c)
	}

	return CommentData{
		Leading:         removeExcessSpace(loc.GetLeadingComments()),
		Trailing:        removeExcessSpace(loc.GetTrailingComments()),
		LeadingDetached: detached,
	}
}

type FilePath struct {
	f *descriptorpb.FileDescriptorProto
	s string
}

type PackageData struct {
	Comments map[FilePath]CommentData
	Messages []DescriptorData
	Services []ServiceData
	Enums    []EnumData
}

type DescriptorData struct {
	*descriptor.DescriptorProto

	Path      string
	FD        *descriptorpb.FileDescriptorProto
	IsRequest bool
}

type ServiceData struct {
	*descriptor.ServiceDescriptorProto

	Path string
	FD   *descriptorpb.FileDescriptorProto
}

type EnumData struct {
	*descriptor.EnumDescriptorProto

	Path string
	FD   *descriptorpb.FileDescriptorProto
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

func getAllEnums(in []*descriptorpb.DescriptorProto, file *descriptorpb.FileDescriptorProto, parent *DescriptorData) (out []EnumData) {
	var rin = []DescriptorData{}

	for mi, msg := range in {
		var path string

		if parent == nil {
			path = fmt.Sprintf("%d.%d", messageCommentPath, mi)
		} else {
			path = fmt.Sprintf("%s.%d.%d", parent.Path, messageCommentPath, mi)
		}

		for ei, enum := range msg.EnumType {
			epath := fmt.Sprintf("%s.%d.%d", path, enumCommentPath, ei)
			out = append(out, EnumData{enum, epath, file})
		}

		n := msg.GetName()

		isReq :=
			len(msg.NestedType) == 0 && parent == nil && (strings.HasSuffix(n, "Request") || strings.HasSuffix(n, "Response"))

		rin = append(rin, DescriptorData{msg, path, file, isReq})
	}

	for _, kind := range rin {
		for _, item := range getAllEnums(kind.NestedType, file, &kind) {
			*item.Name = *kind.Name + "." + *item.Name
			out = append(out, item)
		}
	}

	return
}

func getAllTypes(in []*descriptorpb.DescriptorProto, file *descriptorpb.FileDescriptorProto, parent *DescriptorData) (out []DescriptorData) {
	var rin = []DescriptorData{}

	for i, msg := range in {
		var path string

		if parent == nil {
			path = fmt.Sprintf("%d.%d", messageCommentPath, i)
		} else {
			path = fmt.Sprintf("%s.%d.%d", parent.Path, messageCommentPath, i)
		}

		n := msg.GetName()

		isReq :=
			len(msg.NestedType) == 0 && parent == nil && (strings.HasSuffix(n, "Request") || strings.HasSuffix(n, "Response"))

		rin = append(rin, DescriptorData{msg, path, file, isReq})
	}

	for _, kind := range rin {
		out = append(out, kind)
		for _, item := range getAllTypes(kind.NestedType, file, &kind) {
			*item.Name = *kind.Name + "." + *item.Name
			out = append(out, item)
		}
	}

	return
}

var rex = regexp.MustCompile(`(\w+)=(\w+)`)

func outputMessageType(item *DescriptorData, stuff *PackageData, file *strings.Builder, resolveLink func(s string) string, isSub bool) {
	pfx := ""
	if isSub {
		pfx = "###"
	}
	file.WriteString(fmt.Sprintf("%s## %s%s\n", pfx, iconMessage, item.GetName()))

	comments := stuff.Comments[FilePath{item.FD, item.Path}]
	file.WriteString(comments.Leading)
	file.WriteString("\n\n")

	if len(item.Field) == 0 {
		file.WriteString("This item has no fields.")
	} else {
		if useHTML {
			if isSub {
				file.WriteString(`<span class="h5" aria-level="5">Fields</span>`)
			} else {
				file.WriteString(`<span class="h3" aria-level="3">Fields</span>`)
			}
		} else {
			file.WriteString(pfx + "### Fields\n\n")
		}
	}

	if isSub {
		if useHTML {
			pfx = "###"
		} else {
			pfx = "##"
		}
	} else if useHTML {
		pfx = ""
	}

	file.WriteString("\n")
	for idx, field := range item.Field {
		path := fmt.Sprintf("%s.%d.%d", item.Path, messageFieldCommentPath, idx)
		comments := stuff.Comments[FilePath{item.FD, path}]

		file.WriteString(fmt.Sprintf("%s### %s%s\n", pfx, iconMessageField, field.GetName()))

		label :=
			descriptorpb.FieldDescriptorProto_Label(field.Label.Number())

		isOptional :=
			label == descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL

		isRepeated :=
			label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED

		modifier := ""

		if isOptional {
			modifier = "optional "
		}
		if isRepeated {
			modifier = "repeated "
		}

		file.WriteString("Type: ")

		switch field.GetType() {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			file.WriteString(fmt.Sprintf("%s[%s](%s)", modifier, field.GetTypeName()[1:], resolveLink(field.GetTypeName()[1:])))
		default:
			if v, ok := wellKnown[field.GetType()]; ok {
				file.WriteString(fmt.Sprintf("%s`%s`", modifier, v))
			} else {
				file.WriteString("UNHANDLED | TYPE")
			}
		}

		file.WriteString("\n\n")

		file.WriteString(comments.Leading)
		file.WriteString("\n")
	}
	file.WriteString("\n")
}

var (
	iconMessage      string
	iconMessageField string

	iconEnum      string
	iconEnumValue string

	iconService       string
	iconServiceMethod string
)

var useHTML bool

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

	rgx := rex.FindAllStringSubmatch(gen.GetParameter(), -1)
	res := make(map[string]string)
	for _, kv := range rgx {
		k := kv[1]
		v := kv[2]
		res[k] = v
	}

	codicons :=
		res["codicon"] == "yes"

	useHTML =
		res["usehtml"] == "yes"

	if codicons {
		f := func(icon string) string {
			return fmt.Sprintf(`<span class="codicon codicon-%s %s"></span>`, icon, icon)
		}

		iconMessage = f("symbol-structure")
		iconMessageField = f("symbol-field")
		iconEnum = f("symbol-enum")
		iconEnumValue = f("symbol-enum-member")
		iconService = f("symbol-class")
		iconServiceMethod = f("symbol-method")
	}

	docs := Docs{}

	for _, item := range gen.ProtoFile {
		out := docs[*item.Package]
		if out.Comments == nil {
			out.Comments = map[FilePath]CommentData{}
		}

		for _, location := range item.GetSourceCodeInfo().GetLocation() {
			if location.GetLeadingComments() == "" && location.GetTrailingComments() == "" && len(location.GetLeadingDetachedComments()) == 0 {
				continue
			}

			path := location.GetPath()
			key := make([]string, len(path))
			for idx, p := range path {
				key[idx] = strconv.Itoa(int(p))
			}

			out.Comments[FilePath{item, strings.Join(key, ".")}] = commentDataFrom(location)
		}

		out.Enums = append(out.Enums, getAllEnums(item.MessageType, item, nil)...)
		for idx, enum := range item.EnumType {
			out.Enums = append(out.Enums, EnumData{enum, fmt.Sprintf("%d.%d", enumCommentPath, idx), item})
		}
		out.Messages = append(out.Messages, getAllTypes(item.MessageType, item, nil)...)
		for idx, service := range item.Service {
			out.Services = append(out.Services, ServiceData{service, fmt.Sprintf("%d.%d", serviceCommentPath, idx), item})
		}
		docs[*item.Package] = out
	}

	response := pluginpb.CodeGeneratorResponse{}

	for pkg, stuff := range docs {
		resolveLink := func(s string) string {
			transform := func(s string) string {
				return strings.ReplaceAll(strings.ToLower(s), ".", "")
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

		goto serv

	serv:

		if len(stuff.Services) == 0 {
			goto msg
		}

		file.WriteString("# Services \n\n")
		for idx, serv := range stuff.Services {
			path := fmt.Sprintf("%d.%d", serviceCommentPath, idx)
			comment := stuff.Comments[FilePath{serv.FD, path}]

			file.WriteString(fmt.Sprintf("## %s%s\n\n", iconService, serv.GetName()))

			file.WriteString(comment.Leading)
			file.WriteString("\n")

			if useHTML {
				file.WriteString(`<span class="h3" aria-level="3">Fields</span>` + "\n")
			} else {
				file.WriteString("### Methods\n\n")
			}

			for idx, method := range serv.Method {
				file.WriteString(fmt.Sprintf("#### %s%s\n", iconServiceMethod, method.GetName()))

				cStream := method.GetClientStreaming()
				sStream := method.GetServerStreaming()

				cType := fmt.Sprintf("[%s](%s)", method.GetInputType()[1:], resolveLink(method.GetInputType()[1:]))
				sType := fmt.Sprintf("[%s](%s)", method.GetOutputType()[1:], resolveLink(method.GetOutputType()[1:]))

				if cStream {
					file.WriteString("streaming ")
				}
				file.WriteString(cType)

				file.WriteString(" -> ")

				if sStream {
					file.WriteString("streaming ")
				}
				file.WriteString(sType)

				file.WriteString("\n\n")

				mpath := fmt.Sprintf("%s.%d.%d", path, serviceMethodCommentPath, idx)
				comment := stuff.Comments[FilePath{serv.FD, mpath}]

				file.WriteString(comment.Leading)
				file.WriteString("\n")

				fstOK := false

				for _, item := range stuff.Messages {
					if !item.IsRequest {
						continue
					}

					if item.FD.GetPackage()+"."+item.GetName() == method.GetInputType()[1:] {
						fstOK = true
						file.WriteString("\n<br/>\n\n")
						outputMessageType(&item, &stuff, &file, resolveLink, true)
						break
					}
				}
				for _, item := range stuff.Messages {
					if !item.IsRequest {
						continue
					}

					if item.FD.GetPackage()+"."+item.GetName() == method.GetOutputType()[1:] {
						if fstOK {
							file.WriteString("\n<br/>\n\n")
						}
						outputMessageType(&item, &stuff, &file, resolveLink, true)
						break
					}
				}

				if idx < len(serv.Method)-1 {
					file.WriteString("------\n")
				}
			}
		}

	msg:

		if len(stuff.Messages) == 0 {
			goto enum
		}

		file.WriteString("# Standalone Message Types \n\n")
		for idx, item := range stuff.Messages {
			if item.IsRequest {
				continue
			}
			outputMessageType(&item, &stuff, &file, resolveLink, false)
			if idx < len(stuff.Messages)-1 {
				file.WriteString("------\n")
			}
		}

	enum:
		if len(stuff.Enums) == 0 {
			goto after
		}

		file.WriteString("# Enums \n\n")
		for idx, enum := range stuff.Enums {
			comment := stuff.Comments[FilePath{enum.FD, enum.Path}]

			file.WriteString(fmt.Sprintf("## %s%s\n\n", iconEnum, enum.GetName()))

			file.WriteString(comment.Leading)
			file.WriteString("\n")

			for idx, value := range enum.Value {
				vpath := fmt.Sprintf("%s.%d.%d", enum.Path, enumValueCommentPath, idx)
				comment := stuff.Comments[FilePath{enum.FD, vpath}]

				file.WriteString(fmt.Sprintf("### %s%s\n", iconEnumValue, value.GetName()))

				file.WriteString(comment.Leading)
				file.WriteString("\n\n")
			}

			if idx < len(stuff.Enums)-1 {
				file.WriteString("------\n")
			}
		}

	after:

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
