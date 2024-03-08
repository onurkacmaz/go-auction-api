package config

import (
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Schema struct {
	Environment   string `env:"environment"`
	HttpHost      string `env:"http_host"`
	HttpPort      int    `env:"http_port"`
	DatabaseURI   string `env:"database_uri"`
	RedisURI      string `env:"redis_uri"`
	RedisPassword string `env:"redis_password"`
	RedisDB       int    `env:"redis_db"`
	AuthSecret    string `env:"auth_secret"`
}

const (
	ProductionEnv       = "production"
	DatabaseTimeout     = 5 * time.Second
	AuctionsCachingTime = 1 * time.Minute
	AuctionCachingTime  = 1 * time.Minute
	ArtworkCachingTime  = 1 * time.Minute
)

var (
	cfg Schema
)

func LoadConfig() *Schema {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	err := godotenv.Load(filepath.Join(currentDir, "config.yaml"))
	if err != nil {
		log.Printf("Error on load configuration file, error: %v", err)
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error on parsing configuration file, error: %v", err)
	}

	return &cfg
}

func GetConfig() *Schema {
	return &cfg
}
