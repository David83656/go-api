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
	id := c.Param("id")
	var courses []models.Course
	initializers.DB.Find(&courses, id)

	c.JSON(200, gin.H{
		"courses": courses,
	})
}

func CourseUpdate(c *gin.Context) {
	id := c.Param("id")

	var courseName struct {
		Name string
	}
	c.Bind(&courseName)

	var courses []models.Course
	initializers.DB.Find(&courses, id)

	initializers.DB.Model(&courses).Updates(models.Course{
		Name: courseName.Name,
	})

	c.JSON(200, gin.H{
		"courses": courses,
	})
}

func CourseDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Course{}, id)

	c.Status(200)
}
