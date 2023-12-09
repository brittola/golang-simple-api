package routes

import (
	"brittola-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {

	tweetController := controllers.NewTweetController()

	api := router.Group("")
	{
		api.GET("/tweets", tweetController.FindAll)
		api.GET("/tweets/:id", tweetController.FindById)
		api.GET("/tweets/user/:user", tweetController.FindByUser)
		api.POST("/tweets", tweetController.Create)
		api.PUT("/tweets/:id", tweetController.Update)
		api.DELETE("/tweets/:id", tweetController.Delete)
	}

	return api
}
