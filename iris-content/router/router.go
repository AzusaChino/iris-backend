package router

import (
	c "github.com/azusachino/iris-content/controller"

	"github.com/gin-gonic/gin"
)

func ApplyRouter(app *gin.Engine) {

	apiV1 := app.Group("/api/v1")
	{
		apiV1.POST("/login", c.Login)
		apiV1.GET("/articles", c.QueryArticle)
		apiV1.GET("/article/:id", c.QueryArticleDetail)
		apiV1.GET("/article/:id/comments", c.QueryComment)

		// todo 需要鉴权的接口
		// authorized := apiV1.Group("/auth", gin.BasicAuth(gin.Accounts{}))
		// authorized.GET("")
	}
}
