# 🚀 Project Overview – Multi-Tenant Subscription Platform

### *(Go + gRPC + Gin + RabbitMQ + MySQL + Authboss + Casbin + go-clean-template)*


## 🎯 Objective

Build a scalable SaaS platform where users can:

* Register/login
* Create and manage their own websites
* Subscribe to flexible plans (Free, 3mo, 6mo, 1yr)
* Manage products and categories
* Track traffic and usage stats
* Enforce access control via roles and policies

---

## ⚙️ Tech Stack

| Category              | Tool/Library                                                       |
| --------------------- | ------------------------------------------------------------------ |
| Language              | Go 1.24.3                                                           |
| Architecture Template | [`go-clean-template`](https://github.com/evrone/go-clean-template) |
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

---

## 🧱 Core Business Services

Each service:

* Uses `go-clean-template` architecture
* Runs a **gRPC server** and a **Gin HTTP server**
* Has its own database tables

| Service                | Description                                    | gRPC Port | HTTP Port |
| ---------------------- | ---------------------------------------------- | --------- | --------- |
| `auth-service`         | User login/registration, JWT, Authboss, Casbin | `50051`   | `8081`    |
| `site-service`         | Site creation, user roles, URLs                | `50052`   | `8082`    |
| `subscription-service` | Plan activation, expiry, billing events        | `50053`   | `8083`    |
| `catalog-service`      | Categories and product management              | `50054`   | `8084`    |
| `stat-service`         | Site analytics: visits, referrals, platforms   | `50055`   | `8085`    |

---

## 🛠 Supporting Infrastructure Services

| Service              | Purpose                                             | Ports                                 |
| -------------------- | --------------------------------------------------- | ------------------------------------- |
| `gateway-service`    | Entry point: routes frontend APIs, JWT verification | `80`, `443`, `8080`                   |
| `discovery-service`  | Service registry via Consul                         | `8500`                                |
| `broker-service`     | RabbitMQ for async events                           | `5672` (AMQP), `15672` (UI)           |
| `logging-service`    | Centralized logging with Loki                       | `3100`                                |
| `monitoring-service` | Metrics collection + dashboards                     | `9090` (Prometheus), `3000` (Grafana) |

---

## 🔌 Communication Design

### Internal Communication

* ✅ All services use **gRPC** for performance and strict contracts
* ✅ Services discover each other via **Consul**

### External Communication

* ✅ Exposed via **Gin HTTP APIs**
* ✅ All traffic enters via **gateway-service**

---

## 📄 gRPC & Protobuf

* Each service defines its API in `/proto/*.proto`
* Code generated with:

```bash
protoc --go_out=. --go-grpc_out=. proto/your_service.proto
```

Example:

```proto
service AuthService {
  rpc Register(RegisterRequest) returns (RegisterResponse);
}
```

---

## 🐇 Messaging Events (RabbitMQ)

| Event                    | Publisher              | Subscribers                |
| ------------------------ | ---------------------- | -------------------------- |
| `user.registered`        | `auth-service`         | email, analytics           |
| `subscription.activated` | `subscription-service` | stat-service, notification |
| `site.viewed`            | `stat-service`         | analytics, marketing       |

---

## 🌐 Frontend Access Points

| Client         | Description                          | Path Prefix | Auth      |
| -------------- | ------------------------------------ | ----------- | --------- |
| Admin Portal   | Admin-level control of platform      | `/admin/**` | Admin JWT |
| User Dashboard | Manage site, products, subscriptions | `/user/**`  | User JWT  |
| Public Website | Display site and product info        | `/site/**`  | Optional  |

---

## 📁 Folder Structure

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

---

## ✅ Load Balancing Strategy

* Each service runs on a fixed **internal port**
* Public traffic is routed via `gateway-service` (Traefik or Gin reverse proxy)
* Horizontal scaling is handled by:

  * Registering all replicas in **Consul**
  * Traefik (or a custom client) performs round-robin or weighted balancing
* gRPC clients can use service discovery to load balance gRPC calls

✅ Port design is **load-balancer friendly**
✅ No port conflicts due to containerized isolation

---

## 🚀 Development & Deployment Strategy

### 👨‍💻 Local Development

* Use monorepo folder layout
* Run services with Docker Compose
* Expose gRPC and Gin ports for each service
* Share proto definitions in `proto/`

### 🚀 Production

* Split into **polyrepos** for each service
* CI/CD pipelines per repo
* Deploy via Docker Swarm or Kubernetes
* Use service discovery + gateway in cloud setup

---

## ✅ Summary

| Item              | Status                 |
| ----------------- | ---------------------- |
| Core Architecture | ✅ Defined              |
| go-clean-template | ✅ Used                 |
| Messaging Layer   | ✅ RabbitMQ             |
| Database & ORM    | ✅ MySQL + GORM         |
| Authentication    | ✅ Authboss             |
| Authorization     | ✅ Casbin               |
| Service Discovery | ✅ Consul               |
| Monitoring & Logs | ✅ Prometheus + Loki    |
| Load Balancing    | ✅ via Gateway & Consul |
| API Design        | ✅ gRPC + Gin           |

---
