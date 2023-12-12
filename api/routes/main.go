package routes

import (
	"brittola-api/api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AppRoutes(router *gin.Engine, db *gorm.DB) *gin.RouterGroup {

	tweetController := controllers.NewTweetController(db)

	api := router.Group("tweets")
	{
		api.GET("/", tweetController.FindAll)
		api.GET("/:id", tweetController.FindById)
		api.GET("/user/:user", tweetController.FindByUser)
		api.POST("/", tweetController.Create)
		api.PUT("/:id", tweetController.Update)
		api.DELETE("/:id", tweetController.Delete)
	}

	return api
}
