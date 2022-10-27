package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-app/pkg/model/mysql"

	"order-app/pkg/model"
	"order-app/pkg/usecase"
)

type CustomerController interface {
	GetCustomer(ctx *gin.Context)
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
	customerId := ctx.Param("customer_id")

	data, err := c.customerUsecase.GetCustomer(customerId)
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

	data, err := c.customerUsecase.InsetCustomer(request)
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
