/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2019/12/31 17:52
*/
package main

import (
	"leeblog.com/routers"
)

// @title LeeGo
// @Schemes http https
// @version 1.0
// @description 这个API接口文档
// @termsOfService http://swagger.io/terms/
// @Security bearer
// @securityDefinitions.apikey bearer
// @in header
// @name Authorization
// @contact.name 寒云
// @contact.url http://www.swagger.io/support
// @contact.email 1355081829@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	routers.Run()
}
