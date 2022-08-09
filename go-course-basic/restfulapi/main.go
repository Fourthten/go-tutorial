package main

import (
	"go-course-basic/restfulapi/app"
	"go-course-basic/restfulapi/controller"
	"go-course-basic/restfulapi/helper"
	"go-course-basic/restfulapi/middleware"
	"go-course-basic/restfulapi/repository"
	"go-course-basic/restfulapi/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
