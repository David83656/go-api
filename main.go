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
	r.GET("/", controllers.CourseCreate)
	Init()
	r.Run()

}
