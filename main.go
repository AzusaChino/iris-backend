package main

import (
	"fmt"
	"iris/model"
	"iris/router"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(cors.Default())

	router.ApplyRouter(app)

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	app.POST("/login", func(c *gin.Context) {
		var username = c.Param("username")
		var password = c.Param("password")
		fmt.Println(username, password)
		var user = model.User{
			Username: username,
			Password: password,
			Avatar:   "https://img.luozhiyun.com/20210110110753.png",
		}
		c.JSON(200, gin.H{
			"message": "OK",
			"data":    user,
		})
	})

	app.Run()
}
