Real-Time Analytics Platform (Golang + MySQL + Redis)

Clean Architecture (Handler → Service → Repository)
Caching + Rate Limiting
Structured Logging

Features
1. Real-Time Summary Metrics API
GET /api/v1/metrics/summary
Returns:
active users (from Redis)
active users by platform
API rejections last 5 minutes
new projects created in last 7 days
total live projects
sensors online/offline

2. Metrics History API
GET /api/v1/metrics/history?type=activeUsers&interval=5m
Returns time-series graph data for:
active users
API rejections
Any future metric

3. Sensor Type Distribution API
GET /api/v1/sensors/type-breakdown
Returns aggregated counts of sensors grouped by sensor type.
Includes Redis caching.

4. Modules Metadata API
GET /api/v1/modules
Returns module name, current version, upcoming version, and release info.

# Additional Features
Redis caching layer
Global rate-limiting middleware (100 req/min per IP)
Request-ID middleware
Structured logging via Zap
Config system using Viper
Graceful shutdown

# Clean folder structure
- Project Structure
realtime-analytics/
├── cmd/
│   └── server/main.go
├── internal/
│   ├── config/
│   ├── db/
│   ├── http/
│   ├── metrics/
│   ├── sensors/
│   ├── modules/
│   ├── models/
│   ├── common/
│   └── logger/
├── migrations/
│   └── 001_init_schema.sql
├── go.mod
└── README.md

# Requirements
Go 1.20+
MySQL 8+
Redis 6+
Git

# Environment Variables
Create a .env file and copy from .example.env file

# Local Deployment
git clone https://github.com/nm370130/realtime-analytics.git
cd realtime-analytics

Make sure Docker & Docker Compose are installed.

Start all services:

docker-compose up --build

Build your Go application

Start MySQL with seeded schema + mock data

Start Redis

Expose the Go API at http://localhost:8080


# Stopping the server
docker-compose down

# To delete MySQL + Redis data too:
docker-compose down -v



Author
Nitesh Mishra
Backend Engineer — Golang