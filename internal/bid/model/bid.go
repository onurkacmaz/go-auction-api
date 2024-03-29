package model

import (
	"auction/pkg/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Bid struct {
	database.Model
	Amount    float64 `json:"amount"`
	UserID    uint32  `json:"user_id" gorm:"index:idx_bid_user_id"`
	ArtworkID uint32  `json:"artwork_id" gorm:"index:idx_bid_artwork_id"`
}

func (u *Bid) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().ID()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return nil
}
