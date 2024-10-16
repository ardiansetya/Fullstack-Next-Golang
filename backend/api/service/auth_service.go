package service

import (
	"github.com/gin-gonic/gin"
	web2 "jamal/api/api/models/web"
)

type AuthService interface {
	Register(ctx *gin.Context, register web2.AuthRequestRegister) web2.WebResponse
	Login(ctx *gin.Context, login web2.AuthRequestLogin) web2.WebResponse
	Logout(ctx *gin.Context) web2.WebResponse
}
