package controllers

import (
	"github.com/David83656/go-api/initializers"
	"github.com/David83656/go-api/models"
	"github.com/gin-gonic/gin"
)

func CourseCreate(c *gin.Context) {

	var courseInfo struct {
		Description string
		Name        string
	}
	c.Bind(&courseInfo)

	course := models.Course{Name: "Java", Category: "Programming", Duration: 20, Description: "XDDDD"}

	result := initializers.DB.Create(&course)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Course": course,
	})
}
