package http

import (
	"auction/internal/auction/repository"
	"auction/internal/auction/service"
	"auction/pkg/database"
	"auction/pkg/middleware"
	"auction/pkg/redis"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, db database.IDatabase, cache redis.IRedis) {
	auctionRepo := repository.NewAuctionRepository(db)
	auctionService := service.NewAuctionService(auctionRepo)
	auctionHandler := NewAuctionHandler(auctionService, cache)

	_ = middleware.JWTAuth(db)
	adminMiddleware := middleware.AdminAuth()
	authMiddleware := middleware.JWTAuth(db)
	route := r.Group("/auctions")
	{
		route.GET("", auctionHandler.GetAuctions)
		route.GET("/:id", auctionHandler.GetAuctionByID)
		route.POST("/", authMiddleware, adminMiddleware, auctionHandler.CreateAuction)
		route.PUT("/:id", authMiddleware, adminMiddleware, auctionHandler.UpdateAuction)
	}
}
