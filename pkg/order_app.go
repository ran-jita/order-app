package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"order-app/pkg/controller"
	"order-app/pkg/repository"
	"order-app/pkg/usecase"
	"os"
)

func InitOrderAppHttpHandler(order_app_mysql *gorm.DB) {
	router := gin.Default()
	pingController := controller.NewPingController()

	userRepository := repository.NewUserRepository(order_app_mysql)
	authUsecase := usecase.NewAuthUsecase(userRepository)
	authController := controller.NewAuthController(authUsecase)

	group := router.Group("/v1")
	group.GET("/ping", pingController.Ping)

	authGroup := group.Group("/auth")
	{
		authGroup.GET("/login/:username", authController.GetLogin)
		authGroup.POST("/login", authController.PostLogin)
	}

	serverString := fmt.Sprintf("%s:%s",
		os.Getenv("SERVER_ADDRESS"),
		os.Getenv("SERVER_PORT"),
	)

	router.Run(serverString)

}
