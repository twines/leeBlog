/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/20 16:40
*/
package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func WebAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if user := session.Get("user"); user == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
		}

		c.Next()
	}
}
