package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"order-app/middleware"
	"order-app/pkg/controller"
	"order-app/pkg/repository"
	"order-app/pkg/usecase"
	"os"
)

func InitOrderAppHttpHandler(orderAppMysql *gorm.DB) {
	router := gin.Default()

	middleware := middleware.NewMiddleware()

	pingController := controller.NewPingController()

	userRepository := repository.NewUserRepository(orderAppMysql)
	authUsecase := usecase.NewAuthUsecase(userRepository)
	authController := controller.NewAuthController(authUsecase)

	customerRepository := repository.NewCustomerRepository(orderAppMysql)
	customerUsecase := usecase.NewCustomerUsecase(customerRepository)
	customerController := controller.NewCustomerController(customerUsecase)

	orderRepository := repository.NewOrderRepository(orderAppMysql)
	orderUsecase := usecase.NewOrderUsecase(customerUsecase, orderRepository)
	orderController := controller.NewOrderController(orderUsecase)

	group := router.Group("/v1")
	group.GET("/ping", pingController.Ping)

	JwtRoutes := group.Group("")
	JwtRoutes.Use(middleware.JwtAuth())

	authGroup := group.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
	}

	customerGroup := JwtRoutes.Group("/customer")
	{
		customerGroup.GET("/:customer_id", customerController.GetCustomer)
		customerGroup.POST("/all", customerController.GetAllCustomer) // search with pagination
		customerGroup.POST("", customerController.InsertCustomer)
		customerGroup.PUT("/:customer_id", customerController.UpdateCustomer)
		customerGroup.DELETE("/:customer_id", customerController.DeleteCustomer)
	}

	orderGroup := JwtRoutes.Group("/order")
	{
		orderGroup.GET("/:order_id", orderController.GetOrder)
		orderGroup.POST("/all", orderController.GetAllOrder) // search with pagination
		orderGroup.POST("", orderController.InsertOrder)
		orderGroup.PUT("/:order_id", orderController.UpdateOrder)
		orderGroup.DELETE("/:order_id", orderController.DeleteOrder)
	}

	serverString := fmt.Sprintf("%s:%s",
		os.Getenv("SERVER_ADDRESS"),
		os.Getenv("SERVER_PORT"),
	)

	router.Run(serverString)

}
