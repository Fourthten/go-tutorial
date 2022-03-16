package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func main() {
	r := echo.New()

	// => curl -X GET http://localhost:9000
	r.GET("/", func(ctx echo.Context) error {
		data := "Hello World"
		return ctx.String(http.StatusOK, data)
	})

	// => curl -X GET http://localhost:9000/index
	r.GET("/index", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	// => curl -X GET http://localhost:9000/html
	r.GET("/html", func(ctx echo.Context) error {
		data := "Hello from /html"
		return ctx.String(http.StatusOK, data)
	})

	// => curl -I -X GET http://localhost:9000/index
	r.GET("/index", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	// => curl -X GET http://localhost:9000/json
	r.GET("/json", func(ctx echo.Context) error {
		data := M{"message": "Hello", "Counter": 2}
		return ctx.JSON(http.StatusOK, data)
	})

	// => curl -X GET http://localhost:9000/page1?name=setia
	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)
		return ctx.String(http.StatusOK, data)
	})

	// => curl -X GET http://localhost:9000/page2/ajung
	r.GET("/page2/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)
		return ctx.String(http.StatusOK, data)
	})

	// => curl -X GET http://localhost:9000/page3/setia/welcome
	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")
		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)
		return ctx.String(http.StatusOK, data)
	})

	// => curl -X POST -F name=tia -F message=welcome http://localhost:9000/page4
	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")
		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, strings.Replace(message, "/", "", 1))
		return ctx.String(http.StatusOK, data)
	})

	r.Start(":9000")
}
