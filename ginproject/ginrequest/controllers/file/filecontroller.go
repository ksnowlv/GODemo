package file

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const FILEUPLOADPATH = "./fileres/"

type XFileController struct {
}

func (c XFileController) FileUpload(ctx *gin.Context) {

	fmt.Println("----FileUpload---")
	// bodyData, _ := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("---body/--- \r\n " + string(bodyData))

	file, err := ctx.FormFile("file")
	fmt.Println(file.Filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	dstFilePath := FILEUPLOADPATH + file.Filename
	// 上传文件至指定的完整文件路径
	ctx.SaveUploadedFile(file, dstFilePath)
	fmt.Println("----", dstFilePath)

	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func (c XFileController) MultiFileUpload(ctx *gin.Context) {

	// fmt.Println("----MultiFileUpload---")
	// bodyData, _ := ioutil.ReadAll(ctx.Request.Body)
	// fmt.Println("---body/--- \r\n " + string(bodyData))

	// Multipart form
	form, err := ctx.MultipartForm()

	if err != nil {
		fmt.Println("---MultiFileUpload---", err)
		return
	}

	files := form.File["upload[]"]
	fmt.Println("----MultiFileUpload---", files)

	for _, file := range files {
		dstFilePath := FILEUPLOADPATH + file.Filename

		// 上传文件至指定目录
		ctx.SaveUploadedFile(file, dstFilePath)
		fmt.Println("----dstFilePath---", dstFilePath)
	}
	ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
