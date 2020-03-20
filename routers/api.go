package routers

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	v1 "leeblog.com/app/http/controller/api/v1"
	"leeblog.com/app/http/middleware"
	_ "leeblog.com/docs"
)

// api Api router
func apiRouter() {
	router := router()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiRouter := router.Group("/api/v1/")
	{
		apiRouter.GET("/", v1.Index)
		apiRouter.POST("/login", v1.LoginIndex)
		apiRouter.POST("/index/post", v1.IndexPost)
		apiRouter.POST("/register/post", v1.RegisterPost)
	}
	authApi := apiRouter.Group("")
	authApi.Use(middleware.JWT())
	{
		authApi.GET("/profile", v1.ProfileIndex)
	}
}
