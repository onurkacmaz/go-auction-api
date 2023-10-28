package model

import (
	"time"
)

type Artwork struct {
	ID                   string     `json:"id" gorm:"unique;not null;index;primary_key"`
	AuctionId            string     `json:"auction_id" gorm:"index;not null"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Status               int        `json:"status"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at"`
	StartPrice           float64    `json:"start_price"`
	EndPrice             float64    `json:"end_price"`
	EstimatedMarketPrice float64    `json:"estimated_market_price"`
}
