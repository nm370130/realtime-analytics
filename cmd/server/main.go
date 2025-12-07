package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	httpRoutes "github.com/nm370130/realtime-analytics/internal/http"
)

func main() {

	//LOGGER INIT
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("starting real-time analytics service")

	// LOAD ENV
	mysqlDSN := os.Getenv("MYSQL_DSN")   // Example: user:pass@tcp(localhost:3306)/analytics?parseTime=true
	redisAddr := os.Getenv("REDIS_ADDR") // Example: localhost:6379

	if mysqlDSN == "" || redisAddr == "" {
		log.Fatal("Missing MYSQL_DSN or REDIS_ADDR environment variables")
	}

	//MYSQL INIT
	db, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
	if err != nil {
		logger.Fatal("failed to connect to MySQL", zap.Error(err))
	}
	logger.Info("connected to MySQL")

	// REDIS INIT
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		logger.Fatal("failed to connect to Redis", zap.Error(err))
	}
	logger.Info("connected to Redis")

	//ROUTER INIT
	deps := httpRoutes.Dependencies{
		MySQL:  db,
		Redis:  redisClient,
		Logger: logger,
	}

	router := httpRoutes.NewRouter(deps)

	// HTTP SERVER
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 12 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run server in goroutine
	go func() {
		logger.Info("HTTP server listening", zap.String("port", "8080"))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("server error", zap.Error(err))
		}
	}()

	//GRACEFUL SHUTDOWN
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	logger.Warn("shutdown signal received")

	// Wait max 10 seconds to finish ongoing requests
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("server forced shutdown", zap.Error(err))
	}

	logger.Info("server exited gracefully")
}
