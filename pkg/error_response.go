package pkg

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{
		Message: message,
	})
}
