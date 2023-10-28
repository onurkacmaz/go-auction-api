package repository

import (
	"auction/internal/auction/dto"
	"auction/internal/auction/model"
	"auction/pkg/database"
	"context"
)

type IAuctionRepository interface {
	GetAuctions(ctx context.Context, req *dto.GetAuctionsReq) []*model.Auction
}

type AuctionRepo struct {
	db database.IDatabase
}

func NewAuctionRepository(db database.IDatabase) *AuctionRepo {
	return &AuctionRepo{db: db}
}

func (r *AuctionRepo) GetAuctions(ctx context.Context, req *dto.GetAuctionsReq) []*model.Auction {
	var auctions []*model.Auction

	r.db.Find(
		ctx,
		&auctions,
		database.WithQuery(database.NewQuery("status = ?", model.AuctionStatusActive)),
		database.WithPreload([][]string{{"Artworks", "status = 1"}}),
		database.WithOrder("id desc"),
	)

	return auctions
}
