package controller

import "github.com/gin-gonic/gin"

/**
 * URL控制器
 */
type Handler interface {
	Method() string
	Path() string
	HandlerFunc(ctx *gin.Context)
}
