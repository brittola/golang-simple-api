package main

import (
	"brittola-api/api/entities"
	"brittola-api/api/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	app := gin.Default()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Tweet{})

	routes.AppRoutes(app, db)

	app.Run("localhost:3000")
}
