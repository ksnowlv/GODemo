package routers

import (
	"github.com/gin-gonic/gin"

	"ginrequest/controllers/file"
)

func FileRoutersInit(r *gin.Engine) {
	fileGroup := r.Group("/file")
	{
		//单文件上传
		/*
			curl -X POST http://localhost:8080/file/upload \
			-F "file=@/Users/lvwei/Documents/1.txt" \
			-H "Content-Type: multipart/form-data"

			curl -X POST http://localhost:8080/file/upload \
			-F "file=@/Users/ksnowlv/Documents/1.txt" \
			-H "Content-Type: multipart/form-data"
		*/
		fileGroup.POST("/upload", file.XFileController{}.FileUpload)

		//多文件上传
		/*
		 curl -X POST http://localhost:8080/file/multifileupload \
		  -F "upload[]=@/Users/lvwei/Documents/1.txt" \
		  -F "upload[]=@/Users/lvwei/Documents/2.txt" \
		  -F "upload[]=@/Users/lvwei/Documents/3.txt" \
		  -H "Content-Type: multipart/form-data"

		  curl -X POST http://localhost:8080/file/multifileupload \
		  -F "upload[]=@/Users/ksnowlv/Documents/1.txt" \
		  -F "upload[]=@/Users/ksnowlv/Documents/2.txt" \
		  -F "upload[]=@/Users/ksnowlv/Documents/3.txt" \
		  -H "Content-Type: multipart/form-data"
		*/
		fileGroup.POST("/multifileupload", file.XFileController{}.MultiFileUpload)
	}
}
