/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/12/26 10:12
*/
package routers

import (
	"encoding/gob"
	"github.com/gin-contrib/cors"
	"leeblog.com/app/http/controller/admins"
	"leeblog.com/app/http/middleware"
	"leeblog.com/app/models"
)

//import "leeblog.com/app/http/controller/admin"

// adminRouter 后台管理路由
func adminRouter() {
	gob.Register(models.Admin{})
	router := router()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Authorization"}
	router.Use(cors.New(config))
	//router.HTMLRender = loadTemplates("templates", "admin")
	router.Static("/layuiadmin", "./static/layuiadmin")
	//adminRouter := router.Group("admin")
	//adminRouter.Use(middleware.Session("admin"))
	//{
	//	adminRouter.GET("/login", admin.LoginIndex)
	//	adminRouter.GET("/captcha", admin.Captcha)
	//	adminRouter.GET("/captcha/:value", admin.CaptchaVerify)
	//	adminRouter.POST("/login", admin.LoginDoLogin)
	//}
	//adminAuth := adminRouter.Group("")
	//adminAuth.Use(middleware.AdminAuth())
	//{
	//	//人力管理
	//	adminAuth.GET("/user/index", admin.UserIndex)
	//}
	adminRouter := router.Group("admin")

	apiV1 := adminRouter.Group("v1")
	{

		apiV1.POST("/login", admins.LoginDoLogin)
		apiV1.GET("/admin/detail", admins.AdminDetail)
	}
	apiV1Auth := adminRouter.Group("v1")
	apiV1Auth.Use(middleware.JWT())
	{
		apiV1Auth.POST("/logout", admins.Logout)
		apiV1Auth.GET("/admin/info", admins.AdminIndex)
		//管理员列表
		apiV1Auth.GET("/admin/list", admins.AdminList)
		apiV1Auth.GET("/order/index", admins.OrderIndex)
		apiV1Auth.GET("/order/list", admins.OrderIndex)
		apiV1Auth.GET("/user/list", admins.UserList)
		apiV1Auth.GET("/role/list", admins.RoleList)
		//删除角色
		apiV1Auth.DELETE("/role/delete/:roleId", admins.RoleDelete)
		//添加角色
		apiV1Auth.POST("/role/add", admins.RoleAdd)
		apiV1Auth.GET("/role/permission/:roleId", admins.RolePermission)
		apiV1Auth.POST("/role/permission/:roleId", admins.RolePermissionUpdate)
	}
}
