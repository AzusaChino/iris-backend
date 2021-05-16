package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"iris/router"
)

func main() {
	app := gin.Default()

	app.Use(cors.Default())

	router.ApplyRouter(app)

	app.Run(":7878")
}
