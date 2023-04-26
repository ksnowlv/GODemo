package routers

import (
	"github.com/gin-gonic/gin"

	"ginrequest/controllers/user"
)

func UserRoutersInit(r *gin.Engine) {
	userGroup := r.Group("/user")
	{

		userGroup.GET("/home", user.XUserController{}.Index)
		userGroup.GET("/cookie", user.XUserController{}.CookieTest)

		userGroup.GET("/login", user.XUserController{}.UserLogin)
		/*{
			"phone": "133",
			"code": "11"
		} raw数据请求*/
		userGroup.POST("/login_json", user.XUserController{}.UserLoginWithJson)

		//POST form表单提交
		userGroup.POST("/login_form", user.XUserController{}.UserLoginWithForm)
	}
}
