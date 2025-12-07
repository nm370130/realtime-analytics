package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv string

	HTTPPort        string
	ShutdownTimeout time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration

	MySQLDSN string

	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

func Load() *Config {
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()

	// defaults
	viper.SetDefault("ENV", "local")
	viper.SetDefault("HTTP_PORT", "8080")
	viper.SetDefault("SHUTDOWN_TIMEOUT", "15s")
	viper.SetDefault("READ_TIMEOUT", "10s")
	viper.SetDefault("WRITE_TIMEOUT", "10s")
	viper.SetDefault("IDLE_TIMEOUT", "60s")


	cfg := &Config{
		AppEnv:        viper.GetString("ENV"),
		HTTPPort:      viper.GetString("HTTP_PORT"),
		MySQLDSN:      viper.GetString("MYSQL_DSN"),
		RedisAddr:     viper.GetString("REDIS_ADDR"),
		RedisPassword: viper.GetString("REDIS_PASSWORD"),
		RedisDB:       viper.GetInt("REDIS_DB"),
	}

	// Parse durations
	var err error
	cfg.ShutdownTimeout, err = time.ParseDuration(viper.GetString("SHUTDOWN_TIMEOUT"))
	if err != nil {
		log.Fatalf("invalid SHUTDOWN_TIMEOUT: %v", err)
	}
	cfg.ReadTimeout, err = time.ParseDuration(viper.GetString("READ_TIMEOUT"))
	if err != nil {
		log.Fatalf("invalid READ_TIMEOUT: %v", err)
	}
	cfg.WriteTimeout, err = time.ParseDuration(viper.GetString("WRITE_TIMEOUT"))
	if err != nil {
		log.Fatalf("invalid WRITE_TIMEOUT: %v", err)
	}
	cfg.IdleTimeout, err = time.ParseDuration(viper.GetString("IDLE_TIMEOUT"))
	if err != nil {
		log.Fatalf("invalid IDLE_TIMEOUT: %v", err)
	}

	if cfg.MySQLDSN == "" {
		log.Fatal("MYSQL_DSN is required (e.g. user:pass@tcp(localhost:3306)/rt_analytics?parseTime=true&loc=Local)")
	}
	if cfg.RedisAddr == "" {
		log.Fatal("REDIS_ADDR is required (e.g. localhost:6379)")
	}

	return cfg
}
