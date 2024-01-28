package database

import "time"

type Model struct {
	ID        string     `json:"id" gorm:"unique;not null;index;primaryKey;"`
	CreatedAt time.Time  `gorm:"column:created_at;index"`
	UpdatedAt time.Time  `gorm:"column:updated_at;index"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index"`
}
