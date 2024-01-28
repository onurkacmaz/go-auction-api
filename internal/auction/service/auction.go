package service

import (
	"auction/internal/auction/dto"
	"auction/internal/auction/model"
	"auction/internal/auction/repository"
	"auction/pkg/paging"
	"auction/pkg/upload"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type IAuctionService interface {
	GetAuctions(c *gin.Context, req *dto.GetAuctionsReq) ([]*model.Auction, *paging.Pagination, error)
	GetAuctionByID(c *gin.Context, id string) (*model.Auction, error)
	CreateAuction(c *gin.Context, req *dto.CreateAuctionReq) (*model.Auction, error)
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

func (s *AuctionService) CreateAuction(c *gin.Context, req *dto.CreateAuctionReq) (*model.Auction, error) {
	path, err := upload.FileBag{
		File: req.Image,
		Name: uuid.New().String(),
		Dest: "public/",
	}.Upload(c)

	if err != nil {
		return nil, err
	}

	var startDate, _ = time.Parse("2006-01-02 15:04:05", req.StartDate)
	var endDate, _ = time.Parse("2006-01-02 15:04:05", req.EndDate)

	auction := &model.Auction{
		Name:        req.Name,
		Slug:        req.Name,
		Description: req.Description,
		StartDate:   &startDate,
		EndDate:     &endDate,
		Status:      req.Status,
		Image:       path,
	}

	_, err = s.repo.CreateAuction(c, auction)

	return auction, err
}
