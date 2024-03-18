package api

import (
	"example-api/internal/api/v1/echo"
	"example-api/internal/api/v1/rand"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.RouterGroup) {
	v1 := router.Group("api/v1")
	// add all echo endpoints
	echo.AddRoutes(v1)
	// add all rand endpoints
	rand.AddRoutes(v1)
}
