package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-app/pkg/model/dto"

	"order-app/pkg/model"
	"order-app/pkg/usecase"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthController(
	authUsecase usecase.AuthUsecase,
) *authController {
	return &authController{
		authUsecase: authUsecase,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var (
		request    dto.LoginRequest
		statusCode int
		err        error
	)

	err = ctx.BindJSON(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	data, err := c.authUsecase.Login(request)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, data))
}

func (c *authController) Register(ctx *gin.Context) {
	var (
		request    dto.Register
		statusCode int
		err        error
	)

	err = ctx.BindJSON(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	data, err := c.authUsecase.Register(request)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, data))
}
