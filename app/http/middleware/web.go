/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/12/26 16:11
*/
package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"leeblog.com/app/services/web"
)

func Web() gin.HandlerFunc {
	return func(c *gin.Context) {
		categorySlice := web.NewCategory().GetParentCategory()
		session := sessions.Default(c)
		user := session.Get("user")
		c.Set("common", gin.H{"nav": c.Request.URL.Path, "user": user})
		c.Set("categorySlice", categorySlice)
		c.Set("user", user)
		c.Next()
	}
}
