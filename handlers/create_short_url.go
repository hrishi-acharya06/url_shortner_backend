package handlers

import (
	controller "url_shortner/controller"
	model "url_shortner/model"

	"github.com/gin-gonic/gin"
)

func CreateShortUrl(ctx *gin.Context) {
	var shortUrlRequest model.ShortUrlRequest
	if err := ctx.ShouldBindJSON(&shortUrlRequest); err != nil {
		panic("Invalid Request Body")
	}
	statusCode, shortUrl := controller.GenerateShortUrl(shortUrlRequest.Url)
	ctx.JSON(statusCode, gin.H{
		"orignal_url": shortUrlRequest.Url,
		"short_url":   shortUrl,
	})
}
