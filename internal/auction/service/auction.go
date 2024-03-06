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
	GetAuctionByID(c *gin.Context, id uint32) (*model.Auction, error)
	CreateAuction(c *gin.Context, req *dto.CreateAuctionReq) (*model.Auction, error)
	UpdateAuction(c *gin.Context, id uint32, req *dto.UpdateAuctionReq) (*model.Auction, error)
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

func (s *AuctionService) GetAuctionByID(c *gin.Context, id uint32) (*model.Auction, error) {
	auction := s.repo.GetAuctionByID(c, id)

	if auction == nil {
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

func (s *AuctionService) UpdateAuction(c *gin.Context, id uint32, req *dto.UpdateAuctionReq) (*model.Auction, error) {
	auction := s.repo.GetAuctionByID(c, id)

	if auction == nil {
		return nil, errors.New("auction not found")
	}

	if req.Image != nil {
		path, err := upload.FileBag{
			File: *req.Image,
			Name: uuid.New().String(),
			Dest: "public/",
		}.Upload(c)

		if err != nil {
			return nil, err
		}

		auction.Image = path
	}
	if req.Name != nil {
		auction.Name = *req.Name
	}
	if req.Description != nil {
		auction.Description = *req.Description
	}
	if req.StartDate != nil {
		var startDate, _ = time.Parse("2006-01-02 15:04:05", *req.StartDate)
		auction.StartDate = &startDate
	}
	if req.EndDate != nil {
		var endDate, _ = time.Parse("2006-01-02 15:04:05", *req.EndDate)
		auction.EndDate = &endDate
	}
	if req.Status != nil {
		auction.Status = *req.Status
	}

	var _, err = s.repo.UpdateAuction(c, auction)

	if err != nil {
		return nil, err
	}

	return auction, nil
}
