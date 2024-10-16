package service

import (
	"github.com/gin-gonic/gin"
	web2 "jamal/api/api/models/web"
)

type ProductService interface {
	Create(ctx *gin.Context, createReq web2.ProductCreate) web2.WebResponse
	Delete(ctx *gin.Context, productId int) web2.WebResponse
	Update(ctx *gin.Context, updateReq web2.ProductUpdate, productId int) web2.WebResponse
	FindById(ctx *gin.Context, productId int) web2.WebResponse
	FindAll(ctx *gin.Context) web2.WebResponse
}
