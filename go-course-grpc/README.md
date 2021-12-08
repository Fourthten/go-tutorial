## Simple GRPC
```bash
go mod init go-course-grpc (*parent folder)

https://github.com/grpc/grpc-go
go get -u google.golang.org/grpc

https://grpc.io/docs/protoc-installation/
brew install protobuf
protoc --version

https://github.com/ktr0731/evans
brew tap ktr0731/evans
brew install evans
evans -p 50051 -r

https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/
brew services start mongodb-community@5.0

https://github.com/mongodb/mongo-go-driver#installation
go get go.mongodb.org/mongo-driver/mongo

https://gopkg.in/mgo.v2/bson
go get gopkg.in/mgo.v2/bson

https://github.com/simplesteph/grpc-go-course

go run blog/blog_server/server.go
go run blog/blog_client/client.go

go run calculator/calculator_server/server.go
go run calculator/calculator_client/client.go

go run greet/greet_server/server.go
go run greet/greet_client/client.go
```

## MongoDB
```bash
To run as a macOS service
    brew services start mongodb-community@5.0
To run manually as a background process
    Intel : mongod --config /usr/local/etc/mongod.conf --fork
    Apple M1 : mongod --config /opt/homebrew/etc/mongod.conf --fork
```

* Unary, Server Streaming, Client Streaming, BiDi Streaming
* Error Handling, Deadlines, SSL Encryption
* Greeting, Calculator Service
* Blog API CRUD w/ MongoDB

[gRPC [Golang] Master Class: Build Modern API & Microservices](https://www.udemy.com/share/101Zo03@hLg16RK2Dy4Bqev4WGrRQNN4e06juzLsi2hHw-T0girQ21D6EbTpVx43gDxdo9xZ/)