package logging

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"

	"github.com/google/uuid"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Request ID хедерден алу немесе генерациялау
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Контекстке және жауапқа орнату
		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		// Жалғастыру
		c.Next()

		duration := time.Since(start)

		// Лог шығару
		log.Printf("[%s] [RequestID: %s] %s %s %d %s",
			start.Format(time.RFC3339),
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}
