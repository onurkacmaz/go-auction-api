package model

import (
	"auction/pkg/database"
)

type Bid struct {
	database.Model
	ArtworkId string  `json:"artwork_id" gorm:"index;not null"`
	UserId    string  `json:"user_id" gorm:"index;not null"`
	Amount    float64 `json:"amount"`
}
