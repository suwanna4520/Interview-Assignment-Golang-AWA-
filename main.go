package main

import (
	"golang/controllers"
	"golang/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := models.InitialDatabase(); err != nil {
		panic(err)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:id", controllers.GetUserByID)

	e.Logger.Fatal(e.Start(":8080"))
}
