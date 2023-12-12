package controllers

import (
	"brittola-api/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type tweetController struct {
	db *gorm.DB
}

func NewTweetController(db *gorm.DB) *tweetController {
	return &tweetController{db: db}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	var tweets []entities.Tweet
	t.db.Find(&tweets)
	ctx.JSON(http.StatusOK, tweets)
}

func (t *tweetController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
	var tweet entities.Tweet

	if err := t.db.First(&tweet, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

	ctx.JSON(http.StatusOK, tweet)
}

func (t *tweetController) FindByUser(ctx *gin.Context) {
	user := ctx.Param("user")
	var userTweets []entities.Tweet

	if err := t.db.Where("user = ?", user).Find(&userTweets).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving tweets"})
		return
	}

	if len(userTweets) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "This user has no tweets"})
		return
	}

	ctx.JSON(http.StatusOK, userTweets)
}

func (t *tweetController) Create(ctx *gin.Context) {
	var tweet entities.Tweet

	if err := ctx.BindJSON(&tweet); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := t.db.Create(&tweet).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating tweet"})
		return
	}

	ctx.JSON(http.StatusCreated, tweet)
}

func (t *tweetController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var data entities.Tweet

	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if data.Description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Field 'description' cannot be empty"})
		return
	}

	var tweet entities.Tweet

	if err := t.db.First(&tweet, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

	tweet.Description = data.Description
	t.db.Save(&tweet)

	ctx.JSON(http.StatusOK, tweet)
}

func (t *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var tweet entities.Tweet

	if err := t.db.First(&tweet, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Tweet not found"})
		return
	}

	t.db.Delete(&tweet)

	ctx.JSON(http.StatusOK, gin.H{"message": "Tweet deleted"})
}
