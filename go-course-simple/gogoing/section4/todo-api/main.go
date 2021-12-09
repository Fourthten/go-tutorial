package main

import (
	"fmt"
	"log"
	"net/http"

	// "os"

	"go-course-simple/gogoing/section4/todo-api/controller"
	"go-course-simple/gogoing/section4/todo-api/model"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
	// port := os.Getenv("PORT")
	// log.Fatal(http.ListenAndServe(":"+port, mux))
}
