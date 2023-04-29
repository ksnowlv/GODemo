package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("---test gin")

	r := gin.Default()
	//config router
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello gin",
			"code":    200,
			"data":    "",
		})
	})

	//start http server
	r.Run()
}
