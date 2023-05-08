package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type XNewsController struct {
}

func (newsController XNewsController) Home(ctx *gin.Context) {

	ctx.HTML(
		http.StatusOK,
		"index.tmpl",
		gin.H{
			"title": "POSTS",
		},
	)
}
