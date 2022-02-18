package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type responseError struct {
	Error string `json:"error" example:"message"`
}

type responseErrors struct {
	Errors []string `json:"error"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, responseError{msg})
}

func errorsResponse(c *gin.Context, code int, errors []validator.FieldError) {
	var response responseErrors
	for _, err := range errors {
		response.Errors = append(response.Errors, err.Error())
	}
	c.AbortWithStatusJSON(code, response)
}
