package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// TODO generate middleware implement function, delete after code implementation

	// Passthrough to next handler if need
	c.Next()
}
