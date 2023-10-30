package model

import (
	"auction/pkg/database"
	"time"
)

const (
	AuctionStatusActive   = "active"
	AuctionStatusInactive = "inactive"
)

type AuctionStatus string

type Auction struct {
	database.Model
	Name        string        `json:"name"`
	Description string        `json:"description"`
	StartDate   *time.Time    `json:"start_date"`
	EndDate     *time.Time    `json:"end_date"`
	Status      AuctionStatus `json:"status"`
	Image       string        `json:"image"`
	Artworks    []*Artwork    `json:"artworks"`
}
