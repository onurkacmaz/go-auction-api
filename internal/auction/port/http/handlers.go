package http

import (
	"auction/internal/auction/dto"
	"auction/internal/auction/service"
	"auction/pkg/response"
	"auction/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuctionHandler struct {
	service service.IAuctionService
}

func NewAuctionHandler(service service.IAuctionService) *AuctionHandler {
	return &AuctionHandler{service: service}
}

func (h *AuctionHandler) GetAuctions(c *gin.Context) {
	var req dto.GetAuctionsReq
	if err := c.ShouldBindQuery(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	auctions, err := h.service.GetAuctions(c, &req)

	var res dto.GetAuctionsRes
	utils.Copy(&res.Auctions, &auctions)

	if err != nil {
		log.Println("Failed to get auctions ", err)
		response.Error(c, http.StatusBadRequest, err, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, res)
}

func (h *AuctionHandler) GetAuctionByID(c *gin.Context) {
	auction, err := h.service.GetAuctionByID(c, c.Param("id"))

	var res dto.GetAuctionRes
	utils.Copy(&res.Auction, &auction)

	if err != nil {
		log.Println("Failed to get auction ", err)
		response.Error(c, http.StatusNotFound, err, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, res)
}
