package repository

import (
	"auction/internal/auction/dto"
	"auction/internal/auction/model"
	"auction/pkg/database"
	"context"
)

type IAuctionRepository interface {
	GetAuctions(ctx context.Context, req *dto.GetAuctionsReq) []*model.Auction
	GetAuctionByID(ctx context.Context, id string) *model.Auction
}

type AuctionRepo struct {
	db database.IDatabase
}

func NewAuctionRepository(db database.IDatabase) *AuctionRepo {
	return &AuctionRepo{db: db}
}

func (r *AuctionRepo) GetAuctions(ctx context.Context, req *dto.GetAuctionsReq) []*model.Auction {
	var auctions []*model.Auction

	var opts []database.FindOption

	if req.Preload {
		opts = append(opts, database.WithPreload([][]string{
			{"Artworks", "status = 1"},
			{"Artworks.Images", "deleted_at is null"},
			{"Artworks.Bids", "deleted_at is null"},
			{"Artworks.Artist", "deleted_at is null"},
		}))
	}

	opts = append(opts, database.WithQuery(database.NewQuery("status = ?", model.AuctionStatusActive)))
	opts = append(opts, database.WithOrder("id desc"))

	r.db.Find(
		ctx,
		&auctions,
		opts...,
	)

	return auctions
}

func (r AuctionRepo) GetAuctionByID(ctx context.Context, id string) *model.Auction {
	var auction model.Auction

	var opts []database.FindOption

	opts = append(opts, database.WithQuery(database.NewQuery("id = ?", id)))
	opts = append(opts, database.WithPreload([][]string{
		{"Artworks", "status = 1"},
		{"Artworks.Images", "deleted_at is null"},
		{"Artworks.Bids", "deleted_at is null"},
		{"Artworks.Artist", "deleted_at is null"},
	}))

	r.db.FindOne(
		ctx,
		&auction,
		opts...,
	)

	return &auction
}
