package echo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Echo(c *gin.Context) {
	echo, _ := c.Params.Get("echo")
	c.String(http.StatusOK, echo)
}
