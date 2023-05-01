package routers

import (
	"github.com/gin-gonic/gin"

	"ginrequest/controllers/user"
)

func UserRoutersInit(r *gin.Engine) {
	userGroup := r.Group("/user")
	{

		/*接口请求
		curl -X GET http://localhost:8080/user/home \
		 -H 'Content-Type: application/json' \
		 -H 'Accept-Encoding: gzip, deflate, br' \
		 -H 'Connection: keep-alive' \
		 -H 'User-Agent: ksnowlv关于Gin框架学习实践' \
		*/
		userGroup.GET("/home", user.XUserController{}.Home)
		/*接口请求
		curl -X GET http://localhost:8080/user/cookie \
		 -H 'Content-Type: application/json' \
		 -H 'Accept-Encoding: gzip, deflate, br' \
		 -H 'Connection: keep-alive' \
		 -H 'User-Agent: ksnowlv关于Gin框架学习实践'
		*/
		userGroup.GET("/cookie", user.XUserController{}.Cookie)

		/*Get接口请求,-G不通去掉，去掉就变成post请求了
				  curl  -G http://localhost:8080/user/login \
		 		  --data-urlencode 'phone=1' \
		 		  --data-urlencode 'code=1' \
				  -H 'Content-Type: application/json' \
				  -H 'Accept-Encoding: gzip, deflate, br' \
				  -H 'Connection: keep-alive' \
				  -H 'User-Agent: ksnowlv关于Gin框架学习实践'
		*/
		userGroup.GET("/login", user.XUserController{}.UserLogin)
		/* -d/data HTTP POST方式传送数据   -D/–dump-header 把header信息写入到该文件中
				 curl -X POST http://localhost:8080/user/login_json \
		 		  -d '{"phone": "12", "code": "12"}' \
				  -H 'Content-Type: application/json' \
				  -H 'Accept-Encoding: gzip, deflate, br' \
				  -H 'Connection: keep-alive' \
				  -H 'User-Agent: ksnowlv关于Gin框架学习实践'
		*/
		userGroup.POST("/login_json", user.XUserController{}.UserLoginWithJson)

		//POST form表单提交 curl中使用--data,而不是-data；--data如果有多个键值对，必须加上双引号
		/*
			curl --data "phone=152&code=152" "http://localhost:8080/user/login_form"

			curl --data "phone=152&code=152" "http://localhost:8080/user/login_form" > tmp.gz ; gzip -d tmp.gz & cat tmp
		*/

		userGroup.POST("/login_form", user.XUserController{}.UserLoginWithForm)
		/* -d/data HTTP POST方式传送数据   -D/–dump-header 把header信息写入到该文件中
				 curl -X POST http://localhost:8080/user/add \
		 		  -d '{"name": "ksnowlv", "age": 30,"Phone": "152"}' \
				  -H 'Content-Type: application/json' \
				  -H 'Accept-Encoding: gzip, deflate, br' \
				  -H 'Connection: keep-alive' \
				  -H 'User-Agent: ksnowlv关于Gin框架学习实践'

				  curl -X POST http://localhost:8080/user/add \
		 		  -d '{"name": "kair", "age": 31,"Phone": "151"}' \
				  -H 'Content-Type: application/json' \
				  -H 'Accept-Encoding: gzip, deflate, br' \
				  -H 'Connection: keep-alive' \
				  -H 'User-Agent: ksnowlv关于Gin框架学习实践'
		*/
		userGroup.POST("/add", user.XUserController{}.UserAdd)

		/* -d/data HTTP POST方式传送数据   -D/–dump-header 把header信息写入到该文件中
				 curl -X POST http://localhost:8080/user/getalluser \
		 		  -d '{"token": "123"}' \
				  -H 'Content-Type: application/json' \
				  -H 'Accept-Encoding: gzip, deflate, br' \
				  -H 'Connection: keep-alive' \
				  -H 'User-Agent: ksnowlv关于Gin框架学习实践' \
				  -H 'Accept: application/json, text/html'  > tmp.gz ; gzip -d tmp.gz & cat tmp
		*/
		userGroup.POST("/getalluser", user.XUserController{}.GetAllUser)

		/* -d/data HTTP POST方式传送数据   -D/–dump-header 把header信息写入到该文件中
		curl -X GET http://localhost:8080/user/getallusers \
		 -H 'Content-Type: application/json' \
		 -H 'Accept-Encoding: gzip, deflate, br' \
		 -H 'Connection: keep-alive' \
		 -H 'User-Agent: ksnowlv关于Gin框架学习实践' > tmp.gz ; gzip -d tmp.gz & cat tmp

		 curl -X GET http://localhost:8080/user/getallusers \
		 -H 'Content-Type: application/json' \
		 -H 'Accept-Encoding: gzip, deflate, br' \
		 -H 'Connection: keep-alive' \
		 -H 'User-Agent: ksnowlv关于Gin框架学习实践' --output tmp.gz;cat tmp
		*/
		userGroup.GET("/getallusers", user.XUserController{}.GetAllUser)
	}
}
