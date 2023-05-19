package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		reqUrl := c.Request.URL
		c.Next()
		latency := time.Since(t)
		status := c.Writer.Status()
		log.Printf("%s请求完成, 耗时 %d, 状态 %d\n", reqUrl, latency, status)
	}
}
