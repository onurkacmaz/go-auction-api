package dto

import "time"

type ArtworkImage struct {
	ID        uint32    `json:"id"`
	ArtworkId uint32    `json:"artwork_id"`
	Path      string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
