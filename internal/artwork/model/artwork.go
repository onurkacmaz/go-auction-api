package model

import (
	artistModel "auction/internal/artist/model"
	bidModel "auction/internal/bid/model"
	"auction/pkg/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
)

type ArtworkStatus string

type Artwork struct {
	database.Model
	AuctionId            uint32               `json:"auction_id" gorm:"index;not null"`
	Title                string               `json:"title"`
	Description          string               `json:"description"`
	Status               ArtworkStatus        `json:"status"`
	StartPrice           float64              `json:"start_price"`
	EndPrice             float64              `json:"end_price"`
	EstimatedMarketPrice float64              `json:"estimated_market_price"`
	Images               []*ArtworkImage      `json:"images"`
	Bids                 []*bidModel.Bid      `json:"bids"`
	Artist               []artistModel.Artist `json:"artist" gorm:"many2many:artwork_artists;"`
}

type MinBid struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func (a *Artwork) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New().ID()

	return nil
}

func (a *Artwork) MinBid() MinBid {
	lastBid := a.Bids[len(a.Bids)-1]

	return MinBid{
		Amount:   lastBid.Amount + 1,
		Currency: "USD",
	}
}
