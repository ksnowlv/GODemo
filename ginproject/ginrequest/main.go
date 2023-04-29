package main

import (
	"github.com/gin-gonic/gin"

	"ginrequest/config"
	"ginrequest/logger"
	"ginrequest/routers"
)

// localhost:8080 -> 127.0.0.1:8080

func main() {
	// 1. 加载配置
	// 2. 配置日志
	// 3. 配置MySQL
	// 4. 配置Redis
	// 5. 配置路由
	// 6. 启动服务（优雅关机和平滑重启）

	config.InitConfig()
	logger.InitLog()

	r := gin.Default()

	r.Use(logger.GinLogger, logger.GinRecovery(true))
	//配置路由
	routers.UserRoutersInit(r)
	routers.FileRoutersInit(r)
	r.Run()
}
