package api

import (
	"github.com/gin-gonic/gin"
)

func Handel(r *gin.Engine) *gin.Engine {
	rg := r.Group("/wchat")
	rg.GET("/verification", Validate)

	return r
}
