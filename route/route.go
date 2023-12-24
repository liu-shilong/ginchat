package route

import (
	"ginchat/api"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(engine *gin.Engine) {
	// ping测
	engine.GET("/ping", api.Ping)
}
