package exception

import (
	"github.com/gin-gonic/gin"
	"korie/api-product/model/web"
	"net/http"
)

func ErrorOrSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, web.WebResponse{
		Code:   statusCode,
		Status: http.StatusText(statusCode),
		Data:   data,
	})
	return
}
