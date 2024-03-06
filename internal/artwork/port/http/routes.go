package http

import (
	"auction/internal/artwork/repository"
	"auction/internal/artwork/service"
	"auction/pkg/database"
	"auction/pkg/middleware"
	"auction/pkg/redis"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, db database.IDatabase, cache redis.IRedis) {
	artworkRepo := repository.NewArtworkRepository(db)
	artworkService := service.NewArtworkService(artworkRepo)
	artworkHandler := NewArtworkHandler(artworkService, cache)

	_ = middleware.JWTAuth(db)
	adminMiddleware := middleware.AdminAuth()
	authMiddleware := middleware.JWTAuth(db)
	route := r.Group("/artworks")
	{
		route.POST("/", authMiddleware, adminMiddleware, artworkHandler.CreateArtwork)
		route.PUT("/:id", authMiddleware, adminMiddleware, artworkHandler.UpdateAuction)
	}
}
