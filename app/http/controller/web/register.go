package web

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"leeblog.com/app/models"
	"leeblog.com/app/services/web"
	"leeblog.com/pkg/utils"
	"net/http"
)

func RegisterIndex(c *gin.Context) {
	session := sessions.Default(c)
	err := session.Get("error")
	name := session.Get("name")
	session.Delete("error")
	session.Delete("name")
	session.Save()
	utils.View(c, "/web/register/index.html", gin.H{"error": err, "name": name})
}
func RegisterPost(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")
	session := sessions.Default(c)
	if name == "" || password == "" {
		session.Set("error", "用户名或者密码不可以为空")
		session.Set("name", name)
		_ = session.Save()
		c.Redirect(http.StatusFound, "/register")
		return
	}
	if password != confirmPassword {
		session.Set("error", "两次输入密码不一致")
		session.Set("name", name)
		_ = session.Save()
		c.Redirect(http.StatusFound, "/register")
		return
	}
	user := models.User{}
	user.Name = name
	user.Password = utils.EncodeMD5(password)
	u := web.NewUser().GetUser(user)
	if u.ID > 0 {
		session.Set("error", "用户名已经存在，换个试试")
		session.Set("name", name)
		_ = session.Save()
		c.Redirect(http.StatusFound, "/register")
	} else {
		web.NewUser().AddUser(user)
		c.Redirect(http.StatusFound, "/login")
	}

}
