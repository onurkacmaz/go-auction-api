package service

import (
	"auction/internal/auction/dto"
	"auction/internal/auction/model"
	"auction/internal/auction/repository"
	"auction/pkg/paging"
	"errors"
	"github.com/gin-gonic/gin"
)

type IAuctionService interface {
	GetAuctions(c *gin.Context, req *dto.GetAuctionsReq) ([]*model.Auction, *paging.Pagination, error)
	GetAuctionByID(c *gin.Context, id string) (*model.Auction, error)
}

type AuctionService struct {
	repo repository.IAuctionRepository
}

func NewAuctionService(repo repository.IAuctionRepository) *AuctionService {
	return &AuctionService{
		repo: repo,
	}
}

func (s *AuctionService) GetAuctions(c *gin.Context, req *dto.GetAuctionsReq) ([]*model.Auction, *paging.Pagination, error) {
	auctions, pagination, err := s.repo.GetAuctions(c, req)

	if err != nil {
		return nil, nil, err
	}

	return auctions, pagination, err
}

func (s *AuctionService) GetAuctionByID(c *gin.Context, id string) (*model.Auction, error) {
	auction := s.repo.GetAuctionByID(c, id)

	if auction.ID == "" {
		return nil, errors.New("auction not found")
	}

	return auction, nil
}
