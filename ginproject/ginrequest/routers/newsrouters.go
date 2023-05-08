package routers

import (
	"ginrequest/controllers/news"

	"github.com/gin-gonic/gin"
)

func NewsRoutersInit(r *gin.Engine) {
	newsGroup := r.Group("/news")
	{
		//http://localhost:8080/news/home
		newsGroup.GET("/home", news.XNewsController{}.Home)
	}
}
