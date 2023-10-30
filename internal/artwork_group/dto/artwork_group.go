package dto

type ArtworkGroup struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Begin float64 `json:"begin"`
	End   float64 `json:"end"`
	Order int     `json:"order"`
}
