package main

import (
	"fmt"
	"path"
	"strings"

	"google.golang.org/protobuf/types/pluginpb"
)

func GenerateValaClient(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)
	for _, f := range d.ProtoFile {
		if len(f.Service) == 0 {
			continue
		}

		file := new(pluginpb.CodeGeneratorResponse_File)
		file.Name = new(string)
		*file.Name = path.Join(strings.TrimSuffix(*f.Name, ".proto")) + ".client.hrpc.vala"

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

		add(`using GLib;`)
		add(`using Soup;`)
		add(`using Json;`)

		for _, service := range f.Service {
			add(`public class %sClient {`, *service.Name)
			indent++
			add(`public bool secure { get; set; }`)
			add(`public string host { get; set; }`)
			add(`public string? authorization { get; set; }`)

			add(`private string unaryPrefix { get { return this.secure ? "https" : "http" ; } }`)
			add(`private string wsPrefix { get { return this.secure ? "wss" : "ws" ; } }`)
			add(`private Soup.Session session;`)

			add(`%sClient(string host, bool secure = false) {
		this.secure = secure;
		this.host = host;
		this.session = new Soup.Session();
	}`, *service.Name)

			for _, meth := range service.Method {
				if meth.GetClientStreaming() && !meth.GetServerStreaming() {
					continue
				} else if meth.GetClientStreaming() && meth.GetServerStreaming() {
					add(`public Variant %s() {
						// TODO
					}`, meth.GetName())
				} else if !meth.GetServerStreaming() && !meth.GetClientStreaming() {
					add(`public Variant %s() {
		var message = new Soup.Message("POST", @"$(this.unaryPrefix)://$(this.host)/%s.%s/%s");
	}`, meth.GetName(), f.GetPackage(), service.GetName(), meth.GetName())
				} else if meth.GetServerStreaming() && !meth.GetClientStreaming() {
					add(`public Variant %s() {}`, meth.GetName())
				}
			}
			indent--
			add(`}`)
		}

		r.File = append(r.File, file)
	}
	return
}
