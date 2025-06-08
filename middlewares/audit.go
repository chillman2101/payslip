package middlewares

import (
	"bytes"
	"io/ioutil"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/payslip/models"
	"github.com/payslip/services"
)

func AuditMiddleware(s *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		var body []byte
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" || c.Request.Method == "DELETE" {
			bodyBytes, err := ioutil.ReadAll(c.Request.Body)
			if err == nil {
				body = bodyBytes
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		c.Next()
		user := c.GetUint("user_id")
		request_id := c.GetString("request_id")
		endpoint := c.Request.URL.Path
		method := c.Request.Method
		statusCode := c.Writer.Status()
		duration := time.Since(start)
		clientIP := c.ClientIP()

		log.Printf("[AUDIT] user=%d request_id=%s  method=%s endpoint=%s status=%d duration=%s body=%s\n",
			user, request_id, method, endpoint, statusCode, duration, string(body),
		)

		audit := models.AuditRequest{
			RequestId:  request_id,
			User:       user,
			Endpoint:   endpoint,
			Method:     method,
			StatusCode: statusCode,
			Duration:   duration,
			ClientIp:   clientIP,
		}

		s.AuditLog(c.Request.Context(), audit)
	}
}
