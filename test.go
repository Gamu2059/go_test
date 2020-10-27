package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/hello", handle)
	i := e.Start(":" + os.Getenv("PORT"))
	e.Logger.Fatal(i)
}

func handle(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
