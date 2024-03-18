package echo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gotest.tools/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	AddRoutes(&r.RouterGroup)
	return r
}

func TestEcho(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", apiBasePath+"/static_string", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "static_string", w.Body.String())
}
