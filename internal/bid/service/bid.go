package service

import (
	artworkModel "auction/internal/artwork/model"
	artworkRepository "auction/internal/artwork/repository"
	auctionRepository "auction/internal/auction/repository"
	"auction/internal/bid/dto"
	bidModel "auction/internal/bid/model"
	bidRepository "auction/internal/bid/repository"
	userModel "auction/internal/user/model"
	"context"
	"errors"
	"log"
)

type IBidService interface {
	CreateBid(ctx context.Context, user userModel.User, artwork artworkModel.Artwork, req *dto.CreateBidRequest) (*bidModel.Bid, error)
	GetArtworkByID(ctx context.Context, id uint32) *artworkModel.Artwork
}

type Service struct {
	bidRepo     bidRepository.IBidRepository
	artworkRepo artworkRepository.IArtworkRepository
	auctionRepo auctionRepository.IAuctionRepository
}

func (s *Service) GetArtworkByID(ctx context.Context, id uint32) *artworkModel.Artwork {
	return s.artworkRepo.GetArtworkByID(ctx, id)
}

func NewBidService(
	bidRepo bidRepository.IBidRepository,
	artworkRepo artworkRepository.IArtworkRepository,
	auctionRepo auctionRepository.IAuctionRepository,
) *Service {
	return &Service{
		bidRepo:     bidRepo,
		artworkRepo: artworkRepo,
		auctionRepo: auctionRepo,
	}
}

func (s *Service) CreateBid(ctx context.Context, user userModel.User, artwork artworkModel.Artwork, req *dto.CreateBidRequest) (*bidModel.Bid, error) {
	auction := s.auctionRepo.GetAuctionByID(ctx, artwork.AuctionId)

	if auction == nil {
		log.Println("Auction not found")
		return nil, errors.New("auction not found")
	}

	if auction.IsActive() == false {
		log.Println("Auction is not active")
		return nil, errors.New("auction is not active")
	}

	bid, err := s.bidRepo.CreateBid(ctx, user, req)

	if err != nil {
		return nil, err
	}

	if auction.IsRangeEndDateIsLastXMins(5) == true {
		auction := auction.ExtendEndDate(5)
		_, _ = s.auctionRepo.UpdateAuction(ctx, auction)
	}

	artwork.EndPrice = bid.Amount

	_, _ = s.artworkRepo.UpdateArtwork(ctx, &artwork)

	return bid, nil
}
