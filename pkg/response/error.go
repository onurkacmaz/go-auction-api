package response

import (
	"auction/pkg/config"
	"github.com/gin-gonic/gin"
	"strings"
)

func Error(c *gin.Context, status int, err error, message string) {
	cfg := config.GetConfig()
	errorRes := map[string]interface{}{
		"message": strings.ToUpper(message[:1]) + strings.ToLower(message[1:]),
	}

	if cfg.Environment != config.ProductionEnv {
		errorRes["debug"] = err.Error()
	}

	c.JSON(status, Response{Error: errorRes})
}
