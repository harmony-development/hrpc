# hRPC

hRPC is a simple means of accessing services defined in Protobuf.

## Routing

All routes are of the form `/{PackageName}.{ServiceName}/{RPC}%s`.

For example, `/protocol.chat.v1.ChatService/CreateGuild`.

All requests are POST requests unless otherwise specified.

## Headers

Set the MIME-Type to `application/hrpc` for requests.
This may be expanded to other formats in the future.

## Unary Requests (S -> S)

This is straightforward.
Request body is the serialised message of the input type.
Response body is the serialised message of the output type.

## Requests

Requests to a streaming route are always upgraded to WebSockets.

### Dual Streaming Requests (stream S -> stream S)

Clients send binary messages including serialised Protobuf messages.
Servers send binary messages including serialised Protobuf messages.

### Unary Streaming Requests (S -> stream S)

Clients send a single binary message over the socket.
Once the client has sent a message, the server begins sending binary messages including serialised Protobuf messages back.
