package main

import (
	"github.com/David83656/go-api/initializers"
	"github.com/David83656/go-api/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Course{})
}
