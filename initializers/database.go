package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global Variables
var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=babar.db.elephantsql.com user=mxirbqxf password=bnNOFF5_Lp6Yl-JHWQ7yU9pM5ZNng7LM dbname=mxirbqxf port=5432 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
}
