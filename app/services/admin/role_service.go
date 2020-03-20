package admin

import (
	"leeblog.com/app/models"
)

type RoleService struct {
}

func NewRoleService() *RoleService {
	return &RoleService{}
}
func (rs *RoleService) GetRoleList(page int, limit int) (roleSlice []models.Role, count int) {
	models.DB().Order("id desc").Offset((page - 1) * limit).Limit(limit).Find(&roleSlice).Offset(-1).Limit(-1).Count(&count)
	return
}
func (rs *RoleService) GetRoleById(id uint) (role models.Role) {
	models.DB().First(&role, id)
	return
}
func (rs *RoleService) UpdateRole(role models.Role) {
	models.DB().Save(&role)
}
func (rs *RoleService) DeleteRole(role models.Role) bool {
	res := models.DB().Delete(&role)
	if res.RowsAffected == 1 {
		return true
	}
	return false
}
func (rs *RoleService) AddRole(role models.Role) bool {
	res := models.DB().Create(&role)
	if res.RowsAffected > 0 {
		return true
	}
	return false
}
func (rs *RoleService) GetRolePermission(role models.Role) (permissionSlice []models.Permission) {
	models.DB().Model(&role).Related(&permissionSlice, "permission")
	return
}
func (rs *RoleService) DeleteRolePermission(role models.Role) bool {
	res := models.DB().Where("role_id=?", role.ID).Unscoped().Delete(&models.RolePermission{})
	if res.RowsAffected > 0 {
		return true

	}
	return false
}
