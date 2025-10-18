package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var LoggerKey = "logrusLogger"

func LoggingMiddleware(baseLogger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		entry := baseLogger.WithFields(logrus.Fields{
			"request_id": c.Request.Header.Get("X-Request-ID"),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
		})
		c.Set(LoggerKey, entry)
		c.Next()
	}
}
