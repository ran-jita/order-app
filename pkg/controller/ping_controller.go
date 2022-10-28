package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-app/pkg/model"
)

type PingController interface {
	Ping(ctx *gin.Context)
}

type pingController struct{}

func NewPingController() *authController {
	return &authController{}
}

func (c *authController) Ping(ctx *gin.Context) {
	statusCode := http.StatusOK

	ctx.JSON(statusCode, model.ResponseSuccess(
		statusCode,
		"ping",
	))
}
