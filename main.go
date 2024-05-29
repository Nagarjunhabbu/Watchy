package main

import (
	"fmt"
	"time"
	"watchy/internal/controller"
	"watchy/internal/service"
	"watchy/internal/sql_data"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := getDb()

	// all initialisation
	storer := sql_data.NewWatchEventStore(db)
	service := service.NewWatchEventService(storer)
	ctrl := controller.NewWatchEventController(service)

	e.GET("/v1/event/:user_id", ctrl.GetWatchEvents) // endpoint to get specific user watch events
	e.POST("/v1/event", ctrl.CreateWatchEvent)       // endpoint to insert user watch event logs

	e.Logger.Fatal(e.Start(":8000"))
}

func getDb() *gorm.DB {
	var db *gorm.DB
	var err error

	// Wait for MySQL to become available
	for {
		db, err = gorm.Open(mysql.Open("root:root123@tcp(mysql:3306)/eventdb"), &gorm.Config{})
		if err == nil {
			break // MySQL is available, break the loop
		}

		fmt.Println("MySQL is not available yet. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	// MySQL is now available, start your application logic here
	fmt.Println("MySQL is now available. Starting the application...")
	return db
}
