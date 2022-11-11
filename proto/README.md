# Protobuf

This folder contains all [protobuf](https://developers.google.com/protocol-buffers) definitions.

## Usage

Import go clients:

```Go
import (
	pb "github.com/streadydevs/stready/proto/gen/go/<servicename>"
)

// Creating a gRPC stub client
var opts []grpc.DialOption
...
conn, err := grpc.Dial(*serverAddr, opts...)
if err != nil {
  ...
}
defer conn.Close()

client := pb.New<ServiceName>Client(conn)


// Call an rpc
thing, err := client.GetThing(context.Background(), &pb.Thing{"large", "red"})
if err != nil {
  ...
}
```

See [this guide](https://grpc.io/docs/languages/go/basics/) or Google around to learn more

## Build


### Installation

Make sure that you have buf installed first.

```
$ brew install bufbuild/buf/buf
```

You may need to create a temporary go.mod file in order to install the dependencies from `make install`.
```
$ go mod init example
```

Install the required tools to build gRPC.
```
$ brew install gcc@5
$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ make install
```

If you created the go mod files, make sure to remove them afterwards.
```
$ rm go.mod go.sum
```

Make sure `$ go env GOPATH` is in your $PATH.
(Add `export PATH=$PATH:$(go env GOPATH)/bin` to your .bashrc or .zshrc)


### Generating protobuf/gRPC stubs

Simply run `$ make`.

This runs `$ buf generate` which builds everything based on the contents of `buf.gen.yaml`.

### Linting

`$ buf` supports linting and checking for breaking changes. To lint and check for breaking changes run:
```
$ make lint
```
