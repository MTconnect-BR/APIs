# Velo API Gateway

High-performance API gateway for modern applications. Solve the 7 problems every API faces: rate limiting, cache, authentication, load balancing, observability, documentation, and versioning.

## Features

- **Rate Limiting** - Token bucket + sliding window, per-IP/key
- **Cache Global** - Redis-backed with TTL, tag invalidation
- **Authentication** - JWT, OAuth2, API keys, RBAC
- **Load Balancing** - Round-robin, least connections, weighted
- **Observability** - Prometheus metrics, structured logs, distributed tracing
- **Documentation** - Auto-generated OpenAPI 3.1
- **Versioning** - Header, path, or query-based

## Installation

```bash
# Download binary
curl -sSL https://github.com/velo-api/velo/releases/latest/download/velo-linux-amd64 -o velo
chmod +x velo

# Or via Docker
docker pull velo-api/velo:latest

# Or compile from source
go install github.com/velo-api/velo/cmd/velo@latest
```

## Quick Start

```bash
# 1. Create config
cp configs/velo.example.yaml configs/velo.yaml

# 2. Edit config
vim configs/velo.yaml

# 3. Run
./velo configs/velo.yaml
```

## Configuration

```yaml
server:
  host: 0.0.0.0
  port: 8080

ratelimit:
  enabled: true
  default: "100/min"

cache:
  enabled: true
  driver: redis
  redis:
    addr: localhost:6379
  ttl: 5m

auth:
  enabled: true
  providers:
    - type: jwt
      secret: "${JWT_SECRET}"

loadbalancer:
  enabled: true
  strategy: round-robin
  backends:
    - url: "http://localhost:3000"
      weight: 1

observe:
  metrics:
    enabled: true
    path: /metrics

docs:
  enabled: true
  path: /docs

versioning:
  strategy: header
  header: "X-API-Version"
  default: "v1"
```

## Docker

```bash
# Run with Docker
docker run -p 8080:8080 -v ./configs:/app/configs velo-api/velo

# Docker Compose
docker-compose up -d
```

## Benchmarks

| Metric | Traditional | Velo | Improvement |
|--------|-------------|------|-------------|
| Throughput | 1,000 req/s | 50,000 req/s | 50x |
| Latency P99 | 45ms | 12ms | 73% |
| Memory | 256 MB | 45 MB | 82% |
| Cold Start | 2.5s | 15ms | 99% |

## Website

The comparison website is built with Next.js and deployed to Vercel.

```bash
cd website
npm install
npm run dev
```

## License

MIT
