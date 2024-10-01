package main

import (
	"test1/infrastructure"
	"test1/interfaces/handlers"
	"test1/interfaces/repository"
	"test1/usecases"

	"test1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db := infrastructure.InitDB()

	// Initialize repositories
	userRepo := &repository.GormUserRepository{DB: db}
	productRepo := &repository.ProductRepository{}

	// Initialize usecases
	userUsecase := &usecases.UserUsecase{UserRepo: userRepo}
	productUsecase := &usecases.ProductUsecase{ProductRepo: productRepo}

	// Initialize handlers
	userHandler := &handlers.UserHandler{UserUsecase: userUsecase}
	productHandler := &handlers.ProductHandler{ProductUsecase: productUsecase}

	// Setup the router
	r := gin.Default()

	api := r.Group("/api")

	// Register route modules
	routes.UserRoutes(api, userHandler)
	routes.ProductRoutes(api, productHandler)

	// Run the server
	r.Run(":8080")
}
