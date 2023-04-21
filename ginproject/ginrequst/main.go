package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRequest() {
	//http://localhost:8080/login?phone=ksnow&code=1212

	fmt.Println("---getRequest--phone,code")
	r := gin.Default()

	r.GET("/login", func(ctx *gin.Context) {
		phone := ctx.Query("phone")
		code := ctx.Query("code")
		handleResponseData(ctx, phone, code)
	})

	r.POST("user/login", postParamsJsonHandle)

	r.Run()
}

func handleResponseData(ctx *gin.Context, phone string, code string) {
	if phone == "" {
		fmt.Println("phone 为空!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "手机号为空!",
		})
		return
	}

	if code == "" {
		fmt.Println("为空!")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "验证码为空!",
		})
		return
	}

	fmt.Printf("phone:%s, code:%s", phone, code)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "请求成功!",
	})
}

func postParamsJsonHandle(ctx *gin.Context) {

	data, err := ctx.GetRawData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	var body map[string]string
	err = json.Unmarshal(data, &body)

	if err != nil {
		fmt.Println(err)
	}

	phone := body["phone"]
	code := body["code"]
	fmt.Printf("phone:%s,code:%s\n", phone, code)

	json := make(map[string]interface{}) //注意该结构接受的内容
	ctx.BindJSON(&json)
	fmt.Printf("%v\n", &json)

	handleResponseData(ctx, phone, code)
}

func postRequest() {
	//http://localhost:8080/user/login   phone=ksnow&code=1212
	r := gin.Default()

	r.POST("user/login", postParamsJsonHandle)
	r.Run()
}

func formRequest() {

}

func main() {
	getRequest()
	//postRequest()
	formRequest()
}
