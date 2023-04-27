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

	// 记录到文件。
	log, err := os.Create("gin.log")

	if err != nil {
		fmt.Println("log file create fail:", err)
		return
	}

	gin.DefaultWriter = io.MultiWriter(log)

	r := gin.Default()
	routers.UserRoutersInit(r)
	routers.FileRoutersInit(r)

	r.Run()
}
