package http

import (
	auctionHttp "auction/internal/auction/port/http"
	bidHttp "auction/internal/bid/port/http"
	userHttp "auction/internal/user/port/http"
	"auction/pkg/config"
	database "auction/pkg/database"
	"auction/pkg/redis"
	"auction/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Server struct {
	engine *gin.Engine
	cfg    *config.Schema
	db     database.IDatabase
	cache  redis.IRedis
}

func NewServer(cfg *config.Schema, db database.IDatabase, cache redis.IRedis) *Server {
	if cfg.Environment == config.ProductionEnv {
		gin.SetMode(gin.ReleaseMode)
	}

	return &Server{
		engine: gin.Default(),
		cfg:    cfg,
		db:     db,
		cache:  cache,
	}
}

func (s *Server) Run() error {
	_ = s.engine.SetTrustedProxies(nil)

	if err := s.MapRoutes(); err != nil {
		log.Fatalf("MapRoutes Error: %v", err)
	}

	s.engine.GET("/health", func(c *gin.Context) {
		response.JSON(c, http.StatusOK, nil)
		return
	})

	log.Println("HTTP server is listening on PORT: ", s.cfg.HttpPort)
	if err := s.engine.Run(fmt.Sprintf("%s:%d", s.cfg.HttpHost, s.cfg.HttpPort)); err != nil {
		log.Fatalf("Running HTTP server: %v", err)
	}

	return nil
}

func (s *Server) GetEngine() *gin.Engine {
	return s.engine
}

func (s *Server) MapRoutes() error {
	v1 := s.engine.Group("/api/v1")
	userHttp.Routes(v1, s.db)
	auctionHttp.Routes(v1, s.db)
	bidHttp.Routes(v1, s.db)
	return nil
}
