package admins

import (
	"github.com/gin-gonic/gin"
	"leeblog.com/app/models"
	"leeblog.com/app/services/admin"
	"leeblog.com/pkg/utils"
	"net/http"
	"strconv"
)

func RoleList(c *gin.Context) {
	page, limit := 1, 2
	if p, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	roleSlice, count := admin.NewRoleService().GetRoleList(page, limit)
	totalPage := utils.Page(count, limit)
	mp := map[string]interface{}{}
	mp["totalCount"] = count
	mp["limit"] = limit
	mp["totalPage"] = totalPage
	mp["data"] = roleSlice
	c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success", "data": mp})
}
func RoleDelete(c *gin.Context) {
	roleId := 0
	if id, err := strconv.Atoi(c.Param("roleId")); err == nil {
		roleId = id
	}
	if roleId <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "roleId必须大于0"})
	} else {
		role := models.Role{}
		role.ID = uint(roleId)
		res := admin.NewRoleService().DeleteRole(role)
		if res {
			c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "删除失败"})
		}
	}
}
func RoleAdd(c *gin.Context) {
	roleName := c.PostForm("roleName")
	description := c.PostForm("description")
	res := admin.NewRoleService().AddRole(models.Role{RoleName: roleName, Description: description})
	if res {
		c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "添加失败"})
	}
}
func RolePermission(c *gin.Context) {
	roleId := 0
	if id, err := strconv.Atoi(c.Param("roleId")); err == nil {
		roleId = id
	}
	if roleId <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "roleId必须大于0"})
		return
	}
	role := admin.NewRoleService().GetRoleById(uint(roleId))
	if role.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "角色不存在"})
		return
	}
	permissionSlice := admin.NewRoleService().GetRolePermission(role)
	allPermissions := admin.NewPermissionService().GetPermissionList()
	mp := gin.H{}
	mp["permissionSlice"] = permissionSlice
	mp["allPermissions"] = allPermissions
	c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success", "data": mp})
}
func RolePermissionUpdate(c *gin.Context) {
	roleId := 0
	if id, err := strconv.Atoi(c.Param("roleId")); err == nil {
		roleId = id
	}
	if roleId <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "roleId必须大于0"})
		return
	}
	role := admin.NewRoleService().GetRoleById(uint(roleId))
	if role.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 40000, "message": "角色不存在"})
		return
	}
	admin.NewRoleService().DeleteRolePermission(role)
	permissionList := c.PostFormMap("permissionList")
	var permissionSlice []models.Permission
	p := models.Permission{}
	for _, v := range permissionList {
		if i, err := strconv.Atoi(v); err == nil {
			p.ID = uint(i)
			permissionSlice = append(permissionSlice, p)
		}
	}
	role.Permission = permissionSlice
	admin.NewRoleService().UpdateRole(role)
	c.JSON(http.StatusOK, gin.H{"code": 20000, "message": "success"})
}
