package middleware

import (
	"auction/internal/user/model"
	"auction/internal/user/repository"
	"auction/pkg/database"
	"auction/pkg/jtoken"
	"auction/pkg/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "reflect"
)

type IUserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func JWTAuth(db database.IDatabase) gin.HandlerFunc {
	return JWT(jtoken.AccessTokenType, db)
}

func JWTRefresh() gin.HandlerFunc {
	return JWT(jtoken.RefreshTokenType, nil)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetUserByID(ctx, id)
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func JWT(tokenType string, db database.IDatabase) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}

		payload, err := jtoken.ValidateToken(token)
		if err != nil || payload == nil || payload["type"] != tokenType {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			c.Abort()
			return
		}

		if db != nil {
			userService := NewUserService(repository.NewUserRepository(db))
			user, _ := userService.GetUserByID(c, payload["id"].(string))

			if user == nil {
				c.JSON(http.StatusUnauthorized, Response{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				})
				c.Abort()
				return
			}

			c.Set("user", user)
		}

		var roles []string
		for _, role := range payload["roles"].([]interface{}) {
			if str, ok := role.(string); ok {
				roles = append(roles, str)
			}
		}

		c.Set("userId", payload["id"])
		c.Set("roles", roles)
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.InArray("ROLE_ADMIN", c.GetStringSlice("roles")) {
			c.JSON(http.StatusForbidden, Response{
				Code:    http.StatusForbidden,
				Message: "Forbidden",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
