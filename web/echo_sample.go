package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

type response struct {
	Message string
}

type user struct {
	Id   string
	Name string
}

func RunEchoServer(host string, port string) {
	e := echo.New()
	e.HideBanner = true

	// e.Use(middleware.Logger())
	e.GET("/", homeEcho)
	e.GET("/user", userEcho)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", host, port)))
}

func homeEcho(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Message: "Hello"})
}

func userEcho(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	u := user{
		Id:   id,
		Name: name,
	}
	return c.JSON(http.StatusOK, u)
}
