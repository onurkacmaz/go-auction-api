package http

import (
	artworkRepository "auction/internal/artwork/repository"
	auctionRepository "auction/internal/auction/repository"
	"auction/internal/bid/repository"
	"auction/internal/bid/service"
	"auction/pkg/database"
	"auction/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup, db database.IDatabase) {
	bidRepo := repository.NewBidRepository(db)
	artworkRepo := artworkRepository.NewArtworkRepository(db)
	auctionRepo := auctionRepository.NewAuctionRepository(db)
	bidService := service.NewBidService(bidRepo, artworkRepo, auctionRepo)
	bidHandler := NewBidHandler(bidService)

	authMiddleware := middleware.JWTAuth(db)
	route := r.Group("/bids")
	{
		route.POST("", authMiddleware, bidHandler.CreateBid)
	}
}
