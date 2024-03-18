package rand

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	AddRoutes(&r.RouterGroup)
	return r
}

// Notice that this test is focused around testing the HTTP contract and not the logic,
// Our sub function for producing the random number is more throughly tested and this is
// testing that the correct http request will result in a response, we don't actually test that
// the returned number is correct
func TestHandlerRandInt(t *testing.T) {
	router := setupRouter()

	lower := 5
	upper := 10

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/int/%d/%d", apiBasePath, lower, upper), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	returnedNumber, err := strconv.Atoi(strings.ReplaceAll(w.Body.String(), "\n", ""))
	if err != nil {
		t.Logf("returned string was: '%s'", w.Body.String())
		t.Fatal("not an integrer returned")
	}
	t.Logf("parsed number was: %d", returnedNumber)
}
