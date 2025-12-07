package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// RateLimitMiddleware limits the number of requests per IP.
// Example: limit = 100, interval = 1 minute.
func RateLimitMiddleware(redisClient *redis.Client, logger *zap.Logger, limit int, interval time.Duration) gin.HandlerFunc {

	return func(c *gin.Context) {

		ip := c.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", ip)

		ctx := c.Request.Context()

		// Increment the request counter
		val, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			// If Redis is down → Do not Block traffic
			logger.Warn("rate-limit redis error", zap.Error(err))
			c.Next()
			return
		}

		// First request → set TTL
		if val == 1 {
			redisClient.Expire(ctx, key, interval)
		}

		// If request count exceeds limit => block
		if int(val) > limit {
			logger.Warn("rate limit exceeded",
				zap.String("ip", ip),
				zap.Int64("count", val),
				zap.Int("limit", limit),
			)

			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeded",
			})
			c.Abort()
			return
		}

		// Allow request
		c.Next()
	}
}
