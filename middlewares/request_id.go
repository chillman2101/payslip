package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// middleware/request_id.go
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.NewString()
		c.Set("request_id", reqID)
		c.Writer.Header().Set("X-Request-ID", reqID)
		c.Next()
	}
}
