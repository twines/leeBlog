package admins

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/http/response"
	"leeblog.com/app/models"
	"leeblog.com/app/services/admin"
	"leeblog.com/pkg/utils"
	"net/http"
	"strconv"
)

func AdminIndex(c *gin.Context) {
	JWTClaims, _ := utils.VerifyToken(c)
	admin := JWTClaims.User.(map[string]interface{})
	m := map[string]interface{}{}
	m["roles"] = []string{"admin"}
	m["introduction"] = "I am a super administrator"
	m["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
	m["name"] = admin["Name"]
	r := response.Response{Data: m, Code: 20000}
	c.JSON(http.StatusOK, r)
}

// @Tags admin
// @Summary 管理员列表
// @Description # 请求参数
// @Description | 参数 | 说明 | 备注 |
// @Description | :-----| :----- | :----- |
// @Description |page|页码|默认是1，可以不填|
// @Description # 返回参数参数
// @Description | 参数 | 说明 | 备注 |
// @Description | :-----| :----- | :----- |
// @Description |name|用户名||
// @Description |avatar|头像||
// @Description |roleId|角色id||
// @Description |limit|每页数量||
// @Description |totalCount|总数据量||
// @Description |totalPage|总页码数||
// @Description |status|状态|20000表示成功|
// @Description |code|状态码||
// @Description |message|返回的信息||
// @Description |data|返回的具体数据||
// @Produce  json
// @Security bearer
// @Param page query string false "页码"
// @Success 200 {object} response.Response "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/v1/admin/list [get]
func AdminList(c *gin.Context) {
	page, limit := 1, 2
	if p, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	adminSlice, count := admin.NewAdminsService().GetAdminList(page, limit)
	if len(adminSlice) > 0 {
		for k, a := range adminSlice {
			adminSlice[k].Role = admin.NewRoleService().GetRoleById(a.RoleId)
		}
	}
	totalPage := utils.Page(count, limit)
	mp := map[string]interface{}{}
	mp["totalCount"] = count
	mp["limit"] = limit
	mp["totalPage"] = totalPage
	mp["data"] = adminSlice
	c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success", "data": mp})
}
func AdminDetail(c *gin.Context) {
}
func AdminAdd(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "名称不可以为空"})
		return
	}
	if password == "" {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "密码不可以为空"})
		return
	}
	var a models.Admin
	if a = admin.NewAdminsService().GetAdminByName(name); a.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "管理员已经存在"})
		return
	}
	a.Name = name
	a.Password = utils.EncodeMD5(password)
	if res := admin.NewAdminsService().AddAdmin(a); res {
		c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "失败"})
	}
}
