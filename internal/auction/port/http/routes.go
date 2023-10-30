package http

import (
	"auction/internal/auction/repository"
	"auction/internal/auction/service"
	"auction/pkg/database"
	"auction/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, db database.IDatabase) {
	auctionRepo := repository.NewAuctionRepository(db)
	auctionService := service.NewAuctionService(auctionRepo)
	auctionHandler := NewAuctionHandler(auctionService)

	_ = middleware.JWTAuth()
	route := r.Group("/auctions")
	{
		route.GET("/", auctionHandler.GetAuctions)
		route.GET("/:id", auctionHandler.GetAuctionByID)
	}
}
