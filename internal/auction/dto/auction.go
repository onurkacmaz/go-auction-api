package dto

import (
	"auction/internal/auction/model"
	"time"
)

type Auction struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Status      string     `json:"status"`
	Image       string     `json:"image"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   time.Time  `json:"deleted_at"`
}

type GetAuctionsReq struct {
}

type GetAuctionsRes struct {
	Auctions []*model.Auction `json:"auctions"`
}
