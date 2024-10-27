package main

import (
	routerGroup "url_shortner/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	Initialize()
	r := gin.Default()
	routerGroup.Route(r)
	r.Run(":8083")
}
