package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const RequestIDKey = "request_id"

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Accept from client if provided
		rid := c.GetHeader("X-Request-ID")
		if rid == "" {
			rid = uuid.NewString()
		}

		// Store in context using standard key "request_id"
		c.Set(RequestIDKey, rid)

		// Add to response headers
		c.Writer.Header().Set("X-Request-ID", rid)

		c.Next()
	}
}

func LoggingMiddleware(logger *zap.Logger, redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		// Process request
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		// 4xx: API REJECTED COUNTER
		if status >= 400 && status < 500 {
			ctx := c.Request.Context()
			redisClient.Incr(ctx, "api_rejected:5min")
			redisClient.Expire(ctx, "api_rejected:5min", 5*time.Minute)
		}

		// LOGGING
		logger.Info("incoming request",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.String("client_ip", clientIP),
			zap.Duration("latency", latency),
			zap.String("request_id", c.GetString("request_id")),
		)
	}
}
