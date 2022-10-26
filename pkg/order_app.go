package pkg

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"order-app/pkg/controller"
	"os"
)

func InitOrderAppHttpHandler(order_app_mysql *gorm.DB) {
	router := gin.Default()
	pingController := controller.NewPingController()

	group := router.Group("/v1")
	group.GET("/ping", pingController.Ping)

	serverString := fmt.Sprintf("%s:%s",
		os.Getenv("SERVER_ADDRESS"),
		os.Getenv("SERVER_PORT"),
	)

	router.Run(serverString)

}
