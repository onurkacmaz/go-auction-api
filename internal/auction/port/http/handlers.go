package http

import (
	"auction/internal/auction/dto"
	"auction/internal/auction/service"
	"auction/pkg/response"
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

	res, err := h.service.GetAuctions(c, &req)
	if err != nil {
		log.Println("Failed to get auctions ", err)
		response.Error(c, http.StatusBadRequest, err, err.Error())
		return
	}

	response.JSON(c, http.StatusOK, res)
}
