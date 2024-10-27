package handlers

import (
	controller "url_shortner/controller"

	"github.com/gin-gonic/gin"
)

func GetOrignalUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	orignalUrl := controller.GetOrignalUrl(shortUrl)
	ctx.Redirect(302, orignalUrl)
}
