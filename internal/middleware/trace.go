package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头或其他地方获取现有的 Trace ID（如果存在）
		existingTraceID := c.Request.Header.Get("X-Trace-ID")

		// 如果不存在现有的 Trace ID，则生成一个新的 Trace ID
		if existingTraceID == "" {
			traceID := uuid.New().String()

			// 将新的 Trace ID 添加到请求头中
			c.Request.Header.Set("X-Trace-ID", traceID)
		}

		// 继续处理后续中间件和请求处理程序
		c.Next()
	}
}
