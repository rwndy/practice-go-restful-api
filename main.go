package main

import (
	"Users/riwandi/Documents/practice/go-restful-api/app/database"
	"Users/riwandi/Documents/practice/go-restful-api/controller"
	"Users/riwandi/Documents/practice/go-restful-api/helper"
	"Users/riwandi/Documents/practice/go-restful-api/repository"
	"Users/riwandi/Documents/practice/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db := database.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	router.POST("/api/categories", categoryController.Create)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.HandlePanic(err)
}
