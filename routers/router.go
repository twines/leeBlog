/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/12/26 10:28
*/
package routers

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"leeblog.com/app/http/middleware"
	"leeblog.com/config"
	"leeblog.com/pkg/utils"
	"path/filepath"
	"strings"
)

var engine *gin.Engine

func loadTemplates(templatesDir string, module ...string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	for _, moduleName := range module {
		layouts, err := filepath.Glob("./" + templatesDir + "/" + moduleName + "/*.html")
		if err != nil {
			panic(err.Error())
		}

		includes, err := filepath.Glob("./" + templatesDir + "/" + moduleName + "/**/*.html")
		if err != nil {
			panic(err.Error())
		}
		for _, include := range includes {
			layoutCopy := make([]string, len(layouts))
			copy(layoutCopy, layouts)
			files := append(layoutCopy, include)
			//r.AddFromFiles(strings.Replace(filepath.ToSlash(include), templatesDir, "", -1), files...)
			r.AddFromFilesFuncs(strings.Replace(filepath.ToSlash(include), templatesDir, "", -1), template.FuncMap{
				"html":     html,
				"format":   format,
				"adminUrl": adminUrl,
				"url":      url,
			}, files...)
		}

	}

	return r
}

func html(s string) template.HTML {
	return template.HTML(s)
}
func format(t utils.JSONTime, layout string) string {
	//return t.Format("2006-01-02 15:04:05")
	return t.Format(layout)
}
func adminUrl(path string) string {
	return "/admin/" + strings.Trim(path, "/")
}
func url(path string) string {
	return "/admin/" + strings.Trim(path, "/")
}

var server = config.ServerConfig()

func init() {
	if server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine = gin.Default()
	engine.Use(middleware.LoggerToFile())
	engine.HTMLRender = loadTemplates("templates", "web", "admin")

}

// router router
func router() *gin.Engine {
	return engine
}

// Run run
func Run() {
	apiRouter()
	webRouter()
	adminRouter()
	_ = router().Run(":" + server.Port)
}
