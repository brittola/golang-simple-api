package routes

import (
	"brittola-api/api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AppRoutes(router *gin.Engine, db *gorm.DB) *gin.RouterGroup {

	tweetController := controllers.NewTweetController(db)
	userController := controllers.NewUserController(db)

	api := router.Group("")
	{
		tweetsAPI := api.Group("tweets")
		{
			tweetsAPI.GET("/", tweetController.FindAll)
			tweetsAPI.GET("/:id", tweetController.FindById)
			tweetsAPI.GET("/user/:user", tweetController.FindByUser)
			tweetsAPI.POST("/", tweetController.Create)
			tweetsAPI.PUT("/:id", tweetController.Update)
			tweetsAPI.DELETE("/:id", tweetController.Delete)
		}

		usersAPI := api.Group("users")
		{
			usersAPI.POST("/register", userController.Create)
		}
	}

	return api
}
