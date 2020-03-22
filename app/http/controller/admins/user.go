package admins

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/http/response"
	"leeblog.com/app/models"
	"leeblog.com/app/services/admin"
	"leeblog.com/pkg/utils"
	"strconv"
)

// @Tags admin
// @Summary 用户列表
// @Description # 请求参数
// @Description | 参数 | 说明 | 备注 |
// @Description | :-----| :----- | :----- |
// @Description |page|页码||
// @Description |keyWord|页码||
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
	keyWord := c.DefaultQuery("keyWord", "")
	userSlice, count := us.GetUserList(page, limit, keyWord)
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
// @Param userName formData string false "用户名"
// @Param password formData string false "密码"
func AddUser(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	if u := admin.NewUserService().GetUserByName(userName); u.ID > 0 {
		response.APIFailure(c, "用户已经存在")
		return
	}
	user := models.User{}
	user.Name = userName
	user.Password = utils.EncodeMD5(password)
	if res := admin.NewUserService().AddUser(user); !res {
		response.APIFailure(c, "添加失败")
	} else {
		response.APISuccess(c, user)
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
	} else {
		response.APIResponse(c, response.MissingParameters, nil)
	}

}
