# hRPC

## A simple RPC system using protocol buffers over HTTP and WebSockets

hRPC was designed to be used as a very simple alternative to gRPC. gRPC streams are known to be problematic due to how they were designed to be long-living, so as a result hRPC was born.

### Supported Languages

hRPC was made with the intention to be as **simple as possible**, so creating a code generator / transport isn't actually that scary.
That said, we have already built with hRPC using the following languages:

- Go (using `protoc-gen-go-hrpc` **not `protoc-gen-hrpc`**)
- Rust (using [hrpc-rs](https://github.com/harmony-development/hrpc-rs))
- TypeScript (using [@harmony-dev/hrpc](https://github.com/harmony-development/protobuf-ts-transport-hrpc))
- C++ (using `protoc-gen-hrpc` with `--hrpc_opt="cpp_client"`)
- Dart (using `protoc-gen-hrpc` with `--hrpc_opt="dart_client"`)
- Elixir (using `protoc-gen-hrpc` with `--hrpc_opt="elixir_server"`)
