# Installation
Download dependencies with `go mod tidy && go mod vendor`.

Build with `go build main.go` or run `go run main.go` from the root of the repository.

The server will serve on `0.0.0.0:50055` by default.

Files `tls.key`, `tls.crt` have to be in the same directory as the compiled executable in order to work.

Test the server with:
``` Shell
grpcurl -d '{"name":"krixlion"}' -cacert ./ca.crt localhost:50055 helloworld.Greeter/SayHello
```
This should return:
```json
{
  "message": "Hello krixlion!"
}
```


# Bug
## Expected behavior
A gRPC call goes through and either succeeds or a descriptive error is returned.

## Actual behavior
Error `14 No connection established` is returned with no further messages or logs.

## Steps to reproduce
1. Set up basic gRPC server with server-side TLS configured using [my repo](https://github.com/krixlion/insomnia_bug).
2. Create a new collection and a gRPC request.
3. Prefix the url with `grpcs://` e.g `grpcs://localhost:50055`.
4. Add the CA cert in the collection settings.
5. Send the request.