package service

import (
	"auction/internal/artwork/dto"
	"auction/internal/artwork/model"
	"auction/internal/artwork/repository"
	"auction/pkg/upload"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IArtworkService interface {
	CreateArtwork(c *gin.Context, req *dto.CreateArtworkReq) (*model.Artwork, error)
	UpdateArtwork(c *gin.Context, id uint32, req *dto.UpdateArtworkReq) (*model.Artwork, error)
}

type ArtworkService struct {
	repo repository.IArtworkRepository
}

func NewArtworkService(repo repository.IArtworkRepository) *ArtworkService {
	return &ArtworkService{
		repo: repo,
	}
}

func (s *ArtworkService) GetArtworkByID(c *gin.Context, id uint32) *model.Artwork {
	return s.repo.GetArtworkByID(c, id)
}

func (s *ArtworkService) CreateArtwork(c *gin.Context, req *dto.CreateArtworkReq) (*model.Artwork, error) {
	artwork := &model.Artwork{
		AuctionId:            req.AuctionId,
		Title:                req.Title,
		Description:          req.Description,
		StartPrice:           req.StartPrice,
		EndPrice:             req.EndPrice,
		Status:               req.Status,
		EstimatedMarketPrice: req.EstimatedMarketPrice,
	}

	var images []*model.ArtworkImage
	for _, image := range req.Images {
		path, err := upload.FileBag{
			File: image,
			Name: uuid.New().String(),
			Dest: "public/artwork/",
		}.Upload(c)

		if err != nil {
			images = append(images, &model.ArtworkImage{
				ArtworkId: artwork.ID,
				Path:      path,
			})
		}
	}

	artwork.Images = images

	_, err := s.repo.CreateArtwork(c, artwork)

	return artwork, err
}

func (s *ArtworkService) UpdateArtwork(c *gin.Context, id uint32, req *dto.UpdateArtworkReq) (*model.Artwork, error) {
	return nil, nil
}
