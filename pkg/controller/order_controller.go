package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"order-app/pkg/model/dto"
	"order-app/pkg/model/mysql"

	"order-app/pkg/model"
	"order-app/pkg/usecase"
)

type OrderController interface {
	GetOrder(ctx *gin.Context)
	GetAllOrder(ctx *gin.Context)
	InsertOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type orderController struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderController(
	orderUsecase usecase.OrderUsecase,
) *orderController {
	return &orderController{
		orderUsecase: orderUsecase,
	}
}

func (c *orderController) GetOrder(ctx *gin.Context) {
	var statusCode int
	orderId := ctx.Param("order_id")

	data, err := c.orderUsecase.GetOrder(orderId)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, data))
}

func (c *orderController) GetAllOrder(ctx *gin.Context) {
	var (
		request    dto.GetOrders
		statusCode int
		err        error
	)

	err = ctx.BindJSON(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	userInfo, exist := ctx.Get("userInfo")
	if !exist {
		err = errors.New("Unauthorized")
		statusCode = http.StatusUnauthorized
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}
	userInfoParsed := userInfo.(jwt.MapClaims)
	request.UserId = fmt.Sprintf("%s", userInfoParsed["UserId"])

	data, countRecord, err := c.orderUsecase.GetAllOrder(request)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode,
		map[string]interface{}{
			"orders":      data,
			"countorders": countRecord,
		},
	))
}

func (c *orderController) InsertOrder(ctx *gin.Context) {
	var (
		request    mysql.Order
		statusCode int
		err        error
	)

	err = ctx.BindJSON(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	userInfo, exist := ctx.Get("userInfo")
	if !exist {
		err = errors.New("Unauthorized")
		statusCode = http.StatusUnauthorized
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	userInfoParsed := userInfo.(jwt.MapClaims)
	request.UserId = fmt.Sprintf("%s", userInfoParsed["UserId"])

	data, err := c.orderUsecase.InsertOrder(request)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, data))
}

func (c *orderController) UpdateOrder(ctx *gin.Context) {
	var (
		request    mysql.Order
		statusCode int
		err        error
	)

	err = ctx.BindJSON(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	request.Id = ctx.Param("order_id")
	if request.Id == "" {
		err = errors.New("undefined order_id")
		statusCode = http.StatusBadRequest

		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	data, err := c.orderUsecase.UpdateOrder(request)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, data))
}

func (c *orderController) DeleteOrder(ctx *gin.Context) {
	var (
		statusCode int
		err        error
	)

	orderId := ctx.Param("order_id")
	if orderId == "" {
		err = errors.New("undefined order_id")
		statusCode = http.StatusBadRequest

		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return

	}

	err = c.orderUsecase.DeleteOrder(orderId)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusNoContent
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, nil))
}
