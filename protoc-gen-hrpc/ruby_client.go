package main

import (
	"fmt"
	"path"
	"strings"
	"regexp"

	"google.golang.org/protobuf/types/pluginpb"
)

func GenerateRubyClient(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)
	for _, f := range d.ProtoFile {
		if len(f.Service) == 0 {
			continue
		}

		file := new(pluginpb.CodeGeneratorResponse_File)
		file.Name = new(string)
		*file.Name = path.Join(strings.TrimSuffix(*f.Name, ".proto")) + "_hrpc.rb"

		dat := strings.Builder{}
		indent := 0
		add := func(format string, v ...interface{}) {
			dat.WriteString(strings.Repeat("  ", indent))
			dat.WriteString(fmt.Sprintf(format, v...))
			dat.WriteRune('\n')
		}

		empty := func(input ...int) {
			repeat := 1
			if (len(input) != 0) {
				repeat = input[0]
			}
			for i := 0; i < repeat; i++ {
				dat.WriteRune('\n')
			}
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

		for _, dep := range f.Dependency {
			add(`require_relative '../../%s'`, strings.TrimSuffix(dep, ".proto")+"_pb")
		}
		add(`require_relative '%s'`, strings.TrimSuffix(path.Base(*f.Name), ".proto")+"_pb")
		add(`require 'net/http'`)
		add(`require 'websocket-eventmachine-client'`)
		empty()

		kind := func(in string) string {
			split := strings.Split(in, ".")
			for i := 0; i < len(split) - 1; i++ {
				split[i] = strings.Title(split[i])
			}
			return strings.Join(split,"::")
		}

		re := regexp.MustCompile("(\\B[A-Z]+)")

		for _, service := range f.Service {
			add(`module %s`, *service.Name)
			indent++
			add(`class Client`)
			indent++
			add(`def initialize(host, port = 2289, secure = true)`)
			indent++
			add(`@host = host`)
			add(`@pro = secure ? 'https' : 'http'`)
			add(`@uri = "#{host}:#{port}"`)
			add(`@https = Net::HTTP.new(host, port)`)
			add(`@https.use_ssl = secure`)
			indent--
			add(`end`)
			for _, meth := range service.Method {
				if !meth.GetServerStreaming() && !meth.GetClientStreaming() {
					empty()
					add(`def %s(input = {}, headers = {})`, strings.ToLower(re.ReplaceAllString(*meth.Name, "_$1")))
					indent++
					{
						add(`uri = URI("#{@pro}://#{@uri}/%s.%s/%s")`, *f.Package, *service.Name, *meth.Name)
						add(`req = Net::HTTP::Post.new(uri)`)
						add(`body = %s.new(input)`, kind(*meth.InputType))
						add(`req.body = %s.encode(body)`, kind(*meth.InputType))
						add(`req['content-type'] = 'application/hrpc'`)
						add(`headers.each { |k, v| req[k] = v }`)
						add(`res = @https.request(req)`)
						add(`raise res if res.code != '200'`)
						add(`%s.decode(res.body)`, kind(*meth.OutputType))
					}
					indent--
					add(`end`)
				}
			}
			indent--
			add(`end`)
			empty()
			for _, meth := range service.Method {
				if meth.GetServerStreaming() && !meth.GetClientStreaming() {
					add(`class %sSocket < WebSocket::EventMachine::Client`, *meth.Name)
					indent++
					{
						add(`def self.connect(host, port = 2289, secure = true, headers = {})`)
						indent++
							add(`pro = secure ? 'wss' : 'ws'`)
							add(`super(uri: "#{pro}://#{host}:#{port}/%s.%s/%s",`, *f.Package, *service.Name, *meth.Name)
							add(`      headers: { 'Sec-WebSocket-Protocol' => 'harmony' }.merge(headers)`)
							add(`)`)
						indent--
						add(`end`)
						empty()

						add(`def onmessage(&blk)`)
						indent++
							add(`super do |msg, type|`)
							indent++
								add(`yield %s.decode(msg), type`, kind(*meth.OutputType))
							indent--
							add(`end`)
						indent--
						add(`end`)
						empty()

						add(`def submit(data, args = {})`)
						indent++
							add(`body = %s.new(data)`, kind(*meth.InputType))
							add(`send(%s.encode(body),`, kind(*meth.InputType))
							add(`     { type: :binary }.merge(args))`)
						indent--
						add(`end`)
					}
					indent--
					add(`end`)
				}
			}
			indent--
			add(`end`)
		}

		r.File = append(r.File, file)
	}
	return
}
