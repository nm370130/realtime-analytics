package http

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/nm370130/realtime-analytics/internal/metrics"
	"github.com/nm370130/realtime-analytics/internal/modules"
	"github.com/nm370130/realtime-analytics/internal/sensors"

	"github.com/nm370130/realtime-analytics/internal/common"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Dependencies struct {
	MySQL  *gorm.DB
	Redis  *redis.Client
	Logger *zap.Logger
}

// ROUTER SETUP 

func NewRouter(deps Dependencies) *gin.Engine {
	r := gin.New()

	// Middlrewares
	r.Use(gin.Recovery())

	// Request ID middleware
	r.Use(RequestIDMiddleware())

	// Logging Middleware (now counts 4xx API rejections)
	r.Use(LoggingMiddleware(deps.Logger, deps.Redis))

	// Rate limiting (100 req/min per IP)
	r.Use(RateLimitMiddleware(deps.Redis, deps.Logger, 100, time.Minute))

	// CACHE
	cache := common.NewCache(deps.Redis)

	// METRICS MODULE
	metricsRepo := metrics.NewRepository(deps.MySQL, deps.Redis)
	metricsSvc := metrics.NewService(metricsRepo, cache)
	metricsHandler := metrics.NewHandler(metricsSvc, deps.Logger)

	// SENSORS MODULE
	sensorsRepo := sensors.NewRepository(deps.MySQL)
	sensorsSvc := sensors.NewService(sensorsRepo, cache)
	sensorsHandler := sensors.NewHandler(sensorsSvc, deps.Logger)

	// MODULES MODULE
	modulesRepo := modules.NewRepository(deps.MySQL)
	modulesSvc := modules.NewService(modulesRepo)
	modulesHandler := modules.NewHandler(modulesSvc, deps.Logger)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")

	{
		// Metrics APIs
		v1.GET("/metrics/summary", metricsHandler.GetSummary)
		v1.GET("/metrics/history", metricsHandler.GetHistory)

		// Sensors APIs
		v1.GET("/sensors/type-breakdown", sensorsHandler.GetTypeBreakdown)

		// Modules APIs
		v1.GET("/modules", modulesHandler.GetModules)
	}

	return r
}
