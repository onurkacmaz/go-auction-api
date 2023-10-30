package model

import (
	"auction/internal/artist/model"
	"auction/pkg/database"
)

type Artwork struct {
	database.Model
	AuctionId            string          `json:"auction_id" gorm:"index;not null"`
	Title                string          `json:"title"`
	Description          string          `json:"description"`
	Status               int             `json:"status"`
	StartPrice           float64         `json:"start_price"`
	EndPrice             float64         `json:"end_price"`
	EstimatedMarketPrice float64         `json:"estimated_market_price"`
	Images               []*ArtworkImage `json:"images"`
	Bids                 []*Bid          `json:"bids"`
	Artist               []model.Artist  `json:"artist" gorm:"many2many:artwork_artists;"`
}
