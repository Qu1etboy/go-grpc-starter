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

## DB Migration

This template used Goose as a db migration tool

Create new migration

```sh
goose -dir migrations/ create <migration_name> sql
```

Run migration

```sh
goose -dir migrations/ postgres "host=localhost user=test password=test dbname=go-grpc-starter port=5432 sslmode=disable timezone=utc" up
```

## k8s

Create configmap auto from .env file

```sh
kubectl create configmap go-grpc-starter-config --from-env-file=.env
```

Apply k8s config

```sh
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

And port forward to access via localhost

```sh
kubectl port-forward svc/go-grpc-starter-service 8980:8980
```
