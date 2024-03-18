package echo

import "github.com/gin-gonic/gin"

func AddRoutes(router *gin.RouterGroup) {
	internalRouter := router.Group(apiBasePath)
	internalRouter.GET(apiPathEcho, Echo)
}
