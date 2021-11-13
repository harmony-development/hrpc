# hRPC Errors

This document describes hRPC errors and lists error identifiers.

## hRPC Error Identifier List

This is a list of identifiers that are used by hRPC servers in the `Error` type.
The `hrpc.` prefix is reserved by hRPC and should not be used by server
implementations outside of the error identifiers listed here.

| Identifier                   | Description                                                     |
|------------------------------|-----------------------------------------------------------------|
| `hrpc.internal-server-error` | An error occured in the server.                                 |
| `hrpc.resource-exhausted`    | Reached resource quota or rate limited by the server.           |
| `hrpc.not-implemented`       | Endpoint is not implemented by the server.                      |
| `hrpc.not-found`             | Specified endpoint was not found on the server.                 |
| `hrpc.unavailable`           | The service couldn't be reached, used when the service is down. |

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

## Transport Specific Information

### HTTP

#### Additional Identifiers

The `hrpc.http.` prefix is reserved for hRPC HTTP errors.

| Identifier                        | Description                                                     |
|-----------------------------------|-----------------------------------------------------------------|
| `hrpc.http.bad-unary-request`     | The unary request didn't match spec requirements.               |
| `hrpc.http.bad-streaming-request` | The streaming handshake request didn't match spec requirements. |

#### Status Codes

If using hRPC over HTTP, the following response statuses should be used for
the respective identifiers:

| Identifier                               | Status                      |
|------------------------------------------|-----------------------------|
| `hrpc.internal-server-error`             | `500 Internal Server Error` |
| `hrpc.resource-exhausted` (rate limited) | `429 Too Many Requests`     |
| `hrpc.not-found`                         | `404 Not Found`             |
| `hrpc.not-implemented`                   | `501 Not Implemented`       |
| `hrpc.unavailable`                       | `503 Service Unavailable`   |
| `hrpc.http.bad-unary-request`            | `400 Bad Request`           |
| `hrpc.http.bad-streaming-request`        | `400 Bad Request`           |

For any non `hrpc.` identifiers, check the documentation of the hRPC API you
are using.