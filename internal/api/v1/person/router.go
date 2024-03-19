package person

import (
	"example-api/internal/models"

	"github.com/gin-gonic/gin"
)

var db models.DBManager

func AddRoutes(router *gin.RouterGroup) {
	db = models.NewMustFromFile("./postgres.yaml")
	internalRouter := router.Group(apiBasePath)
	internalRouter.POST("/", HandlerCreatePerson)
	internalRouter.GET(apiGetPerson, HandlerGetPerson)
}
