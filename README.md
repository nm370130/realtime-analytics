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

Additional Features
Redis caching layer
Global rate-limiting middleware (100 req/min per IP)
Request-ID middleware
Structured logging via Zap
Config system using Viper
Graceful shutdown

Clean folder structure
Project Structure
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

Requirements
Go 1.20+
MySQL 8+
Redis 6+
Git

Environment Variables
Create a .env file and copy from .example.env file

Or with Docker:
docker exec -i mysql mysql -u root -p realtime < migrations/001_init_schema.sql

Redis Keys Used
Key	                        Description
active_users	            total active users
active_users:web	        web active users
active_users:mobile-app1	mobile app1 users
active_users:mobile-app2	mobile app2 users
api_rejected:5min	        rejected API count (expires 5 min)
sensors:type-breakdown	    cached sensor type distribution

Running the Application
1. Install dependencies
go mod tidy

2. Start MySQL & Redis

Using Docker:

docker run --name mysql8 -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:8
docker run --name redis -p 6379:6379 -d redis

3. Run the server
cd cmd/server
go run main.go

Server runs at:
http://localhost:8080

API Endpoints
Health
GET /health

Summary Metrics
GET /api/v1/metrics/summary

Metrics History
GET /api/v1/metrics/history?type=activeUsers&interval=5m

Sensor Type Breakdown
GET /api/v1/sensors/type-breakdown

Modules Metadata
GET /api/v1/modules




Author
Nitesh Mishra
Backend Engineer — Golang