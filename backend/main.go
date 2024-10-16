package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"korie/api-product/app"
	"korie/api-product/controller"
	"korie/api-product/helper"
	"korie/api-product/repository"
	"korie/api-product/service"
	"github.com/gin-contrib/cors"
	"time"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	ProductRepository := repository.NewProductRepository()
	productService := service.NewProductService(ProductRepository, db, validate)
	productController := controller.NewProductController(productService)

	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Change to your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))


	r.GET("api/products", productController.FindAll)       // GET ALL
	r.POST("api/products", productController.Create)       // CREATE
	r.PUT("api/products/:id", productController.Update)    // UPDATE
	r.DELETE("api/products/:id", productController.Delete) // DELETE
	r.GET("api/products/:id", productController.FindById)  // GET BY ID

	err := r.Run("localhost:3001")
	helper.IfErrNotNul(err)
}
