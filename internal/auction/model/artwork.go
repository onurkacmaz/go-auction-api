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

type MinBid struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func (a *Artwork) MinBid() MinBid {
	lastBid := a.Bids[len(a.Bids)-1]

	return MinBid{
		Amount:   lastBid.Amount + 1,
		Currency: "USD",
	}
}
