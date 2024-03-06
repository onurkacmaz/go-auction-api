package http

import (
	"auction/internal/auction/dto"
	"auction/internal/auction/service"
	"auction/pkg/config"
	"auction/pkg/redis"
	"auction/pkg/response"
	"auction/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type AuctionHandler struct {
	service service.IAuctionService
	cache   redis.IRedis
}

func NewAuctionHandler(service service.IAuctionService, cache redis.IRedis) *AuctionHandler {
	return &AuctionHandler{service: service, cache: cache}
}

func (h *AuctionHandler) GetAuctions(c *gin.Context) {
	var req dto.GetAuctionsReq
	var res dto.GetAuctionsRes

	cacheKey := c.Request.URL.RequestURI()
	if err := h.cache.Get(cacheKey, &res); err == nil {
		response.JSON(c, http.StatusOK, res)
		return
	}

	if err := c.ShouldBindQuery(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	auctions, pagination, err := h.service.GetAuctions(c, &req)

	if err != nil {
		log.Println("Failed to get auctions ", err)
		response.Error(c, http.StatusBadRequest, err, err.Error())
		return
	}

	res.Pagination = pagination
	utils.Copy(&res.Auctions, &auctions)

	response.JSON(c, http.StatusOK, res)
	//_ = h.cache.SetWithExpiration(cacheKey, res, config.AuctionsCachingTime)
}

func (h *AuctionHandler) GetAuctionByID(c *gin.Context) {
	var res dto.GetAuctionRes

	cacheKey := c.Request.URL.RequestURI()
	if err := h.cache.Get(cacheKey, &res.Auction); err == nil {
		response.JSON(c, http.StatusOK, res)
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	auction, err := h.service.GetAuctionByID(c, uint32(id))

	if err != nil {
		log.Println("Failed to get auction ", err)
		response.JSON(c, http.StatusNotFound, err.Error())
		return
	}

	utils.Copy(&res.Auction, &auction)
	response.JSON(c, http.StatusOK, res)
	_ = h.cache.SetWithExpiration(cacheKey, res.Auction, config.AuctionCachingTime)
}

func (h *AuctionHandler) CreateAuction(c *gin.Context) {
	var req dto.CreateAuctionReq

	if err := c.ShouldBind(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	auction, err := h.service.CreateAuction(c, &req)

	if err != nil {
		log.Println("Failed to create auction ", err)
		response.Error(c, http.StatusUnprocessableEntity, err, err.Error())
		return
	}

	var res dto.GetAuctionRes
	utils.Copy(&res.Auction, &auction)

	response.JSON(c, http.StatusCreated, res)
}

func (h *AuctionHandler) UpdateAuction(c *gin.Context) {
	var req dto.UpdateAuctionReq

	if err := c.ShouldBind(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	auction, err := h.service.UpdateAuction(c, uint32(id), &req)

	if err != nil {
		log.Println("Failed to update auction ", err)
		response.Error(c, http.StatusUnprocessableEntity, err, err.Error())
		return
	}

	var res dto.GetAuctionRes
	utils.Copy(&res.Auction, &auction)

	response.JSON(c, http.StatusOK, res)
}
