package controllers

import (
	"github.com/David83656/go-api/initializers"
	"github.com/David83656/go-api/models"
	"github.com/gin-gonic/gin"
)

func CourseCreate(c *gin.Context) {

	var courseName struct {
		Name string
	}
	c.Bind(&courseName)

	course := models.Course{Name: courseName.Name, Category: "Programming", Duration: 20}

	result := initializers.DB.Create(&course)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Course": course,
	})
}

func CourseIndex(c *gin.Context) {

	var courses []models.Course
	initializers.DB.Find(&courses)

	c.JSON(200, gin.H{
		"Courses": courses,
	})
}

func CourseShow(c *gin.Context) {

}
