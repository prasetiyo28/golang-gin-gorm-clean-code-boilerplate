package handlers

import (
	"net/http"
	"test1/entities"
	"test1/usecases"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase *usecases.UserUsecase
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input entities.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUsecase.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.UserUsecase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, user)
	}
	c.JSON(http.StatusOK, user)

}
