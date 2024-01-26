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
