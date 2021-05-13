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
		json := model.User{}
		c.BindJSON(&json)
		fmt.Println(json.Username, json.Password)
		var user = model.User{
			Username: json.Username,
			Password: "*********",
			Avatar:   "https://azusachino.cn/img/avatar_hu665707f40d683dd23d0f4e3ce66154da_29874_300x0_resize_q75_box.jpg",
		}
		c.JSON(200, gin.H{
			"message": "OK",
			"data":    user,
		})
	})

	app.Run()
}
