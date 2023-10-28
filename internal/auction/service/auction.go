package service

import (
	"auction/internal/auction/dto"
	"auction/internal/auction/repository"
	"github.com/gin-gonic/gin"
)

type IAuctionService interface {
	GetAuctions(c *gin.Context, req *dto.GetAuctionsReq) (*dto.GetAuctionsRes, error)
}

type AuctionService struct {
	repo repository.IAuctionRepository
}

func NewAuctionService(repo repository.IAuctionRepository) *AuctionService {
	return &AuctionService{
		repo: repo,
	}
}

func (s *AuctionService) GetAuctions(c *gin.Context, req *dto.GetAuctionsReq) (*dto.GetAuctionsRes, error) {
	auctions := s.repo.GetAuctions(c, req)

	return &dto.GetAuctionsRes{
		Auctions: auctions,
	}, nil
}
