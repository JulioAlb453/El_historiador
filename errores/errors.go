package errores

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SendErrorRespose(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, ErrorResponse{
		Status:  http.StatusText(statusCode),
		Message: err.Error(),
	})
}
