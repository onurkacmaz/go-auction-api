package model

import "auction/pkg/database"

type UserFollow struct {
	database.Model
	UserID    uint64 `gorm:"not null"`
	ArtworkID uint64 `gorm:"not null"`
}
