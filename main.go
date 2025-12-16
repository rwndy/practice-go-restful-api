package main

import (
	"fmt"
	"github.com/rwndy/practice-go-restful-api/app/database"
	"github.com/rwndy/practice-go-restful-api/controller"
	"github.com/rwndy/practice-go-restful-api/repository"
	"github.com/rwndy/practice-go-restful-api/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using defaults")
	}

	db := database.NewDB()
	defer db.Close()

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
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("🚀 Server starting on http://localhost:3000")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
