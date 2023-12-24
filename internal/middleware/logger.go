package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RequestLoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始记录请求日志

		logger.WithFields(logrus.Fields{
			"path":       c.Request.URL.Path,
			"method":     c.Request.Method,
			"ip":         c.ClientIP(),
			"status":     c.Writer.Status(),
			"trace_id":   c.Request.Header.Get("X-Trace-ID"),
			"user_agent": c.Request.UserAgent(),
		}).Info("Request received")

		// 继续处理后续中间件和请求处理程序
		c.Next()

		// 记录响应日志
		logger.WithFields(logrus.Fields{
			"path":     c.Request.URL.Path,
			"method":   c.Request.Method,
			"status":   c.Writer.Status(),
			"size":     c.Writer.Size(),
			"trace_id": c.Request.Header.Get("X-Trace-ID"),
		}).Info("Response sent")
	}
}
