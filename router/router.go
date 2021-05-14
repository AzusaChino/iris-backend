package router

import (
	"iris/controller"

	"github.com/gin-gonic/gin"
)

func ApplyRouter(app *gin.Engine) {
	login := new(controller.LoginHandler)
	ac := new(controller.ArticleController)
	adc := new(controller.ArticleDetailController)
	cc := new(controller.CommentController)
	doApply(app, login, ac, adc, cc)
}

func doApply(app *gin.Engine, handlers ...controller.Handler) {
	for _, handler := range handlers {
		app.Handle(handler.Method(), handler.Path(), handler.HandlerFunc)
	}
}
