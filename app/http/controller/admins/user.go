package admins

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/services/admin"
	"net/http"
	"strconv"
)

// @Tags admin
// @Summary 用户列表
// @Description # 请求参数
// @Description | 参数 | 说明 | 备注 |
// @Description | :-----| :----- | :----- |
// @Description |page|页码||
// @Description # 返回参数参数
// @Description | 参数 | 说明 | 备注 |
// @Description | :-----| :----- | :----- |
// @Description |code|状态码||
// @Description |message|返回的信息||
// @Description |data|返回的具体数据||
// @Produce  json
// @Security bearer
// @Param page query string false "页码"
// @Success 200 {object} response.Response "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin/v1/user/list [get]
func UserList(c *gin.Context) {
	page, limit := 1, 1
	if p, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	us := admin.NewUserService()
	userSlice, count := us.GetUserList(page, limit)
	mp := map[string]interface{}{}
	mp["totalCount"] = count
	mp["limit"] = limit
	mp["totalPage"] = count
	mp["data"] = userSlice
	c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success", "data": mp})
}
