// SPDX-FileCopyrightText: 2021 Carson Black <uhhadd@gmail.com>
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"path"
	"strings"

	"google.golang.org/protobuf/types/pluginpb"
)

// import Vapor
func GenerateSwiftServer(d *pluginpb.CodeGeneratorRequest) (r *pluginpb.CodeGeneratorResponse) {
	r = new(pluginpb.CodeGeneratorResponse)

	for _, f := range d.ProtoFile {
		if len(f.Service) == 0 {
			continue
		}

		file := new(pluginpb.CodeGeneratorResponse_File)
		file.Name = new(string)
		*file.Name = path.Join(strings.TrimSuffix(*f.Name, ".proto")) + ".hrpc.swift"

		dat := strings.Builder{}
		indent := 0
		add := func(format string, a ...interface{}) {
			dat.WriteString(strings.Repeat("\t", indent))
			dat.WriteString(fmt.Sprintf(format, a...))
			dat.WriteRune('\n')
		}
		addI := func(format string, a ...interface{}) {
			add(format, a...)
			indent++
		}
		addD := func(format string, a ...interface{}) {
			indent--
			add(format, a...)
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

		add(`import Vapor`)
		add(`import SwiftProtobuf`)
		add(`import Foundation`)

		swiftName := func(kind string) string {
			split := strings.Split(kind[1:], ".")
			for idx := range split {
				split[idx] = strings.Title(split[idx])
			}
			return strings.Join(split, "_")
		}

		for _, service := range f.Service {
			addI(`protocol %sServer {`, service.GetName())

			for _, rpc := range service.Method {
				if rpc.GetClientStreaming() && rpc.GetServerStreaming() {
					add(`func %s (req: Request, out: (%s) -> Void) -> (%s) -> Void`, rpc.GetName(), swiftName(rpc.GetOutputType()), swiftName(rpc.GetInputType()))
				} else if rpc.GetServerStreaming() && !rpc.GetClientStreaming() {
					add(`func %s (req: Request, in: %s, out: (%s) -> Void)`, rpc.GetName(), swiftName(rpc.GetInputType()), swiftName(rpc.GetOutputType()))
				} else {
					add(`func %s (req: Request, in: %s) async throws -> %s`, rpc.GetName(), swiftName(rpc.GetInputType()), swiftName(rpc.GetOutputType()))
				}
			}

			addD(`}`)

			addI(`extension %sServer {`, service.GetName())

			for _, rpc := range service.Method {
				if rpc.GetClientStreaming() && rpc.GetServerStreaming() {
					add(`func %s (req: Request, out: (%s) -> Void) -> (%s) -> Void { return { _ in } }`, rpc.GetName(), swiftName(rpc.GetOutputType()), swiftName(rpc.GetInputType()))
				} else if rpc.GetServerStreaming() && !rpc.GetClientStreaming() {
					add(`func %s (req: Request, in: %s, out: (%s) -> Void) { }`, rpc.GetName(), swiftName(rpc.GetInputType()), swiftName(rpc.GetOutputType()))
				} else {
					add(`func %s (req: Request, in: %s) async throws -> %s { throw Abort(.internalServerError, reason: "unimplemented") }`, rpc.GetName(), swiftName(rpc.GetInputType()), swiftName(rpc.GetOutputType()))
				}
			}

			addI(`func registerRoutes(withBuilder builder: RoutesBuilder) {`)

			addI(`builder.group("%s.%s") { builder in`, f.GetPackage(), service.GetName())

			for _, rpc := range service.Method {
				route := rpc.GetName()

				if rpc.GetClientStreaming() && rpc.GetServerStreaming() {
					addI(`builder.webSocket("%s") { request, ws in`, route)
					{

						addI(`let callback = self.%s(req: request) { message in`, rpc.GetName())
						{
							addI(`do {`)
							{
								add(`ws.send([UInt8](try message.serializedData()))`)
							}
							addD(`} catch {`)
							indent++
							{
								add(`_ = ws.close()`)
							}
							addD(`}`)
						}
						addD(`}`)

						addI(`ws.onBinary { ws, bb in`)
						{
							addI(`do {`)
							{
								add(`try callback(%s(serializedData: bb.allData()))`, swiftName(rpc.GetInputType()))
							}
							addD(`} catch {`)
							indent++
							{
								add(`_ = ws.close()`)
							}
							addD(`}`)
						}
						addD(`}`)
					}
					addD(`}`)
				} else if rpc.GetServerStreaming() && !rpc.GetClientStreaming() {

					addI(`builder.webSocket("%s") { request, ws in`, route)
					{
						addI(`ws.onBinary { ws, bb in`)
						{
							addI(`do {`)
							{
								add(`let message = try %s(serializedData: bb.allData())`, swiftName(rpc.GetInputType()))

								addI(`self.%s(req: request, in: message) { message in`, rpc.GetName())
								{
									addI(`do {`)
									{
										add(`ws.send([UInt8](try message.serializedData()))`)
									}
									addD(`} catch {`)
									indent++
									{
										add(`_ = ws.close()`)
									}
									addD(`}`)
								}
								addD(`}`)
							}
							addD(`} catch {`)
							indent++
							{
								add(`_ = ws.close()`)
							}
							addD(`}`)
						}
						addD(`}`)
					}
					addD(`}`)

				} else {
					addI(`builder.post("%s") { request -> Response in`, route)

					addI(`do {`)

					add(`let message: %s = try request.decodeMessage()`, swiftName(rpc.GetInputType()))
					add(`let response = try await self.%s(req: request, in: message)`, rpc.GetName())
					add(`return try response.toResponse(on: request)`)

					addD(`} catch {`)

					addI(`throw Abort(.internalServerError, reason: "something did an oops \(error)")`)

					addD(`}`)

					addD(`}`)
				}
			}

			addD(`}`)

			addD(`}`)

			addD(`}`)
		}

		r.File = append(r.File, file)
	}
	return
}
