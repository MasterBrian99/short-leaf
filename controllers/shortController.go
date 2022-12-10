package controllers

import "github.com/gin-gonic/gin"

type ShortInput struct {
	url string
}

func CreateShortUrl(c *gin.Context) {
	c.JSON(200, gin.H{
		"asd": "sd",
	})
}
