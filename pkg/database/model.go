package database

import "time"

type Model struct {
	ID        string    `json:"id" gorm:"unique;not null;index;primary_key"`
	CreatedAt CreatedAt `gorm:"column:created_at"`
	UpdatedAt UpdatedAt `gorm:"column:updated_at"`
	DeletedAt DeletedAt `gorm:"column:deleted_at"`
}

type CreatedAt time.Time
type UpdatedAt time.Time
type DeletedAt *time.Time
type ID string
