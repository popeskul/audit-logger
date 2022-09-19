# Audit Logger

Audit Logger is a simple gRPC interceptor that logs all requests and write it to the database.

## Setup
### 1. Build project
```sh
go build -o audit-logger cmd/main.go
```
### 2. Run built project
```sh
./audit-logger
```

### Generating gRPC code
```sh
protoc --go_out=. --go-grpc_out=. proto/audit.proto
```

#### Docker
```sh
docker run --rm -d --name audit-logo-mongo -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=123 -p 27017:27017 mongo:latest
docker run --rm -d --name audit-logo-mongo-test -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=123 -p 27011:27017 mongo:latest
```
