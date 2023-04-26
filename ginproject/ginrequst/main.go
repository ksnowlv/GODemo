package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// localhost:8080 -> 127.0.0.1:8080

func getRequest() {
	//http://localhost:8080/user/login?phone=131&code=1212

	fmt.Println("---getRequest--phone,code")
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/login", func(ctx *gin.Context) {
			phone := ctx.Query("phone")
			code := ctx.Query("code")
			handleResponseData(ctx, phone, code)
		})

		/*{
			"phone": "133",
			"code": "11"
		} raw数据请求*/
		userGroup.POST("/login_json", postParamsJsonHandle)

		//POST form表单提交
		userGroup.POST("/login_form", postParamsFormHandle)

	}

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

	fmt.Println("----postParamsJsonHandle---")
	bodyData, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println("---body/--- \r\n " + string(bodyData))

	var body map[string]string
	err := json.Unmarshal(bodyData, &body)

	if err != nil {
		fmt.Println("json error!", err)
	} else {
		fmt.Println("----json---", body)
	}
	phone := body["phone"]
	code := body["code"]

	handleResponseData(ctx, phone, code)
}

func postParamsFormHandle(ctx *gin.Context) {

	fmt.Println("----postParamsFormHandle---")
	// bodyData, _ := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("---body/--- \r\n " + string(bodyData))
	// ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))

	// fmt.Println(ctx.Request.Form)

	phone := ctx.PostForm("phone")
	code := ctx.PostForm("code")

	handleResponseData(ctx, phone, code)
}

func main() {
	getRequest()
}
