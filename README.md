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

1. Run calculator server:

        $ go run calculator/calculator_server/server.go

1. Run calculator client:

        $ go run calculator/calculator_client/client.go


(Update) How to run `Sum` RPC via REST:

1. Download the following packages:

        $ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
        $ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

1. Compile `.proto` file to `.go` file:

        $ protoc -I/usr/local/include -I. \
            -I$GOPATH/src \
            -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
            --go_out=plugins=grpc:. \
            calculator/calculatorpb/calculator.proto

1. Generate reverse-proxy using `protoc-gen-grpc-gateway`:

        $ protoc -I/usr/local/include -I. \
            -I$GOPATH/src \
            -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
            --grpc-gateway_out=logtostderr=true:. \
            calculator/calculatorpb/calculator.proto
    
    It will generate `calculator/calculatorpb/calculator.pb.gw.go` file

1. (Optional) Generate swagger definitions using `protoc-gen-swagger`:

        $ protoc -I/usr/local/include -I. \
            -I$GOPATH/src \
            -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
            --swagger_out=logtostderr=true:. \
            calculator/calculatorpb/calculator.proto

    It will generate `calculator/calculatorpb/calculator.swagger.json` file

1. Run calculator server:

        $ go run calculator/calculator_server/server.go

1. Run calculator gateway:

        $ go run calculator/calculator_server/gateway.go

1. Send message via `curl`:

        $ curl -X POST -k http://127.0.0.1:8081/sum -d '{"first_number": 5, "second_number": 3}'

    And you will get response `{"sum_result":8}` as result

### SSL certificate

1. Install openssl:

        $ sudo apt-get install openssl

2. Create needed certificates, keys:

        $ cd ssl
        $ chmod +x ./generate_files.sh
        $ ./generate_files.sh

    if you got `Cannot open file .rnd` error - create an empty `.rnd` file

### Blog example

1. Install MongoDB:

    * Download and install MongoDB from [here](https://www.mongodb.com/download-center/community)

    * Create folder to store data:

            $ mkdir ~/mongodata
            $ mkdir ~/mongodata/db

    * Run MongoDB:

            $ /usr/bin/mongod --dbpath ~/mongodata/db

1. Install UI for your MongoDB:

    * Download Robo 3T from [here](https://robomongo.org/download)

    * Switch to download directory and run these commands:

            $ tar -xvzf robo3t*.tar.gz
            $ sudo mkdir /usr/local/bin/robomongo
            $ sudo mv robo3t*/* /usr/local/bin/robomongo
            $ cd /usr/local/bin/robomongo/bin
            $ sudo chmod +x robo3t

    * Open .bashrc file:

            $ sudo vim ~/.bashrc

      And add the following line to the end of the file:

            alias robomongo='/usr/local/bin/robomongo/bin/robo3t'

    * Reload it using the following command:

            $ source ~/.bashrc

    * Run robomongo from your terminal:

            $ robomongo

1. Install MongoDB driver for Golang:

        $ go get go.mongodb.org/mongo-driver/mongo

##### Run Blog Example

1. Compile `.proto` file to `.go` file:

        $ protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.

1. Run blog server:

        $ go run blog/blog_server/server.go

1. Run blog client:

        $ go run blog/blog_client/client.go