# Multi-Tenant Subscription Platform

This repository contains a minimal implementation of a SaaS platform where users can create their own web sites and subscribe to flexible plans. The project follows the [go-clean-template](https://github.com/evrone/go-clean-template) architecture and demonstrates how each service exposes both gRPC and HTTP endpoints.

The current codebase only implements the **Auth Service** as a proof of concept. Other services share the same layout and will be added later.

## Objective

* Users can register and log in
* Each user manages their own sites
* Flexible subscription plans (Free, 3mo, 6mo, 1yr)
* Product and category management
* Tracking of site traffic and usage
* Access control via roles and policies

## Tech Stack

| Category              | Tool/Library                                                       |
| --------------------- | ------------------------------------------------------------------ |
| Language              | Go 1.21                                                            |
| Architecture Template | [go-clean-template](https://github.com/evrone/go-clean-template)   |
| Internal Comms        | gRPC (Protocol Buffers)                                            |
| Public API            | Gin (RESTful HTTP APIs)                                            |
| ORM                   | GORM                                                               |
| Authentication        | Authboss                                                           |
| Authorization         | Casbin (RBAC/ABAC)                                                 |
| Async Messaging       | RabbitMQ (AMQP)                                                    |
| Service Discovery     | Consul                                                             |
| Config Management     | Viper + `.env`                                                     |
| Database              | MySQL 8.x                                                          |
| API Gateway           | Traefik / custom Gin proxy                                         |
| Logging               | Loki + Promtail                                                    |
| Monitoring            | Prometheus + Grafana                                               |
| Containerization      | Docker + Docker Compose                                            |
| CI/CD                 | GitHub Actions / GitLab CI                                         |

## Core Business Services

| Service                | Description                                    | gRPC Port | HTTP Port |
| ---------------------- | ---------------------------------------------- | --------- | --------- |
| `auth-service`         | User login/registration, JWT, Authboss, Casbin | `50051`   | `8081`    |
| `site-service`         | Site creation, user roles, URLs                | `50052`   | `8082`    |
| `subscription-service` | Plan activation, expiry, billing events        | `50053`   | `8083`    |
| `catalog-service`      | Categories and product management              | `50054`   | `8084`    |
| `stat-service`         | Site analytics: visits, referrals, platforms   | `50055`   | `8085`    |

## Supporting Infrastructure Services

| Service              | Purpose                                             | Ports                                 |
| -------------------- | --------------------------------------------------- | ------------------------------------- |
| `gateway-service`    | Entry point: routes frontend APIs, JWT verification | `80`, `443`, `8080`                   |
| `discovery-service`  | Service registry via Consul                         | `8500`                                |
| `broker-service`     | RabbitMQ for async events                           | `5672` (AMQP), `15672` (UI)           |
| `logging-service`    | Centralized logging with Loki                       | `3100`                                |
| `monitoring-service` | Metrics collection + dashboards                     | `9090` (Prometheus), `3000` (Grafana) |

## gRPC and Protobuf

Protobuf definitions live in the `proto/` directory. Generate stubs with:

```bash
protoc --go_out=. --go-grpc_out=. proto/<service>.proto
```

For example, to generate the Auth service stubs:

```bash
protoc --go_out=. --go-grpc_out=. proto/auth.proto
```

## Messaging Events (RabbitMQ)

| Event                    | Publisher              | Subscribers                |
| ------------------------ | ---------------------- | -------------------------- |
| `user.registered`        | `auth-service`         | email, analytics           |
| `subscription.activated` | `subscription-service` | stat-service, notification |
| `site.viewed`            | `stat-service`         | analytics, marketing       |

## Folder Structure

```
subscription-platform/
├── proto/
│   └── *.proto
├── services/
│   ├── auth-service/
│   ├── site-service/
│   ├── subscription-service/
│   ├── catalog-service/
│   └── stat-service/
├── supporting/
│   ├── gateway-service/
│   ├── discovery-service/
│   ├── broker-service/
│   ├── logging-service/
│   └── monitoring-service/
├── docker-compose.yaml
└── README.md
```

## Running the Auth Service

Build and run the example Auth service from the repository root:

```bash
# build dependencies
go build ./...

# start the gRPC and HTTP servers
go run ./services/auth-service
```

The service listens on gRPC port `50051` and HTTP port `8081`.

---

This repository serves as a starting point. Each additional service should mirror the structure of `auth-service` and register with Consul for discovery. Docker Compose can be used for local development to spin up all services, the gateway, and supporting infrastructure.
