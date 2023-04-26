package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"ginrequest/routers"
)

// localhost:8080 -> 127.0.0.1:8080

func getRequest() {
	//http://localhost:8080/user/login?phone=131&code=1212

	fmt.Println("---getRequest--phone,code")
	r := gin.Default()
	routers.UserRoutersInit(r)
	routers.FileRoutersInit(r)

	r.Run()
}

func main() {
	getRequest()
}
