package model

import "auction/pkg/database"

type ArtworkImage struct {
	database.Model
	ArtworkId string `json:"artwork_id" gorm:"index;not null"`
	Path      string `json:"image"`
}
