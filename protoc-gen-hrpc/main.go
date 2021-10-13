// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: MPL-2.0

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
	"unicode"
	"unicode/utf8"

	"github.com/alecthomas/repr"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type QualPair struct {
	Package string
	ID      string
}

func badToUnderscore(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
		return r
	}
	return '_'
}

var isGoKeyword = map[string]bool{
	"break":       true,
	"case":        true,
	"chan":        true,
	"const":       true,
	"continue":    true,
	"default":     true,
	"else":        true,
	"defer":       true,
	"fallthrough": true,
	"for":         true,
	"func":        true,
	"go":          true,
	"goto":        true,
	"if":          true,
	"import":      true,
	"interface":   true,
	"map":         true,
	"package":     true,
	"range":       true,
	"return":      true,
	"select":      true,
	"struct":      true,
	"switch":      true,
	"type":        true,
	"var":         true,
}

func cleanPackageName(name string) string {
	name = strings.Map(badToUnderscore, name)
	// Identifier must not be keyword or predeclared identifier: insert _.
	if isGoKeyword[name] {
		name = "_" + name
	}
	// Identifier must not begin with digit: insert _.
	if r, _ := utf8.DecodeRuneInString(name); unicode.IsDigit(r) {
		name = "_" + name
	}
	return name
}

func goPackageOption(d *descriptorpb.FileDescriptorProto) (impPath string, pkg string, ok bool) {
	opt := d.GetOptions().GetGoPackage()
	if opt == "" {
		return "", "", false
	}
	// A semicolon-delimited suffix delimits the import path and package name.
	sc := strings.Index(opt, ";")
	if sc >= 0 {
		return opt[:sc], cleanPackageName(opt[sc+1:]), true
	}
	// The presence of a slash implies there's an import path.
	slash := strings.LastIndex(opt, "/")
	if slash >= 0 {
		return opt, cleanPackageName(opt[slash+1:]), true
	}
	return "", cleanPackageName(opt), true
}

//go:generate esc -o data_gen.go -pkg main ../templates
func main() {
	input := pluginpb.CodeGeneratorRequest{}
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = proto.Unmarshal(data, &input)
	if err != nil {
		panic(err)
	}

	response := pluginpb.CodeGeneratorResponse{}
	response.SupportedFeatures = new(uint64)
	*response.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, item := range strings.Split(*input.Parameter, ":") {
		builtins := map[string]func(d *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorResponse{
			"qt_cpp_client": GenerateQtCxxClient,
			"d_client":      GenerateDClient,
			"ts_client":     GenerateTSClient,
			"dart_client":   GenerateDartClient,
			"swift_server":  GenerateSwiftServer,
			"elixir_server": GenerateElixirServer,
		}

		if fn, ok := builtins[item]; ok {
			r := fn(&input)

			for _, item := range r.File {
				response.File = append(response.File, item)
			}

			continue
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
			"goFilename": func(item *descriptorpb.FileDescriptorProto, interfix string) string {
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
			},
			"fanciedName": func(item *descriptorpb.FileDescriptorProto) string {
				return strings.Title(strings.ReplaceAll(strings.ReplaceAll(*item.Name, "/", "ᐳ"), ".proto", ""))
			},
			"methodData": func(meth *descriptorpb.MethodDescriptorProto) string {
				data, err := proto.Marshal(meth)
				if err != nil {
					panic(err)
				}
				return repr.String(data)
			},
			"fileData": func(item *descriptorpb.FileDescriptorProto) string {
				data, err := proto.Marshal(item)
				if err != nil {
					panic(err)
				}
				return repr.String(data)
			},
			"serviceData": func(item *descriptorpb.ServiceDescriptorProto) string {
				data, err := proto.Marshal(item)
				if err != nil {
					panic(err)
				}
				return repr.String(data)
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
			},
			"newset": func() map[string]struct{} {
				return map[string]struct{}{}
			},
			"appendSet": func(m map[string]struct{}, v string) string {
				if v == "" {
					return ""
				}
				m[v] = struct{}{}
				return ""
			},
		})

		data, err = ioutil.ReadFile(item)
		if err != nil {
			panic(err)
		}

		tmpl, err = tmpl.Parse(string(data))
		if err != nil {
			panic(err)
		}

		for _, item := range input.ProtoFile {
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
	}

	msg, err := proto.Marshal(&response)
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(msg)
}
