package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	"ginrequest/routers"
)

// localhost:8080 -> 127.0.0.1:8080

func main() {

	gin.ForceConsoleColor()
	// 记录到文件。
	logFile, err := os.Create("gin.log")

	if err != nil {
		fmt.Println("log file create fail:", err)
		return
	}

	gin.DefaultWriter = io.MultiWriter(logFile)

	r := gin.Default()

	//定义路由日志的格式
	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	log.Printf("interface %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	//配置路由
	routers.UserRoutersInit(r)
	routers.FileRoutersInit(r)

	r.Run()
}
