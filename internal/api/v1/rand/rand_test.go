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

func TestRandInt(t *testing.T) {
	router := setupRouter()

	lower := 5
	upper := 10

	for i := 0; i < 100; i++ {
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

		assert.Equal(t, true, returnedNumber >= lower)
		assert.Equal(t, true, returnedNumber <= upper)
	}
}
