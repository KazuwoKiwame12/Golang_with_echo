package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	port := os.Getenv("PORT")

	e := echo.New()
	e.GET("/:string", printString)
	e.Logger.Fatal(e.Start(":" + port))
}

func printString(c echo.Context) error {
	data := c.Param("string")
	return c.String(http.StatusOK, data)
}
