package dto

import (
	"auction/internal/artwork/model"
	bidModel "auction/internal/bid/model"
	"mime/multipart"
	"time"
)

type Artwork struct {
	ID                   uint32          `json:"id"`
	AuctionId            uint32          `json:"auction_id"`
	Title                string          `json:"title"`
	Description          string          `json:"description"`
	Status               int             `json:"status"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
	StartPrice           float64         `json:"start_price"`
	EndPrice             float64         `json:"end_price"`
	EstimatedMarketPrice float64         `json:"estimated_market_price"`
	Images               []*ArtworkImage `json:"images"`
	Bids                 []*bidModel.Bid `json:"bids"`
}

type CreateArtworkReq struct {
	AuctionId            uint32                 `form:"auction_id" binding:"required"`
	Title                string                 `form:"title" binding:"required"`
	Description          string                 `form:"description" binding:"required"`
	StartPrice           float64                `form:"start_price" binding:"required"`
	EndPrice             float64                `form:"end_price" binding:"required"`
	Images               []multipart.FileHeader `form:"images" binding:"required"`
	Status               model.ArtworkStatus    `form:"status" binding:"required"`
	EstimatedMarketPrice float64                `form:"estimated_market_price" binding:"required"`
}

type UpdateArtworkReq struct {
	AuctionId            uint32                 `form:"auction_id" binding:"required"`
	Title                string                 `form:"name" binding:"required"`
	Description          string                 `form:"description" binding:"required"`
	StartPrice           float64                `form:"start_price" binding:"required"`
	EndPrice             float64                `form:"end_price" binding:"required"`
	Images               []multipart.FileHeader `form:"image" binding:"required"`
	Status               model.ArtworkStatus    `form:"status" binding:"required"`
	EstimatedMarketPrice float64                `form:"estimated_market_price" binding:"required"`
}

type GetArtworkRes struct {
	Artwork Artwork `json:"artwork"`
}
