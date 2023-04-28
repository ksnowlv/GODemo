package config

import (
	"fmt"

	"ginrequest/core"
	"ginrequest/redisdb"
)

func InitConfig() {

	// 记录到文件。
	// logFile, err := os.Create("gin.log")

	// if err != nil {
	// 	fmt.Println("log file create fail:", err)

	// }
	// gin.DefaultWriter = io.MultiWriter(logFile)

	//定义路由日志的格式
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("interface %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	fmt.Println("---InitConfig---")

	core.InitViper()
	err := redisdb.InitRedisClient()

	if err != nil {
		fmt.Println("redis client error:", err)
	} else {
		fmt.Println("redis client success!")
	}
}
