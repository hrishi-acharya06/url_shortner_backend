package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	Initialize()
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "This is go url shortner app"})
	})
	r.Run(":8083")
}
