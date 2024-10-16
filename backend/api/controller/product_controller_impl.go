package controller

import (
	"github.com/gin-gonic/gin"
	web2 "jamal/api/api/models/web"
	"jamal/api/api/service"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductControllerImpl {
	return &ProductControllerImpl{
		ProductService: productService, // Assign the service
	}
}

func (controller ProductControllerImpl) Create(c *gin.Context) {
	// DECODE DARI JSON KE STRUCT
	var createProduct web2.ProductCreate
	if err := c.ShouldBindJSON(&createProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": "Invalid input data",
		})
		return
	}

	productResponse := controller.ProductService.Create(c, createProduct)

	c.JSON(productResponse.Code, productResponse)
}

func (controller ProductControllerImpl) Delete(c *gin.Context) {
	// MENGAMBIL PARAMETER DI URL
	productIdParam := c.Param("id")
	id, err := strconv.Atoi(productIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": "Invalid Product ID",
		})
		return
	}

	productResponse := controller.ProductService.Delete(c, id)

	c.JSON(productResponse.Code, productResponse)
}

func (controller ProductControllerImpl) Update(c *gin.Context) {
	// MENGAMBIL PARAMETER DI URL
	productIdParam := c.Param("id")
	id, err := strconv.Atoi(productIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": "Invalid Product ID",
		})
		return
	}

	// DECODE DARI JSON KE STRUCT
	var updateProduct web2.ProductUpdate
	if err = c.ShouldBindJSON(&updateProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": "Invalid input data",
		})
		return
	}

	productResponse := controller.ProductService.Update(c, updateProduct, id)

	c.JSON(productResponse.Code, productResponse)
}

func (controller ProductControllerImpl) FindById(c *gin.Context) {
	// MENGAMBIL PARAMETER DI URL
	productIdParam := c.Param("id")
	id, err := strconv.Atoi(productIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": "Invalid product ID",
		})
		return
	}

	productResponse := controller.ProductService.FindById(c, id)

	c.JSON(productResponse.Code, productResponse)
}

func (controller ProductControllerImpl) FindAll(c *gin.Context) {

	productResponse := controller.ProductService.FindAll(c)
	c.JSON(productResponse.Code, productResponse)
}
