package dto

type Bid struct {
	ID     uint32  `json:"id"`
	Amount float64 `json:"amount"`
}

type CreateBidRequest struct {
	Amount    float64 `json:"amount"`
	ArtworkID uint32  `json:"artwork_id"`
}

type CreateBidResponse struct {
	Bid      *Bid    `json:"bid"`
	EndPrice float64 `json:"end_price"`
	Message  string  `json:"message"`
	BidCount int     `json:"bid_count"`
	MinBid   struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"min_bid"`
}
