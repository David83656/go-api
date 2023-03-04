package main

import (
	"github.com/David83656/go-api/controllers"
	"github.com/David83656/go-api/initializers"
	"github.com/gin-gonic/gin"
)

func Init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/Courses", controllers.CourseCreate)
	r.PUT("/Courses/:id", controllers.CourseUpdate)
	r.GET("/Courses", controllers.CourseIndex)
	r.GET("/Courses/:id", controllers.CourseShow)
	r.DELETE("/Courses/:id", controllers.CourseDelete)
	Init()
	r.Run()

}
