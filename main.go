package main

import (
	"github.com/MasterBrian99/short-leaf/config"
	"github.com/MasterBrian99/short-leaf/controllers"
	"github.com/MasterBrian99/short-leaf/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.DatabaseSetup()
	r := gin.Default()
	router := r.Group("api")
	short := router.Group("short")
	{
		short.POST("/", controllers.CreateShortUrl)
	}
	user := router.Group("user")
	{
		user.POST("/register", controllers.CreateNewUser)
		user.POST("/login", controllers.Login)
	}
	// config.DB.AutoMigrate(models.User{}, models.ShortUrls{})

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
