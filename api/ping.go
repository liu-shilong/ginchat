package api

import (
	"github.com/gin-gonic/gin"
	logrus "github.com/sirupsen/logrus"
	"net/http"
)

func Ping(c *gin.Context) {
	traceID := c.Request.Header.Get("X-Trace-ID")
	logrus.Info("124123123")
	c.JSON(http.StatusOK, gin.H{
		"trace_id": traceID,
	})
}
