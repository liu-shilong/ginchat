package cmd

import (
	"fmt"
	"ginchat/internal/middleware"
	"ginchat/internal/model"
	"ginchat/route"
	"ginchat/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"os"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "启动服务",
	Example: "./ginchat start",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap()
	},
}

// 启动
func bootstrap() {
	// 数据库连接
	model.Connect()
	// 运行模式 debug
	gin.SetMode(gin.DebugMode)

	engine := gin.Default()

	// TraceId中间件
	engine.Use(middleware.TraceMiddleware())

	// 日志中间件
	setupLogger, err := logger.SetupLogger()
	if err != nil {
		fmt.Printf("Failed to set up logger: %s", err)
		os.Exit(1)
	}
	engine.Use(middleware.RequestLoggerMiddleware(setupLogger))

	// 初始化路由
	route.InitApiRouter(engine)

	// 运行
	err = engine.Run(":8080")
	if err != nil {
		return
	}
}
