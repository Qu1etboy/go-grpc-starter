# Go gRPC Starter Template

## Prerequisite

install protoc

```sh
brew install protobuf
```

install go protobuf plugin

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## Generate Proto

generate proto into go code

```sh
protoc --go_out=. --go-grpc_out=. proto/user.proto
```

or using makefile

```sh
make generate
```
