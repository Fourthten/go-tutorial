package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	r := echo.New()
	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})

	fmt.Println("Server started at http://localhost:9000")
	r.Start(":9000")
}

// json payload:
// curl -X POST http://localhost:9000/user -H 'Content-Type: application/json' -d '{"name":"Setia","email":"test@mail.com"}'

// xml payload:
// curl -X POST http://localhost:9000/user -H 'Content-Type: application/xml' -d '<?xml version="1.0"?><Data><Name>Setia</Name><Email>test@mail.com</Email></Data>'

// form data:
// curl -X POST http://localhost:9000/user -d 'name=Setia' -d 'email=test@mail.com'

// query string:
// curl -X GET http://localhost:9000/user?name=Setia&email=test@mail.com
