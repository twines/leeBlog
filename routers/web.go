/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/12/26 10:12
*/
package routers

import (
	"encoding/gob"
	"leeblog.com/app/http/controller/web"
	"leeblog.com/app/http/middleware"
	"leeblog.com/app/models"
)

// webRouter 网站前台
func webRouter() {
	gob.Register(models.User{})
	router := router()
	webRouter := router.Group("")
	webRouter.StaticFile("/favicon.ico", "./static/favicon.ico")
	webRouter.StaticFile("/logo.jpg", "./static/logo.jpg")
	webRouter.Static("/css", "./static/css")
	webRouter.Static("/js", "./static/js")
	webRouter.Static("/bootstrap", "./static/bootstrap")
	webRouter.Static("/upload", "./static/upload")
	webRouter.Static("/ckplayer", "./static/ckplayer")
	webRouter.Static("/editormd", "./static/editor.md")
	webRouter.Static("/fontawesome", "./static/fontawesome")
	webRouter.Use(middleware.Session("web"), middleware.Web())
	webRouter.GET("/", web.Index)
	webRouter.GET("/category/:id", web.Index)
	webRouter.GET("/blog/detail/:id", web.BlogDetail)
	webRouter.GET("/login", web.LoginIndex)
	webRouter.POST("/login", web.LoginPost)
	webRouter.POST("/upload", web.Upload)
	webAuth := webRouter.Group("")
	webAuth.Use(middleware.WebAuth())
	{
		webAuth.GET("/logout", web.Logout)
		webAuth.GET("/blog/edit/:id", web.BlogEdit)
		webAuth.GET("/blog/add", web.BlogAdd)
		webAuth.POST("/blog/add", web.BlogDoAdd)
		webAuth.POST("/blog/update/:id", web.BlogUpdate)
	}

}
