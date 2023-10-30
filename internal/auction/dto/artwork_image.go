package dto

import "time"

type ArtworkImage struct {
	ID        string    `json:"id"`
	ArtworkId string    `json:"artwork_id"`
	Path      string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
