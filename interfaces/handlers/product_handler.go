package handlers

import (
	"net/http"
	"test1/entities"
	"test1/usecases"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUsecase *usecases.ProductUsecase
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input entities.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.ProductUsecase.CreateProduct(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) CreatOrder(c *gin.Context) {
	var input entities.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.ProductUsecase.CreateOrder(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {

	product, err := h.ProductUsecase.GetProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetOrder(c *gin.Context) {

	product, err := h.ProductUsecase.GetOrder()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	c.JSON(http.StatusCreated, product)
}
