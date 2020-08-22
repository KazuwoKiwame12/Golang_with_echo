package main

import (
	"net/http"
	"os"
	"strconv"

	like "github.com/KazuwoKiwame12/test-echo-with-postgres/DB/Model/Like"
	user "github.com/KazuwoKiwame12/test-echo-with-postgres/DB/Model/User"
	"github.com/labstack/echo/v4"
)

type (
	User struct {
		ID   int    `json; "id"`
		NAME string `json; "name"`
	}

	Like struct {
		ID      int  `json; "id"`
		SelfID  int  `json; "self_id"`
		LoverID int  `json; "lover_id"`
		hasLove bool `json; "hasLove"`
	}
)

func main() {

	port := os.Getenv("PORT")

	e := echo.New()
	//UserモデルのAPI
	e.PUT("/User/update", updateUser)
	e.POST("/User/create", createUser)
	e.GET("/User/:id", getUser)
	//LikeモデルのAPI
	e.DELETE("/Like/delete", deleteLike)
	e.POST("/Like/create", createLike)
	e.GET("/Like/Self/:id", getLikes)

	//defaultのAPI
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

func deleteLike(c echo.Context) error {
	selfID, _ := strconv.Atoi(c.FormValue("self_id"))
	loverID, _ := strconv.Atoi(c.FormValue("lover_id"))
	like := like.Delete(selfID, loverID)

	return c.JSON(http.StatusOK, like)
}

func createLike(c echo.Context) error {
	selfID, _ := strconv.Atoi(c.FormValue("self_id"))
	loverID, _ := strconv.Atoi(c.FormValue("lover_id"))
	hasLove, _ := strconv.ParseBool(c.FormValue("hasLove"))
	like := like.Create(selfID, loverID, hasLove)

	return c.JSON(http.StatusOK, like)
}

func getLikes(c echo.Context) error {
	data := c.Param("id")
	id, _ := strconv.Atoi(data)
	like := like.Get(id)

	return c.JSON(http.StatusOK, like)
}
