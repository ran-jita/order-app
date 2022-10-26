package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController interface {
	Ping(ctx *gin.Context)
}

type pingController struct{}

func NewPingController() *authController {
	return &authController{}
}

func (c *authController) Ping(ctx *gin.Context) {
	result := map[string]interface{}{
		"message":     "ping",
		"status_code": 200,
	}

	ctx.JSON(http.StatusOK, result)
}
