package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("---test gin")

	r := gin.Default()
	//config router
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello gin",
		})
	})

	//start http server
	r.Run()
}
