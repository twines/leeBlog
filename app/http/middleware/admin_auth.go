/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/20 16:40
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/pkg/utils"
	"net/http"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if utils.AuthCheck(c, "admin") == nil {
			c.Redirect(http.StatusFound, "/admin/login")
		} else {
			c.Next()
		}
	}
}
