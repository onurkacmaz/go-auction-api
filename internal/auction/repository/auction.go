package repository

import (
	"auction/internal/auction/dto"
	"auction/internal/auction/model"
	"auction/pkg/database"
	"auction/pkg/paging"
	"context"
)

type IAuctionRepository interface {
	GetAuctions(ctx context.Context, req *dto.GetAuctionsReq) ([]*model.Auction, *paging.Pagination, error)
	GetAuctionByID(ctx context.Context, id uint32) *model.Auction
	UpdateAuction(ctx context.Context, auction *model.Auction) (*model.Auction, error)
	CreateAuction(ctx context.Context, auction *model.Auction) (*model.Auction, error)
}

type AuctionRepo struct {
	db database.IDatabase
}

func NewAuctionRepository(db database.IDatabase) *AuctionRepo {
	return &AuctionRepo{db: db}
}

func (r *AuctionRepo) GetAuctions(ctx context.Context, req *dto.GetAuctionsReq) ([]*model.Auction, *paging.Pagination, error) {

	var opts []database.FindOption

	opts = append(opts, database.WithQuery(database.NewQuery("status = ?", model.AuctionStatusActive)))

	var total int64
	if err := r.db.Count(ctx, &model.Auction{}, &total, opts...); err != nil {
		return nil, nil, err
	}

	pagination := paging.New(req.Page, req.Limit, total)

	if req.Preload {
		opts = append(opts, database.WithPreload([][]string{
			{"Artworks", "status = active"},
			{"Artworks.Images", "deleted_at is null"},
			{"Artworks.Bids", "deleted_at is null"},
			{"Artworks.Artist", "deleted_at is null"},
		}))
	}

	opts = append(opts, database.WithOrder("created_at desc"))
	opts = append(opts, database.WithLimit(int(pagination.Limit)))
	opts = append(opts, database.WithOffset(int(pagination.Skip)))

	var auctions []*model.Auction
	_ = r.db.Find(
		ctx,
		&auctions,
		opts...,
	)

	return auctions, pagination, nil
}

func (r *AuctionRepo) GetAuctionByID(ctx context.Context, id uint32) *model.Auction {
	var auction *model.Auction

	var opts []database.FindOption

	opts = append(opts, database.WithQuery(database.NewQuery("id = ?", id)))
	opts = append(opts, database.WithPreload([][]string{
		{"Artworks", "status = 'active'"},
		{"Artworks.Images", "deleted_at is null"},
		{"Artworks.Bids", "deleted_at is null"},
		{"Artworks.Artist", "deleted_at is null"},
	}))

	var err = r.db.FindOne(
		ctx,
		&auction,
		opts...,
	)

	if err != nil {
		return nil
	}

	return auction
}

func (r *AuctionRepo) UpdateAuction(ctx context.Context, auction *model.Auction) (*model.Auction, error) {
	if err := r.db.Update(ctx, auction); err != nil {
		return nil, err
	}

	return auction, nil
}

func (r *AuctionRepo) CreateAuction(ctx context.Context, auction *model.Auction) (*model.Auction, error) {
	if err := r.db.Create(ctx, auction); err != nil {
		return nil, err
	}

	return auction, nil
}
