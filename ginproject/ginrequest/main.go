package main

import (
	"github.com/gin-gonic/gin"

	"ginrequest/config"
	"ginrequest/routers"
)

// localhost:8080 -> 127.0.0.1:8080

func main() {

	gin.ForceConsoleColor()
	config.InitConfig()
	r := gin.Default()
	//配置路由
	routers.UserRoutersInit(r)
	routers.FileRoutersInit(r)
	r.Run()
}
