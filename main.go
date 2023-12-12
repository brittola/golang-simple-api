package main

import (
	"brittola-api/api/entities"
	"brittola-api/api/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	app := gin.Default()

	dsn := "root:root@tcp(127.0.0.1:3306)/my_twitter?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Tweet{})

	routes.AppRoutes(app, db)

	app.Run("localhost:3000")
}
