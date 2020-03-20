/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/9 9:57
*/
package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("editormd-image-file")

	// 上传文件至指定目录
	if err := c.SaveUploadedFile(file, "./static/upload/"+file.Filename); err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"url": "/upload/" + file.Filename, "message": "上传成功", "success": 1})
}
