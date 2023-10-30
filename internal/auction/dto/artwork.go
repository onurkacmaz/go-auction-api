package dto

import (
	"auction/internal/artist/dto"
	"time"
)

type Artwork struct {
	ID                   string          `json:"id"`
	AuctionId            string          `json:"auction_id"`
	Title                string          `json:"title"`
	Description          string          `json:"description"`
	Status               int             `json:"status"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
	StartPrice           float64         `json:"start_price"`
	EndPrice             float64         `json:"end_price"`
	EstimatedMarketPrice float64         `json:"estimated_market_price"`
	Images               []*ArtworkImage `json:"images"`
	Bids                 []*Bid          `json:"bids"`
	Artist               []*dto.Artist   `json:"artist"`
}
