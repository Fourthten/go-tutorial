package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("App", "Simple app")
	// argAppName = kingpin.Arg("name", "Application name").Required().String()
	// argPort    = kingpin.Arg("port", "Port to listen on").Default("9000").Int()
	// argAppName = app.Arg("name", "Application name").Required().String()
	// argPort    = app.Arg("port", "Port to listen on").Default("9000").Int()
	argAppName = app.Flag("name", "Application name").Required().String()
	argPort    = app.Flag("port", "Port to listen on").Short('p').Default("9000").Int() // -p or --port
)

func main() {
	// kingpin.Parse()
	kingpin.MustParse(app.Parse(os.Args[1:]))

	appName := *argAppName
	port := fmt.Sprintf(":%d", *argPort)
	fmt.Printf("Starting %s on port %s\n", appName, port)
	e := echo.New()
	e.GET("/index", func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, true)
	})
	e.Logger.Fatal(e.Start(port))
}

// go run arg.go "My application" 9000
// go run arg.go --name "My application" -p 9000
