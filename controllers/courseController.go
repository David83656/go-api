package controllers

import "github.com/gin-gonic/gin"

func CourseCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
