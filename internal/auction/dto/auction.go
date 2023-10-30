package dto

import (
	"time"
)

type Auction struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Status      string     `json:"status"`
	Image       string     `json:"image"`
	Artworks    []*Artwork `json:"artworks"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type GetAuctionsReq struct {
	Preload bool `form:"preload" default:"true"`
}

type GetAuctionsRes struct {
	Auctions []*Auction `json:"auctions"`
}

type GetAuctionByIDReq struct {
	ID string `uri:"id" binding:"required"`
}

type GetAuctionRes struct {
	Auction *Auction `json:"auction"`
}
