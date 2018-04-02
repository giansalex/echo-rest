package main

import (
	"net/http"

	"github.com/giansalex/echo-rest/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", index).Name = "Index"
	e.GET("/hello/:name", hello)

	api := e.Group("api/v1")
	api.GET("/users", users)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func index(c echo.Context) error {
	return c.String(http.StatusOK, "ECHO REST API")
}

func hello(c echo.Context) error {
	name := c.Param("name")

	return c.String(http.StatusOK, "Hello "+name)
}

func users(c echo.Context) error {
	list := []*model.User{
		&model.User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		},
		&model.User{
			Name:  "GianCarlos",
			Email: "giansalex@gmail.com",
		},
	}
	return c.JSON(http.StatusOK, list)
}
