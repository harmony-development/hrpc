# hRPC Error Identifier List

This is a list of identifiers that are used by hRPC servers in the `Error` type.
The `hrpc.` prefix is reserved by hRPC and should not be used by server
implementations outside of the error identifiers listed here.

- `hrpc.internal-server-error`: An error occured in the server.
- `hrpc.resource-exhausted`: Reached resource quota or rate limited by the
server.
- `hrpc.not-implemented`: Endpoint is not implemented by the server.
- `hrpc.not-found`: Specified endpoint was not found on the server.
- `hrpc.unavailable`: The server couldn't be reached, most likely means the
server is down.

## Error Retrying

Server implementations should use the `RetryInfo` protobuf message and put
it in the `details` field of `Error`. Client implementations then can look
at the `details` field for `RetryInfo` and use the information contained
there.

All hRPC servers should put a valid `RetryInfo` in the `details` field of
`Error` for the error identifiers listed below.

- `hrpc.unavailable` errors may be retried using exponential backoff. Minimum
retry delay and repetition count depends on the server, but should be 1 second
and 1 repetition if not documented.
- `hrpc.resource-exhausted` errors may be retried depending on the server.