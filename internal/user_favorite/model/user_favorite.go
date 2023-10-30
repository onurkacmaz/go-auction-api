package model

import "auction/pkg/database"

type UserFavorite struct {
	database.Model
	UserID    uint64 `gorm:"not null"`
	ArtworkID uint64 `gorm:"not null"`
}
