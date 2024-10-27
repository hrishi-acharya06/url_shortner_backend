package api

import (
	handlers "url_shortner/handlers"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.RouterGroup) {
	r.POST("/create-short-url", handlers.CreateShortUrl)
	r.GET("/:shortUrl", handlers.GetOrignalUrl)
}
