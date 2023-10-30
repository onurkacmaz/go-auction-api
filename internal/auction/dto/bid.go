package dto

import "time"

type Bid struct {
	ID        string    `json:"id"`
	ArtworkId string    `json:"artwork_id"`
	UserId    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
