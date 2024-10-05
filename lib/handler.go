package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerOk (c *gin.Context, message string, TotalInfo any, data any) {
	c.JSON(http.StatusOK, Message{
		Success: true,
		Message: message,
		ResultsInfo: TotalInfo,
		Results: data,
	})
}

func HandlerUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Message{
		Success: false,
		Message: message,
	})
}

func HandlerNotFound (c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Message{
		Success: false,
		Message: message,
	})
}

func HandlerBadRequest (c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Message{
		Success: false,
		Message: message,
	})
}

func HandlerMaxFile (c *gin.Context, message string) {
	c.JSON(http.StatusRequestEntityTooLarge, Message{
		Success: false,
		Message: message,
	})
}