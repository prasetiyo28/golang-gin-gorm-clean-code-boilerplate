package routes

import (
	"test1/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup, handler *handlers.ProductHandler) {
	users := router.Group("/product")
	{
		users.POST("/", handler.CreateProduct)
		users.GET("/", handler.GetProduct)
		users.POST("/order", handler.CreatOrder)
		users.GET("/order", handler.GetOrder)
	}
}
