package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"ginrequest/redisdb"
)

const XCOOKIE_USERID = "userid"
const XREDIS_KEY_NAME = "name"

type XUserController struct {
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

func (c XUserController) Home(ctx *gin.Context) {
	ctx.SetCookie(XCOOKIE_USERID, "123456", 3600, "/", "localhost", false, true)
	redisdb.RedisSetString(XREDIS_KEY_NAME, "ksnowlv", 3600*24)
	redisdb.RedisSetMultiValues("username", "ksnow", "age", 10, "address", "北京海淀区")

	// ctx.String(http.StatusOK, "cookie OK")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "请求成功",
	})
	zap.L().Error("错误日志")
}

func (c XUserController) Cookie(ctx *gin.Context) {
	userid, _ := ctx.Cookie(XCOOKIE_USERID)
	name := redisdb.RedisGetString(XREDIS_KEY_NAME)
	values := redisdb.RedisGetMultiValues("username", "age", "address")

	fmt.Println("---redis多个values获取:", values)
	ctx.String(http.StatusOK, "cookie userid:"+userid+":name:"+name)
	zap.L().Debug("测试日志")
}

func (c XUserController) UserLogin(ctx *gin.Context) {
	ctx.SetCookie("userid", "123456", 3600, "/", "localhost", false, true)
	phone := ctx.Query("phone")
	code := ctx.Query("code")
	fmt.Printf("---UserLogin:phone:%s,code=%s", phone, code)
	handleResponseData(ctx, phone, code)
}

func (c XUserController) UserLoginWithJson(ctx *gin.Context) {

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

func (c XUserController) UserLoginWithForm(ctx *gin.Context) {

	fmt.Println("----postParamsFormHandle---")
	// bodyData, _ := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("---body/--- \r\n " + string(bodyData))
	// ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))

	// fmt.Println(ctx.Request.Form)

	phone := ctx.PostForm("phone")
	code := ctx.PostForm("code")

	handleResponseData(ctx, phone, code)
}
