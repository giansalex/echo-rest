package main

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/giansalex/echo-rest/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", index).Name = "Index"
	e.GET("/hello/:name", hello)
	e.POST("/api/login", login)

	api := e.Group("api/v1")
	api.Use(middleware.JWT([]byte("secret")))
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

func login(c echo.Context) (err error) {
	auth := new(model.Auth)
	if err = c.Bind(auth); err != nil {
		return
	}

	if auth.Username == "admin" && auth.Password == "123456" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Giancarlos"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token":  t,
			"expire": claims["exp"],
		})
	}

	return echo.ErrUnauthorized
}
