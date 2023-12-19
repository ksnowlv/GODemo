package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"ginrequest/global"
	"ginrequest/models/user"
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

// func (c XUserController) UserLogin(ctx *gin.Context) {
// 	ctx.SetCookie("userid", "123456", 3600, "/", "localhost", false, true)
// 	phone := ctx.Query("phone")
// 	code := ctx.Query("code")
// 	fmt.Printf("---UserLogin:phone:%s,code=%s", phone, code)
// 	handleResponseData(ctx, phone, code)
// }

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

// 返回所有用户
func (c XUserController) GetAllUser(ctx *gin.Context) {
	user := user.XUser{}
	users, err := user.GetAll(global.GMySQL)

	if err != nil {
		fmt.Println("---XUserController GetAllUser err:", err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": users,
		"count":  len(users),
	})
}

// 返回指定id用户
func (c XUserController) GetUserById(ctx *gin.Context) {

	var user user.XUser

	err := ctx.Bind(&user)

	if len(user.UserId) == 0 {
		fmt.Println("---XUserController GetUserById no userid")
	}

	selectUser, err := user.GetUserById(global.GMySQL)

	var res gin.H
	if err != nil {
		res = gin.H{
			"code":    http.StatusOK,
			"message": "请求成功",
			"result":  "没有查询到该用户",
			"count":   0,
		}
	} else {
		res = gin.H{
			"code":    http.StatusOK,
			"message": "请求成功",
			"result":  selectUser,
			"count":   1,
		}
	}

	ctx.JSON(http.StatusOK, res)
}

// 返回指定id用户
func (c XUserController) UserRegist(ctx *gin.Context) {

	var user user.XUser

	err := ctx.Bind(&user)

	if err != nil {
		fmt.Println("---XUserController UserRegist Bind err:", err)
	} else {
		fmt.Printf("---XUserController UserRegist user:%v", user)
	}

	_, err = user.GetUserByPhone(global.GMySQL)

	fmt.Printf("---XUserController UserRegist resUser:%v, err:%v", user, err)

	if err != nil {

		// V4 基于随机数
		u4 := uuid.New()
		fmt.Println(u4.String()) // a0d99f20-1dd1-459b-b516-dfeca4005203
		user.UserId = u4.String()

		insertId, err := user.AddNewUser(global.GMySQL)

		if err != nil {
			fmt.Println("---XUserController UserRegist AddNewUser err:", err)
		} else {
			fmt.Println("---XUserController UserRegist  AddNewUser suceess!,insertId:", insertId)

		}
		age := strconv.Itoa(user.Age)
		info := "userid:" + user.UserId + "-name:" + user.Name + "-age:" + age + "-phone:" + user.Phone

		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "恭喜您，账号已注册成功，请登陆!",
			"data":    info,
		})

	} else {
		resUser, err := user.GetUserByPhone(global.GMySQL)

		if err != nil {
			fmt.Println("---XUserController update err:", err)
		}

		age := strconv.Itoa(resUser.Age)
		info := "userid:" + resUser.UserId + "\nname:" + resUser.Name + "\nage:" + age + "\nphone:" + resUser.Phone

		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "已注册过账号，请登陆",
			"data":    info,
		})
	}
}

func (c XUserController) UserLogin(ctx *gin.Context) {
	var user user.XUser

	err := ctx.Bind(&user)

	if err != nil {
		fmt.Println("---XUserController Bind user err:", err)
	} else {
		fmt.Printf("---user:%v", user)
	}

	_, err = user.GetUserByPhone(global.GMySQL)

	if err != nil {

		fmt.Println("---XUserController UserLogin：查询不到该用户，开始准备创建用户--- err:", err)
		// V4 基于随机数
		u4 := uuid.New()
		fmt.Println(u4.String()) // a0d99f20-1dd1-459b-b516-dfeca4005203
		user.UserId = u4.String()
		insertId, err := user.AddNewUser(global.GMySQL)

		if err != nil {
			fmt.Println("---XUserController Add err:", err)
		} else {
			fmt.Println("---XUserController Add  suceess!,insertId:", insertId)

		}
		age := strconv.Itoa(user.Age)
		info := "userid:" + user.UserId + "\nname:" + user.Name + "\nage:" + age + "\nphone:" + user.Phone

		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "创建用户成功",
			"data":    info,
		})

	} else {

		age := strconv.Itoa(user.Age)
		info := "userid:" + user.UserId + "-name:" + user.Name + "-age:" + age + "-phone:" + user.Phone

		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "用户登陆成功!",
			"data":    info,
		})
	}
}

// // 返回指定id用户
// func (c XUserController) UserAdd(ctx *gin.Context) {

// 	var user user.XUser

// 	err := ctx.Bind(&user)

// 	if err != nil {
// 		fmt.Println("---XUserController UserAdd err:", err)
// 	} else {
// 		fmt.Printf("---user:%v", user)
// 	}

// 	_, err = user.Get(global.GMySQL)

// 	fmt.Printf("---resUser:%v, err:%v", user, err)

// 	if err != nil {
// 		fmt.Println("---UserAdd--- err:", err)
// 		Id, err := user.Add(global.GMySQL)

// 		if err != nil {
// 			fmt.Println("---XUserController Add err:", err)
// 		}

// 		fmt.Println(Id)

// 		age := strconv.Itoa(user.Age)
// 		info := "name:" + user.Name + "age:" + age + "phone:" + user.Phone

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"code":    http.StatusOK,
// 			"message": "添加成功",
// 			"data":    info,
// 		})

// 	} else {
// 		_, err := user.Update(global.GMySQL)

// 		if err != nil {
// 			fmt.Println("---XUserController update err:", err)
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{
// 			"code":    http.StatusOK,
// 			"message": "更新成功",
// 		})
// 	}

// }
