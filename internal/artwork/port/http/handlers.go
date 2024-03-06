package http

import (
	"auction/internal/artwork/dto"
	"auction/internal/artwork/service"
	"auction/pkg/redis"
	"auction/pkg/response"
	"auction/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ArtworkHandler struct {
	service service.IArtworkService
	cache   redis.IRedis
}

func NewArtworkHandler(service service.IArtworkService, cache redis.IRedis) *ArtworkHandler {
	return &ArtworkHandler{service: service, cache: cache}
}

func (h *ArtworkHandler) CreateArtwork(c *gin.Context) {
	var req dto.CreateArtworkReq

	if err := c.ShouldBind(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	artwork, err := h.service.CreateArtwork(c, &req)

	if err != nil {
		log.Println("Failed to create artwork ", err)
		response.Error(c, http.StatusUnprocessableEntity, err, err.Error())
		return
	}

	var res dto.GetArtworkRes
	utils.Copy(&res.Artwork, &artwork)

	response.JSON(c, http.StatusCreated, res)
}

func (h *ArtworkHandler) UpdateAuction(c *gin.Context) {
	var req dto.UpdateArtworkReq

	if err := c.ShouldBind(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	artwork, err := h.service.UpdateArtwork(c, uint32(c.GetUint("id")), &req)

	if err != nil {
		log.Println("Failed to update artwork ", err)
		response.Error(c, http.StatusUnprocessableEntity, err, err.Error())
		return
	}

	var res dto.GetArtworkRes
	utils.Copy(&res.Artwork, &artwork)

	response.JSON(c, http.StatusOK, res)
}
