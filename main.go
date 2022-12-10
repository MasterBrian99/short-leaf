package main

import (
	"github.com/MasterBrian99/short-leaf/config"
	"github.com/MasterBrian99/short-leaf/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.DatabaseSetup()
	router := gin.Default()
	short := router.Group("short")
	{
		short.GET("/", controllers.CreateShortUrl)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
