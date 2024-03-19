package person

import "example-api/internal/models"

type ResponseCreatePerson struct {
	Message      string `json:"message"`
	ErrorMessage string `json:"error_message"`
	ID           uint   `json:"id"`
}

type ResponseGetPerson struct {
	Message      string        `json:"message"`
	ErrorMessage string        `json:"error_message"`
	Person       models.Person `json:"person"`
}
