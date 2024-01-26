package repository

import (
	artworkModel "auction/internal/auction/model"
	"auction/pkg/database"
	"context"
)

type IArtworkRepository interface {
	GetArtworkByID(ctx context.Context, id string) *artworkModel.Artwork
	UpdateArtwork(ctx context.Context, artwork *artworkModel.Artwork) (*artworkModel.Artwork, error)
}

type ArtworkRepo struct {
	db database.IDatabase
}

func (a ArtworkRepo) GetArtworkByID(ctx context.Context, id string) *artworkModel.Artwork {
	var opts []database.FindOption

	opts = append(opts, database.WithQuery(database.NewQuery("id = ? and status = 1", id)))

	opts = append(opts, database.WithPreload([][]string{
		{"Bids", "deleted_at is null order by created_at asc"},
	}))

	var artwork artworkModel.Artwork
	err := a.db.FindOne(ctx, &artwork, opts...)

	if err != nil {
		return nil
	}

	return &artwork
}

func NewArtworkRepository(db database.IDatabase) *ArtworkRepo {
	return &ArtworkRepo{db: db}
}

func (a ArtworkRepo) UpdateArtwork(ctx context.Context, artwork *artworkModel.Artwork) (*artworkModel.Artwork, error) {
	err := a.db.GetDB().WithContext(ctx).Save(artwork).Error
	if err != nil {
		return nil, err
	}
	return artwork, nil
}
