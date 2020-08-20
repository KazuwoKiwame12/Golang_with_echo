package main

import (
	"net/http"
	"os"
	"strconv"

	user "github.com/KazuwoKiwame12/test-echo-with-postgres/DB/Model/User"
	"github.com/labstack/echo/v4"
)

type (
	User struct {
		ID   int    `json; "id"`
		NAME string `json; "name"`
	}
)

func main() {

	port := os.Getenv("PORT")

	e := echo.New()
	e.PUT("/User/create", updateUser)
	e.POST("/User/create", createUser)
	e.GET("/User/:id", getUser)
	e.GET("/", helloWorld)
	e.Logger.Fatal(e.Start(":" + port))
}

func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello World!!")
}

func updateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	newName := c.FormValue("newName")
	user := user.Update(id, newName)

	return c.JSON(http.StatusOK, user)
}

func createUser(c echo.Context) error {
	data := c.FormValue("name")
	user := user.Create(data)

	return c.JSON(http.StatusOK, user)
}

func getUser(c echo.Context) error {
	data := c.Param("id")
	id, _ := strconv.Atoi(data)
	user := user.Get(id)

	return c.JSON(http.StatusOK, user)
}
