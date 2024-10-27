package routes

import (
	v1 "url_shortner/api"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	v1ApiGroup := r.Group("/v1")
	v1.Route(v1ApiGroup)

}
