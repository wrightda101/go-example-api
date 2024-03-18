package rand

import "github.com/gin-gonic/gin"

func AddRoutes(router *gin.RouterGroup) {
	r := New()
	internalRouter := router.Group(apiBasePath)
	internalRouter.GET(apiPathRandInt, r.RandInt)
}
