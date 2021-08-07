package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Respond(c *gin.Context, statusCode int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(statusCode, body)
		return
	}
	c.JSON(statusCode, body)
}

func RespondError(c *gin.Context, apiErr *ApplicationError) {
	if apiErr != nil {
		Respond(c, apiErr.StatusCode, apiErr)
		return
	}
	Respond(c, http.StatusInternalServerError, "failed server request")
}
