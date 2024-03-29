package model

import (
	artworkModel "auction/internal/artwork/model"
	bidModel "auction/internal/bid/model"
	"auction/pkg/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

const (
	AuctionStatusActive   = "active"
	AuctionStatusInactive = "inactive"
)

type AuctionStatus string

type Auction struct {
	database.Model
	Name        string                  `json:"name"`
	Slug        string                  `json:"slug"`
	Description string                  `json:"description"`
	StartDate   *time.Time              `json:"start_date"`
	EndDate     *time.Time              `json:"end_date"`
	Status      AuctionStatus           `json:"status"`
	Image       string                  `json:"image"`
	Artworks    []*artworkModel.Artwork `json:"artworks" gorm:"one2many:artworks;"`
	Bids        []*bidModel.Bid         `json:"bids" gorm:"many2many:bids;"`
}

func (a *Auction) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New().ID()
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	return nil
}

func (a *Auction) BeforeUpdate(tx *gorm.DB) error {
	a.UpdatedAt = time.Now()

	return nil
}

func (a *Auction) IsActive() bool {
	return a.Status == AuctionStatusActive && a.StartDate.Before(time.Now()) && a.EndDate.After(time.Now())
}

func (a *Auction) IsRangeEndDateIsLastXMins(minutes int) bool {
	return a.EndDate.Before(time.Now().Add(time.Duration(minutes) * time.Minute))
}

func (a *Auction) ExtendEndDate(minutes int) *Auction {
	var endDate = a.EndDate.Add(time.Duration(minutes) * time.Minute)
	a.EndDate = &endDate

	return a
}
