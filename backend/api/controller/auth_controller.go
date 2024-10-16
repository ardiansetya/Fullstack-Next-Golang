package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}
