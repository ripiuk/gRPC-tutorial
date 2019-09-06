## gRPC [Golang] tutorial

This is an implementation of [Udemy tutorial](https://www.udemy.com/grpc-golang/)

1. install [GO](https://golang.org/):

        $ wget https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz
        $ sudo tar -C /usr/local -xzf go1.12.9.linux-amd64.tar.gz
        $ export PATH=$PATH:/usr/local/go/bin

1. Install [grpc-go](https://github.com/grpc/grpc-go):

        $ go get -u google.golang.org/grpc

1. Install [Protocol Buffers for GO](https://github.com/golang/protobuf):

        $ go get -u github.com/golang/protobuf/protoc-gen-go

1. Install [Protocol Buffers (protoc)](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)

        $ curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip
        $ unzip protoc-3.9.1-linux-x86_64.zip -d protoc3
        $ sudo mv protoc3/bin/* /usr/local/bin/ && sudo mv protoc3/include/* /usr/local/include/
        $ rm protoc-3.9.1-linux-x86_64.zip && rm -rf protoc3

### Greeting example

1. Compile `.proto` file to `.go` file:

        $ protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

1. Run greet server:

        $ go run greet/greet_server/server.go

1. Run greet client:

        $ go run greet/greet_client/client.go

### Calculator example

1. Compile `.proto` file to `.go` file:

        $ protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.

1. Run greet server:

        $ go run calculator/calculator_server/server.go

1. Run greet client:

        $ go run calculator/calculator_client/client.go