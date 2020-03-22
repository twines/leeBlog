package admins

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"leeblog.com/app/http/response"
	"leeblog.com/app/services/admin"
	"leeblog.com/pkg/utils"
	"net/http"
)

func LoginIndex(c *gin.Context) {
	fmt.Println(utils.AuthCheck(c, "admin"))
	utils.View(c, "/admin/login/index.html")
}

// @Tags admin
// @Summary 登录
// @Description # 请求参数
// @Description | 参数 | 说明 | 备注 |
// @Description | :-----| :----- | :----- |
// @Description |username|用户名||
// @Description |password|密码||
// @Description # 返回参数参数
// @Description | 参数 | 说明 | 备注 |
// @Description | :-----| :----- | :----- |
// @Description |code|状态码||
// @Description |message|返回的信息||
// @Description |data|返回的具体数据||
// @Produce  json
// @Param userName formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} response.Response "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/v1/login [post]
func LoginDoLogin(c *gin.Context) {
	userName := c.PostForm("userName")
	admin := admin.NewAdminsService().GetAdminByName(userName)
	token, _ := utils.GenToken(admin)
	r := response.Response{Data: map[string]string{"token": token}, Code: 20000}
	c.JSON(http.StatusOK, r)
}

// @Tags admin
// @Summary 退出
// @Produce  json
// @Param Authorization header string true "用户token"
// @Success 200 {object} response.Response "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/v1/logout [post]
func Logout(c *gin.Context) {
	r := response.Response{Code: 20000}
	c.JSON(http.StatusOK, r)
}
