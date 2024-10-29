package main

import (
	"os"
	routerGroup "url_shortner/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	Initialize()
	r := gin.Default()
	routerGroup.Route(r)
	r.Run(":" + os.Getenv("server_port"))
}
