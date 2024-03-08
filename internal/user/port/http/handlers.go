package http

import (
	"auction/internal/user/dto"
	"auction/internal/user/service"
	"auction/pkg/response"
	"auction/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid credentials")
		return
	}

	user, accessToken, refreshToken, expiresIn, err := h.service.Login(c, &req)
	if err != nil {
		log.Println("Failed to login ", err)
		response.Error(c, http.StatusBadRequest, err, err.Error())
		return
	}

	var res dto.LoginRes
	utils.Copy(&res.User, &user)
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	res.ExpiresIn = expiresIn
	response.JSON(c, http.StatusOK, res)
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	user, err := h.service.Register(c, &req)
	if err != nil {
		log.Println(err.Error())
		response.Error(c, http.StatusBadRequest, err, err.Error())
		return
	}

	var res dto.RegisterRes
	utils.Copy(&res.User, &user)
	response.JSON(c, http.StatusOK, res)
}

func (h *UserHandler) GetMe(c *gin.Context) {
	userID := c.GetUint("userId")

	user, err := h.service.GetUserByID(c, uint32(userID))
	if err != nil {
		log.Println(err.Error())
		response.Error(c, http.StatusBadRequest, err, err.Error())
		return
	}

	var res dto.User
	utils.Copy(&res, &user)
	response.JSON(c, http.StatusOK, res)
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	userID := c.GetUint("userId")

	accessToken, expiresIn, err := h.service.RefreshToken(c, uint32(userID))
	if err != nil {
		log.Println("Failed to refresh token", err)
		response.Error(c, http.StatusBadRequest, err, err.Error())
		return
	}

	res := dto.RefreshTokenRes{
		AccessToken: accessToken,
		ExpiresIn:   expiresIn,
	}
	response.JSON(c, http.StatusOK, res)
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	var req dto.ChangePasswordReq
	if err := c.ShouldBindJSON(&req); c.Request.Body == nil || err != nil {
		log.Println("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	userID := c.GetUint("userId")
	err := h.service.ChangePassword(c, uint32(userID), &req)
	if err != nil {
		log.Println(err.Error())
		response.Error(c, http.StatusBadRequest, err, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, nil)
}
