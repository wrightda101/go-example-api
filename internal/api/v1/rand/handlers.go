package rand

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *mystruct) HanlderRandInt(c *gin.Context) {
	lowerString, _ := c.Params.Get("lower")
	upperString, _ := c.Params.Get("upper")

	// TODO: throwing away the error for now, should do some error handling if non-int sent over
	lower, _ := strconv.Atoi(lowerString)
	upper, _ := strconv.Atoi(upperString)
	num := r.randInt(int64(lower), int64(upper))

	c.String(http.StatusOK, fmt.Sprintf("%d\n", num))
}

// func RandString(c *gin.Context) {
// 	echo, _ := c.Params.Get("length")
// 	c.String(http.StatusOK, echo)
// }

// func (r *mystruct) randString(length int) string {
// 	// set default length
// 	if length == 0 {
// 		length = 8
// 	}
// 	r.r.
// }
