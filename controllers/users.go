package controllers

import (
	"fmt"
	"golang/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	UserModel := models.User{}
	if err := c.Bind(&UserModel); err != nil {
		log.Fatal(err)
	}

	fmt.Println("UserModel: ", UserModel)

	if err := models.Db.Create(&UserModel).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, map[string]interface{}{"success": false, "message": "Create User error"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}

func GetUserByID(c echo.Context) error {
	id := c.Param("id")

	var users []models.User
	if err := models.Db.Where("id = ?", id).Find(&users).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusNotFound, map[string]interface{}{"success": false, "message": "Get User by ID error"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": users})
}
