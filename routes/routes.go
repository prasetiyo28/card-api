package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/prasetiyo28/card-api/controllers"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, This is Echo !")
	})

	e.GET("/users", controllers.FetchAllUsers)
	e.POST("/users", controllers.StoreUsers)
	e.PUT("/users", controllers.UpdateUsers)
	return e

}
