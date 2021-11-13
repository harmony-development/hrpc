# The hRPC Specification

This document defines and explains design details for hRPC.

## Version

This is hRPC specification version `1`.

## Protobuf Version

All hRPC implementations **MUST** support [Protobuf 3](https://developers.google.com/protocol-buffers/docs/reference/proto3-spec).

All references to Protobuf in this document are references to Protobuf 3.

## RPCs

A unary RPC is where a Protobuf RPC signature is as follows:
```protobuf
rpc ExampleMethod(ExampleMessage) returns (ExampleMessage);
```

A streaming RPC is where a Protobuf RPC signature is as follows:
```protobuf
rpc ExampleMethod(stream ExampleMessage) returns (stream ExampleMessage);
```

A client streaming RPC is where a Protobuf RPC signature is as follows:
```protobuf
rpc ExampleMethod(stream ExampleMessage) returns (ExampleMessage);
```
In client streaming RPCs, the server is only allowed to send **one** message.
Messages sent after that **SHOULD** be ignored by the client.


A server streaming RPC is where a Protobuf RPC signature is as follows:
```protobuf
rpc ExampleMethod(ExampleMessage) returns (stream ExampleMessage);
```
In server streaming RPCs, the client is only allowed to send **one** message.
Messages sent after that **SHOULD** be ignored by the server.

### Errors

Servers can send errors to clients using the [hRPC error] protocol type. How
errors sent are dependent on the transport used. Error identifiers are
described in the [hRPC errors document].

## Transports

hRPC is transport agnostic. Here we describe transport specific design details.

### HTTP

This transport works over HTTP.

#### Request Paths

A request path is a string in the following format, where `<package>` is the
name of the package, `<service>` is the name of the service and `<rpc>` is
the name of the RPC:
```
/<package>.<service>/<rpc>
```

If `<package>` is empty, it becomes the following:
```
/<service>/<rpc>
```

This is used as the path of an URI while making a request to a server.

##### Example

For the Protobuf definition below:
```protobuf
syntax = "proto3";

package example;

message ExampleMessage { }

service ExampleService {
    rpc ExampleMethod(ExampleMessage) returns (ExampleMessage);
}
```

The `ExampleMethod` RPC will have a request path of:
```
/example.ExampleService/ExampleMethod
```

#### Unary RPCs

##### Unary requests

- **MUST** have method set to `POST`.
- **MUST** contain the serialized binary data of the Protobuf message in their
body.
- **MUST** set the [`Content-Type`][http headers] header to `application/hrpc`.
- **MUST** set the [`Content-Length`][http headers] header to the length of
the serialized binary data.
- **SHOULD** add `<current spec version>` to `Hrpc-Version` header, where
`<current spec version>` is the [version defined at the start of this document](#Version).

##### Unary responses

- **MUST** contain the serialized binary data of the Protobuf message in their
body.
- **MUST** set the [`Content-Type`][http headers] header to `application/hrpc`.
- **MUST** add `<current spec version>` to `Hrpc-Version` header, where
`<current spec version>` is the [version defined at the start of this document](#Version).

##### Client behaviour

- After getting a successful unary response, a client **SHOULD** look for the
`Hrpc-Version` header. A client then **SHOULD** check if it can work with the
version of the specification in the header. A client **SHOULD** return an error
to the user notifying of the incompatible spec version.

#### Streaming RPCs

Streaming RPCs use a [`WebSocket`][websocket] to communicate Protobuf messages
between server and client. The initial handshake is done through HTTP.

##### Handshake requests

- **MUST** have method set to `GET`.
- **MUST** contain `hrpc` in the [`Sec-WebSocket-Protocol`][websocket_protocol_header] header.
- **SHOULD** add `hrpc-version=<current spec version>` to [`Sec-WebSocket-Extensions`][websocket_extensions_header] header,
where `<current spec version>` is the [version defined at the start of this document](#Version).

##### Handshake responses

- **IF** the request had `hrpc` in the [`Sec-WebSocket-Protocol`][websocket_protocol_header] header,
**MUST** set the same header to `hrpc`.
- **MUST** add `hrpc-version=<current spec version>` to [`Sec-WebSocket-Extensions`][websocket_extensions_header] header,
where `<current spec version>` is the [version defined at the start of this document](#Version).

##### Client behaviour

- After getting a successful handshake response from the server, a client
**SHOULD** look for `hrpc-version=number` in the [`Sec-WebSocket-Extensions`][websocket_extensions_header] header.
The `number` will be the version of hRPC implemented by the server. A client
then **SHOULD** check if it can work with this version of the specification.
If a client can't work with this version, it **SHOULD** close the [`WebSocket`][websocket].
- [`WebSocket` binary messages][websocket_messages] **MUST** be used to send
RPCs request message in serialized form to the server.

##### Server behaviour

- [`WebSocket` binary messages][websocket_messages] **MUST** be used to send
RPCs response message or a [hRPC error] prefixed with an opcode.
    - If sending RPCs response message, the serialized message **MUST** be
    prefixed with `0`.
    - If sending a [hRPC error], the serialized error **MUST** be prefixed with `1`.

### Errors

When a request fails for whatever reason, an error response should be sent.
This applies to socket handshake requests and unary requests.

Below should be implemented for error responses:

- **MUST** set the body to a serialized [hRPC error] message.
- **MUST** set the status to an **unsuccessful** [HTTP status]. This should be
the status code corresponding to the error identifier. Implementations **MUST**
use the [hRPC errors document] to decide what status to set.
- **MUST** set the [`Content-Type`][http headers] header to `application/hrpc`.
- **MUST** add `<current spec version>` to `Hrpc-Version` header, where
`<current spec version>` is the [version defined at the start of this document](#Version).

[http headers]: https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html
[websocket]: https://datatracker.ietf.org/doc/html/rfc6455
[websocket_messages]: https://datatracker.ietf.org/doc/html/rfc6455#section-5.6
[websocket_protocol_header]: https://datatracker.ietf.org/doc/html/rfc6455#section-11.3.4
[websocket_extensions_header]: https://datatracker.ietf.org/doc/html/rfc6455#section-11.3.2
[hrpc error]: https://github.com/harmony-development/hrpc/blob/8e648895ece3eb1466f457125556cc86feeb92b3/protocol/hrpc.proto#L5-L13
[http status]: https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html
[hrpc errors document]: ./ERRORS.md
