package controllers

import (
	"brittola-api/api/entities"

	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, v := range t.tweets {
		if v.ID == id {
			ctx.JSON(http.StatusOK, t.tweets[i])
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})

}

func (t *tweetController) FindByUser(ctx *gin.Context) {
	user := ctx.Param("user")
	var userTweets []entities.Tweet

	for _, v := range t.tweets {
		if v.User == user {
			userTweets = append(userTweets, v)
		}
	}

	if len(userTweets) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "This user has no tweets",
		})
		return
	}
	ctx.JSON(http.StatusOK, userTweets)
}

func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}
	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusCreated, t.tweets)
}

func (t *tweetController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var data entities.Tweet

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao processar JSON",
		})
		return
	}

	if data.Description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Campo 'description' não pode ser vazio",
		})
		return
	}

	for i, v := range t.tweets {
		if v.ID == id {
			t.tweets[i].Description = data.Description
			ctx.JSON(http.StatusOK, t.tweets[i])
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet não encontrado",
	})
}

func (t *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, v := range t.tweets {
		if v.ID == id {
			t.tweets = append(t.tweets[0:i], t.tweets[i+1:]...)

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Tweet deletado.",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet não encontrado.",
	})

}
