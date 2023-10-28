package model

import (
	"time"
)

const (
	AuctionStatusActive   = "active"
	AuctionStatusInactive = "inactive"
)

type AuctionStatus string

type Auction struct {
	ID          string        `json:"id" gorm:"unique;not null;index;primary_key"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	StartDate   *time.Time    `json:"start_date"`
	EndDate     *time.Time    `json:"end_date"`
	Status      AuctionStatus `json:"status"`
	Image       string        `json:"image"`
	Artworks    []*Artwork    `json:"artworks"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	DeletedAt   time.Time     `json:"deleted_at"`
}
