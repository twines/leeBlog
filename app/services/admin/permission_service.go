package admin

import "leeblog.com/app/models"

type PermissionService struct {
}

func NewPermissionService() *PermissionService {
	return &PermissionService{}
}
func (p *PermissionService) GetPermissionList() (permissionSlice []models.Permission) {
	models.DB().Find(&permissionSlice)
	return
}
