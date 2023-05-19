package main

import (
	"time"

	"github.com/azusachino/iris/pkg/orm"
	"github.com/azusachino/iris/pkg/setting"
	"github.com/azusachino/iris/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	setting.SetUp()
	orm.SetUp()
}

func main() {
	app := gin.Default()

	app.Static("/assets", "./assets")

	app.Use(corsHandler())
	// app.Use(middleware.Logger())

	router.ApplyRouter(app)

	app.Run(":" + setting.AppConfig.Port)
}

func corsHandler() gin.HandlerFunc {
	return cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true允许所有
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 1 * time.Hour,
	})
}
