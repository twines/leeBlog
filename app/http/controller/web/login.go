package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"leeblog.com/app/models"
	"leeblog.com/app/services/web"
	"leeblog.com/pkg/utils"
	"net/http"
)

func LoginIndex(c *gin.Context) {
	session := sessions.Default(c)
	err := session.Get("error")
	session.Delete("error")
	session.Save()
	utils.View(c, "/web/login/index.html", gin.H{"error": err})
}

func LoginPost(c *gin.Context) {
	session := sessions.Default(c)
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" || password == "" {
		session.Set("error", "用户名或者密码不可以为空")
		_ = session.Save()
		c.Redirect(http.StatusFound, "/login")
		return
	}
	var user = models.User{Name: name}
	user = web.NewUser().GetUser(user)

	if user.ID <= 0 {
		session.Set("error", "用户不存在")
		_ = session.Save()
		c.Redirect(http.StatusFound, "/login")
		return
	}
	if user.Password != utils.EncodeMD5(password) {
		session.Set("error", "用户名或者密码错误")
		_ = session.Save()
		c.Redirect(http.StatusFound, "/login")
		return
	}
	session.Set("user", user)
	_ = session.Save()
	c.Redirect(http.StatusFound, "/")
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
	c.Redirect(http.StatusFound, "/")
}
