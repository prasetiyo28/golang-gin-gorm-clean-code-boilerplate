package routes

import (
	"test1/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, handler *handlers.UserHandler) {
	users := router.Group("/users")
	{
		users.POST("/", handler.CreateUser)
		users.GET("/:id", handler.GetUser)
	}
}
