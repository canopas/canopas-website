package main

import (
    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
)

func main() {
    r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./vue-frontend/dist", false)))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

    r.Run()
}
