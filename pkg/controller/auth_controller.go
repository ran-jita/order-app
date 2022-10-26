package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"order-app/pkg/model/dto"

	"order-app/pkg/model"
	"order-app/pkg/usecase"
)

type AuthController interface {
	GetLogin(ctx *gin.Context)
	PostLogin(ctx *gin.Context)
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

func (c *authController) GetLogin(ctx *gin.Context) {
	username := ctx.Param("username")

	data, err := c.authUsecase.GetLogin(username)
	if err != nil {
		statusCode := http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	result := map[string]interface{}{
		"data":        data,
		"status_code": 200,
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *authController) PostLogin(ctx *gin.Context) {
	var (
		request dto.Login
		err error
	)

	err = ctx.BindJSON(&request)
	if err != nil {
		statusCode := http.StatusBadRequest
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	data, err := c.authUsecase.PostLogin(request)
	if err != nil {
		statusCode := http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	result := map[string]interface{}{
		"data":        data,
		"status_code": http.StatusOK,
	}

	ctx.JSON(http.StatusOK, result)
}
