package model

import (
	"auction/pkg/database"
	"auction/pkg/utils"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type User struct {
	database.Model
	Email           string         `json:"email" gorm:"unique;not null;index:idx_user_email"`
	Password        string         `json:"password"`
	Roles           datatypes.JSON `json:"roles" gorm:"type:json;not null;"`
	PhoneNumber     string         `json:"phone_number"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"`
	BirthDate       *time.Time     `json:"birth_date"`
}

const (
	AdminRole = "ROLE_ADMIN"
	UserRole  = "ROLE_USER"
)

type RoleJSON []string

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	u.Password = utils.HashPassword([]byte(u.Password))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	roleJson, _ := json.Marshal(&RoleJSON{UserRole})
	u.Roles = datatypes.JSON([]byte(string(roleJson)))

	return nil
}
