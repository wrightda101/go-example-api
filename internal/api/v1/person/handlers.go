package person

import (
	"example-api/internal/models"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlerCreatePerson(c *gin.Context) {
	person, _ := convertJsonToPerson(c.Request.Body)
	id, _ := db.CreatePerson(person)
	resp := ResponseCreatePerson{
		ID:      id,
		Message: "created",
	}
	c.JSON(http.StatusOK, resp)
}

func HandlerGetPerson(c *gin.Context) {
	idString, _ := c.Params.Get("id")
	id, _ := strconv.ParseUint(idString, 10, 32)
	person, _ := db.GetPerson(uint(id))
	resp := ResponseGetPerson{
		Person:  person,
		Message: "got",
	}
	c.JSON(http.StatusOK, resp)
}

func convertJsonToPerson(input io.ReadCloser) (models.Person, error) {
	// throw away error, bad idea, should be handling this
	data, _ := io.ReadAll(input)
	return models.PersonFromJSON(data), nil
}
