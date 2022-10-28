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

type CustomerController interface {
	GetCustomer(ctx *gin.Context)
	GetAllCustomer(ctx *gin.Context)
	InsertCustomer(ctx *gin.Context)
	UpdateCustomer(ctx *gin.Context)
	DeleteCustomer(ctx *gin.Context)
}

type customerController struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerController(
	customerUsecase usecase.CustomerUsecase,
) *customerController {
	return &customerController{
		customerUsecase: customerUsecase,
	}
}

func (c *customerController) GetCustomer(ctx *gin.Context) {
	var statusCode int
	customerId := ctx.Param("customer_id")

	data, err := c.customerUsecase.GetCustomer(customerId)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, data))
}

func (c *customerController) GetAllCustomer(ctx *gin.Context) {
	var (
		request    dto.GetCustomers
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

	data, countRecord, err := c.customerUsecase.GetAllCustomer(request)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode,
		map[string]interface{}{
			"customers":       data,
			"count_customers": countRecord,
		},
	))
}

func (c *customerController) InsertCustomer(ctx *gin.Context) {
	var (
		request    mysql.Customer
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

	data, err := c.customerUsecase.InsertCustomer(request)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, data))
}

func (c *customerController) UpdateCustomer(ctx *gin.Context) {
	var (
		request    mysql.Customer
		statusCode int
		err        error
	)

	err = ctx.BindJSON(&request)
	if err != nil {
		statusCode = http.StatusBadRequest
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	request.Id = ctx.Param("customer_id")
	if request.Id == "" {
		err = errors.New("undefined customer_id")
		statusCode = http.StatusBadRequest

		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	data, err := c.customerUsecase.UpdateCustomer(request)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusOK
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, data))
}

func (c *customerController) DeleteCustomer(ctx *gin.Context) {
	var (
		statusCode int
		err        error
	)

	customerId := ctx.Param("customer_id")
	if customerId == "" {
		err = errors.New("undefined customer_id")
		statusCode = http.StatusBadRequest

		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return

	}

	err = c.customerUsecase.DeleteCustomer(customerId)
	if err != nil {
		statusCode = http.StatusForbidden
		ctx.JSON(statusCode, model.ResponseError(statusCode, err))
		return
	}

	statusCode = http.StatusNoContent
	ctx.JSON(statusCode, model.ResponseSuccess(statusCode, nil))
}
