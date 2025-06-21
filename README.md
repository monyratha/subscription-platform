# Subscription Platform Skeleton

This repository contains a minimal skeleton implementation for a multi-tenant subscription platform written in Go. It demonstrates how the Auth service exposes both gRPC and HTTP endpoints.

## Services

- **Auth Service** (`services/auth-service`)
  - gRPC server on port `50051`
  - HTTP server on port `8081`

## Usage

Generate protobuf code:

```bash
protoc --go_out=. --go-grpc_out=. proto/auth.proto
```

Run the Auth service:

```bash
go run ./services/auth-service
```

This is only a starting point—you can extend the service implementations and add other services following the same pattern.

