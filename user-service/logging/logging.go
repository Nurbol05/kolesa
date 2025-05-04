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

		// RequestID генерациялау немесе хедерден алу
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Контекстке және жауапқа қосу (response header)
		c.Set("RequestID", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		// Хендлерді шақыру
		c.Next()

		duration := time.Since(start)

		// Логқа шығару
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
