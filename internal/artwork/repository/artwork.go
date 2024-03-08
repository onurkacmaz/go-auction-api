package repository

import (
	model "auction/internal/artwork/model"
	"auction/pkg/database"
	"context"
)

type IArtworkRepository interface {
	GetArtworkByID(ctx context.Context, id uint32) (*model.Artwork, error)
	CreateArtwork(ctx context.Context, artwork *model.Artwork) (*model.Artwork, error)
	UpdateArtwork(ctx context.Context, artwork *model.Artwork) (*model.Artwork, error)
}

type ArtworkRepo struct {
	db database.IDatabase
}

func NewArtworkRepository(db database.IDatabase) *ArtworkRepo {
	return &ArtworkRepo{db: db}
}

func (a ArtworkRepo) UpdateArtwork(ctx context.Context, artwork *model.Artwork) (*model.Artwork, error) {
	err := a.db.GetDB().WithContext(ctx).Save(artwork).Error
	if err != nil {
		return nil, err
	}
	return artwork, nil
}

func (a ArtworkRepo) CreateArtwork(ctx context.Context, artwork *model.Artwork) (*model.Artwork, error) {
	err := a.db.GetDB().WithContext(ctx).Save(artwork).Error
	if err != nil {
		return nil, err
	}
	return artwork, nil
}

func (a ArtworkRepo) GetArtworkByID(ctx context.Context, id uint32) (*model.Artwork, error) {
	var opts []database.FindOption

	opts = append(opts, database.WithQuery(database.NewQuery("id = ?", id), database.NewQuery("status = ?", "active")))

	opts = append(opts, database.WithPreload([][]string{
		{"Images", "deleted_at is null"},
		{"Bids", "deleted_at is null order by created_at asc"},
	}))

	var artwork model.Artwork
	err := a.db.FindOne(ctx, &artwork, opts...)

	if err != nil {
		return nil, err
	}

	return &artwork, nil
}
