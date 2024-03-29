# hRPC

## A simple RPC system using protocol buffers over HTTP and WebSockets

hRPC was designed to be used as a very simple alternative to gRPC. gRPC streams
are known to be problematic due to how they were designed to be long-living, so
as a result hRPC was born.

### Supported Languages

hRPC was made with the intention to be as **simple as possible**, so creating a
code generator / transport isn't actually that scary. That said, we have already
built with hRPC using the following languages:

- Rust (using [hrpc-rs](https://github.com/harmony-development/hrpc-rs))
- TypeScript (using [@harmony-dev/hrpc](https://github.com/harmony-development/protobuf-ts-transport-hrpc))
- C++ (using `protoc-gen-hrpc` with `--hrpc_opt="cpp_client"`)
- C# (using `protoc-gen-hrpc` with `--hrpc-opt="csharp_client"` or `--hrpc-opt="csharp_server"`, and [hrpc.net](https://github.com/harmony-development/hrpc.net) wrapper)
- Dart (using `protoc-gen-hrpc` with `--hrpc_opt="dart_client"`)
- Elixir (using `protoc-gen-hrpc` with `--hrpc_opt="elixir_server"`)

### WIP Languages

- Go (using `protoc-gen-go-hrpc` **not `protoc-gen-hrpc`**)

### Examples

Example applications are located at [the hRPC examples repository](https://github.com/harmony-development/hrpc-examples).

Also check out our [blog post] containing a tutorial.

### Protocol and Spec

See the protobuf files in [`protocol`](./protocol). That folder also contains
documentation regarding to hRPC.

[blog post]: https://dev.to/harmonydevelopment/introducing-hrpc-a-simple-rpc-system-for-user-facing-apis-16ge
