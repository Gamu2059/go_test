package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/hello", handle)
	i := e.Start(":1223")
	e.Logger.Fatal(i)
}

func handle(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
