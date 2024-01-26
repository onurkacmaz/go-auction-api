package repository

import (
	"auction/internal/bid/dto"
	bidModel "auction/internal/bid/model"
	userModel "auction/internal/user/model"
	"auction/pkg/database"
	"context"
)

type IBidRepository interface {
	CreateBid(ctx context.Context, user userModel.User, req *dto.CreateBidRequest) (*bidModel.Bid, error)
}

type BidRepo struct {
	db database.IDatabase
}

func NewBidRepository(db database.IDatabase) *BidRepo {
	return &BidRepo{db: db}
}

func (r *BidRepo) CreateBid(ctx context.Context, user userModel.User, req *dto.CreateBidRequest) (*bidModel.Bid, error) {
	bid := &bidModel.Bid{
		ArtworkID: req.ArtworkID,
		UserID:    user.ID,
		Amount:    req.Amount,
	}

	if err := r.db.Create(ctx, bid); err != nil {
		return nil, err
	}

	return bid, nil
}
