package utils

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse format
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
