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
