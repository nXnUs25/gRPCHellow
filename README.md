# gRPCHellow
Simple Greetings gRPC application to demonstrate Unary server / client 


## Biuldning all
```shell
❯ make all
#protoc -Iproto --go_opt=module=proto --go_out=. --go-grpc_opt=module=proto --go-grpc_out=. proto/*.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
go build -o bin/server ./server
go build -o bin/client ./client
```

## Running

### Server

```shell
❯ bin/server
2022/05/24 08:34:48 listening on 0.0.0.0:50051
2022/05/24 08:35:10 Greetings func called fname:"Auggie"
```

### Client

```shell
❯ bin/client
2022/05/24 08:35:10 Greetings from Hellow grpc app Auggie
```

## Help 

```shell
❯ make
all                            Generate Pbs and build
clean_pb_hellow                Clean generated files for greet
rebuild                        Rebuild the whole project
about                          Display info related to the build
help                           Show this help
```

## Clean all

```shell
❯ make clean
rm -fv proto/*.pb.go
proto/hello.pb.go
proto/hello_grpc.pb.go
rm -vrf bin
bin/server
bin/client
bin
```
