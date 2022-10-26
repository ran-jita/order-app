package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"order-app/pkg"
	"os"
	"sync"
)

func main() {
	godotenv.Load()

	order_app_mysql, err := InitMysql()
	if err!=nil {
		fmt.Println("fail connect to mysql")
		fmt.Println(err)
		os.Exit(2)
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("Starting Paper Payment Sync HTTP Handler\n")
		pkg.InitOrderAppHttpHandler(order_app_mysql)
	}()

	wg.Wait()
}

func InitMysql() (*gorm.DB, error) {
	var connection *gorm.DB

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	var err error

	connection, err = gorm.Open("mysql", connectionString)
	if nil != err {
		fmt.Sprintf("Failed connected to database %s", connectionString)
	} else {
		fmt.Sprintf("Successfully connected to database %s", connectionString)
	}

	fmt.Println("Connection is created")
	return connection, nil
}
