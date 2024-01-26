package dto

import (
	"auction/pkg/paging"
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
	Preload bool  `form:"preload" default:"true"`
	Page    int64 `form:"page" default:"1"`
	Limit   int64 `form:"limit" default:"10"`
}

type GetAuctionsRes struct {
	Pagination *paging.Pagination `json:"pagination"`
	Auctions   []*Auction         `json:"auctions"`
}

type GetAuctionByIDReq struct {
	ID string `uri:"id" binding:"required"`
}

type GetAuctionRes struct {
	Auction *Auction `json:"auction"`
}
