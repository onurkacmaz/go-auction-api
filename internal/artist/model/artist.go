package model

import (
	"auction/pkg/database"
)

type Artist struct {
	database.Model
	Name    string `json:"name"`
	Bio     string `json:"bio" gorm:"type:text"`
	Email   string `json:"email"`
	Phone   string `json:"phone" gorm:"type:varchar(12)"`
	Address string `json:"address" gorm:"type:text"`
}
