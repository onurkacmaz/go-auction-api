package model

import (
	"auction/pkg/database"
)

type ArtworkGroup struct {
	database.Model
	Title string  `json:"title"`
	Begin float64 `json:"begin"`
	End   float64 `json:"end"`
	Order int     `json:"order"`
}
