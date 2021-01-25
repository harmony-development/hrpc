package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type QualPair struct {
	Package string
	ID      string
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

	fname := string("")
	setFName := func(f string) string {
		fname = f
		return ""
	}

	ok := false
	setOK := func(f bool) string {
		ok = f
		return ""
	}

	tmpl := template.New("codegen")
	tmpl.Funcs(template.FuncMap{
		"setFilename":  setFName,
		"trimSuffix":   strings.TrimSuffix,
		"sprintf":      fmt.Sprintf,
		"filepathBase": filepath.Base,
		"setOK":        setOK,
		"deref": func(v interface{}) interface{} {
			return reflect.ValueOf(v).Elem().Interface()
		},
		"hasClientStream": func(v *descriptorpb.MethodDescriptorProto) bool {
			return v.ClientStreaming != nil && *v.ClientStreaming
		},
		"hasServerStream": func(v *descriptorpb.MethodDescriptorProto) bool {
			return v.ServerStreaming != nil && *v.ServerStreaming
		},
		"resolvedGoType": func(item *descriptorpb.FileDescriptorProto, method *descriptorpb.MethodDescriptorProto, in string) QualPair {
			wellKnown := map[string]QualPair{
				".google.protobuf.Empty": {
					Package: "github.com/golang/protobuf/ptypes/empty",
					ID:      "empty.Empty",
				},
			}

			if val, ok := wellKnown[in]; ok {
				return val
			}

			split := strings.Split(*method.InputType, ".")
			final := split[len(split)-1]

			return QualPair{
				Package: *item.Options.GoPackage,
				ID:      path.Base(*item.Options.GoPackage) + "." + final,
			}
		},
		"newset": func() map[string]struct{} {
			return map[string]struct{}{}
		},
		"appendSet": func(m map[string]struct{}, v string) string {
			m[v] = struct{}{}
			return ""
		},
	})

	data, err = ioutil.ReadFile(*gen.Parameter)
	if err != nil {
		panic(err)
	}

	tmpl, err = tmpl.Parse(string(data))
	if err != nil {
		panic(err)
	}

	response := pluginpb.CodeGeneratorResponse{}

	for _, item := range gen.ProtoFile {
		out := strings.Builder{}
		ok = false
		err = tmpl.Execute(&out, item)
		if err != nil {
			panic(err)
		}

		if !ok {
			continue
		}

		outfile := new(pluginpb.CodeGeneratorResponse_File)
		outfile.Name = new(string)
		*outfile.Name = fname

		data := new(string)
		*data = out.String()
		outfile.Content = data

		response.File = append(response.File, outfile)

		continue
	}

	msg, err := proto.Marshal(&response)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(msg)
}
