package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/azusachino/iris-content/common"
	"net/http"
)

const Authorization = "Authorization"

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(Authorization)
		if token == "" {
			c.JSON(
				http.StatusOK,
				common.Error(-1, "没有登录，无法访问"),
			)
			c.Abort()
			return
		}
		c.Next()
	}
}
