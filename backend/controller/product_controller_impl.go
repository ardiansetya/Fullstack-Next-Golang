package controller

import (
	"github.com/gin-gonic/gin"
	"korie/api-product/exception"
	"korie/api-product/helper"
	"korie/api-product/model/web"
	"korie/api-product/service"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller ProductControllerImpl) Create(c *gin.Context) {
	/// SAVE HASIL DECODE
	productCreateRequest := web.ProductCreateRequest{}
	err := c.ShouldBindJSON(&productCreateRequest) // DECODE JSON FROM REQ BODY
	helper.IfErrNotNul(err)

	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT
	ProductResponse := controller.ProductService.Create(c.Request.Context(), productCreateRequest)

	/// GIVE OUTPUT TO CLIENT
	exception.ErrorOrSuccess(c, http.StatusOK, ProductResponse)
}

func (controller ProductControllerImpl) Update(c *gin.Context) {
	/// SAVE HASIL DECODE
	productUpdateRequest := web.ProductUpdateRequest{}
	err := c.ShouldBindJSON(&productUpdateRequest) // DECODE JSON FROM REQ BODY
	helper.IfErrNotNul(err)

	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT )
	productId := c.Param("id") /// MENGAMBIL PARAM
	id, err := strconv.Atoi(productId)
	if err != nil {
		exception.ErrorOrSuccess(c, http.StatusBadRequest, "Invalid product ID")
		return
	}
	productUpdateRequest.Id = id

	///GIVE OUPUT KE CLIENT
	ProductResponse := controller.ProductService.Update(c.Request.Context(), productUpdateRequest)
	exception.ErrorOrSuccess(c, http.StatusOK, ProductResponse)

}

func (controller ProductControllerImpl) Delete(c *gin.Context) {
	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT JSON
	productId := c.Param("id") /// MENGAMBIL PARAM
	id, err := strconv.Atoi(productId)
	helper.IfErrNotNul(err)

	// RESPONSE SUKSES
	controller.ProductService.Delete(c.Request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	/// GIVE OUTPUT TO CLIENT
	c.JSON(200, webResponse)

}

func (controller ProductControllerImpl) FindById(c *gin.Context) {
	/// EXECUTE SERVICE AND IMPLEMENTATION FORMAT OUTPUT JSON
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		exception.ErrorOrSuccess(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	/// OUTPUT JIKA DATA TIDAK TI DEMUKAN DI DB
	productResponse := controller.ProductService.FindById(c.Request.Context(), id)
	if productResponse == (web.ProductResponse{}) {
		exception.ErrorOrSuccess(c, http.StatusNotFound, "Product not found")
		return
	}

	// OUTPUT KETIKA SUKSES
	exception.ErrorOrSuccess(c, http.StatusOK, productResponse)

}

func (controller ProductControllerImpl) FindAll(c *gin.Context) {
	productResponses := controller.ProductService.FindAll(c.Request.Context())

	/// GIVE OUTPUT TO CLIENT
	exception.ErrorOrSuccess(c, http.StatusOK, productResponses)
}
