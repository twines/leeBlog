package admins

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/http/response"
	"leeblog.com/app/services/admin"
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
	response.APISuccess(c, mp)
	//c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success", "data": mp})
}

//@Tags admin
// @Summary 添加人员
// @Description # 请求参数
func AddUser(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	if user, err := admin.NewUserService().AddUser(userName, password); err != nil {
		response.APIFailure(c, err)
		//c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "failure", "data": err})
	} else {
		response.APISuccess(c, user)

		//c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success", "data": user})
	}

}

//@Tags admin
// @Description # 请求参数
func DeleteUserById(c *gin.Context) {
	if userIdStr, ok := c.GetPostForm("userId"); ok {
		if userId, err := strconv.Atoi(userIdStr); err == nil {
			rev := admin.NewUserService().DeleteUserById(uint(userId))
			if rev {
				response.APISuccess(c, nil)

				//c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success"})
			} else {
				response.APIFailure(c, nil)

				//c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "failure"})
			}
		} else {
			response.APIFailure(c, err)

			//c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "success", "data": err})
		}
	}

}
