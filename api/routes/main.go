package routes

import (
	"brittola-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {

	tweetController := controllers.NewTweetController()

	v1 := router.Group("")
	{
		v1.GET("/tweets", tweetController.FindAll)
		v1.POST("/tweets", tweetController.Create)
		v1.PUT("/tweets/:id", tweetController.Update)
		v1.DELETE("/tweets/:id", tweetController.Delete)
	}

	return v1
}
