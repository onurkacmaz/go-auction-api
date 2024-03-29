package dto

import (
	artworkDto "auction/internal/artwork/dto"
	"auction/internal/auction/model"
	"auction/pkg/paging"
	"mime/multipart"
	"time"
)

type Auction struct {
	ID          uint32                `json:"id"`
	Name        string                `json:"name"`
	Slug        string                `json:"slug"`
	Description string                `json:"description"`
	StartDate   *time.Time            `json:"start_date"`
	EndDate     *time.Time            `json:"end_date"`
	Status      model.AuctionStatus   `json:"status"`
	Image       string                `json:"image"`
	Artworks    []*artworkDto.Artwork `json:"artworks"`
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
	ID uint32 `uri:"id" binding:"required"`
}

type GetAuctionRes struct {
	Auction *Auction `json:"auction"`
}

type CreateAuctionReq struct {
	Name        string               `form:"name" binding:"required"`
	Description string               `form:"description" binding:"required"`
	StartDate   string               `form:"start_date" binding:"required"`
	EndDate     string               `form:"end_date" binding:"required"`
	Status      model.AuctionStatus  `form:"status" binding:"required"`
	Image       multipart.FileHeader `form:"image" binding:"required"`
}

type UpdateAuctionReq struct {
	Name        *string               `form:"name"`
	Description *string               `form:"description"`
	StartDate   *string               `form:"start_date"`
	EndDate     *string               `form:"end_date"`
	Status      *model.AuctionStatus  `form:"status"`
	Image       *multipart.FileHeader `form:"image"`
}
