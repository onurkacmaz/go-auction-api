package http

import (
	"auction/internal/bid/dto"
	"auction/internal/bid/service"
	"auction/internal/user/model"
	"auction/pkg/response"
	"auction/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type BidHandler struct {
	service service.IBidService
}

func NewBidHandler(service service.IBidService) *BidHandler {
	return &BidHandler{service: service}
}

func (h *BidHandler) CreateBid(c *gin.Context) {
	var req dto.CreateBidRequest

	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	artwork, err := h.service.GetArtworkByID(c, req.ArtworkID)
	if artwork == nil {
		response.Error(c, http.StatusNotFound, err, "Artwork not found")
		return
	}

	var user model.User
	utils.Copy(c.MustGet("user"), &user)
	bid, err := h.service.CreateBid(c, user, *artwork, &req)

	if err != nil {
		log.Println("Failed to create bid ", err)
		response.Error(c, http.StatusBadRequest, err, err.Error())
		return
	}

	var res dto.CreateBidResponse
	utils.Copy(&res.Bid, &bid)

	minBid := artwork.MinBid()

	res.EndPrice = artwork.EndPrice
	res.Message = "Bid created successfully"
	res.BidCount = len(artwork.Bids)
	res.MinBid.Amount = minBid.Amount
	res.MinBid.Currency = minBid.Currency

	response.JSON(c, http.StatusOK, res)
}
