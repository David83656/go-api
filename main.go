package main

import (
	"github.com/David83656/go-api/initializers"
	"github.com/gin-gonic/gin"
)

func Init() {
	initializers.LoadEnv()
}

func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	Init()
	r.Run()

}
