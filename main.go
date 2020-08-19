package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		NAME string `json; "name"`
	}
)

func main() {

	port := os.Getenv("PORT")

	e := echo.New()
	e.GET("/:string", printString)
	e.Logger.Fatal(e.Start(":" + port))
}

func printString(c echo.Context) error {
	data := c.Param("string")
	user := map[string]User{
		"1": User{
			NAME: "Hoge",
		},
		"2": User{
			NAME: data,
		},
	}
	return c.JSON(http.StatusOK, user)
}
