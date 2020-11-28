package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/prasetiyo28/card-api/models"
)

func FetchAllUsers(c echo.Context) error {
	result, err := models.FetchAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func StoreUsers(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	result, err := models.StoreUsers(username, email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, result)
}

func UpdateUsers(c echo.Context) error {
	id := c.FormValue("id")
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateUsers(conv_id, username, email, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusOK, result)
}
