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
		customerGroup.POST("", customerController.InsertCustomer)
		customerGroup.PUT("/:customer_id", customerController.UpdateCustomer)
		customerGroup.DELETE("/:customer_id", customerController.DeleteCustomer)
	}

	serverString := fmt.Sprintf("%s:%s",
		os.Getenv("SERVER_ADDRESS"),
		os.Getenv("SERVER_PORT"),
	)

	router.Run(serverString)

}
