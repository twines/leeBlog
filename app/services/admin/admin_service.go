package admin

import "leeblog.com/app/models"

type AdminsService struct {
}

func NewAdminsService() *AdminsService {
	return &AdminsService{}
}

func (a *AdminsService) GetAdminByName(name string) (admin models.Admin) {
	models.DB().Where("name=?", name).First(&admin)
	return
}
func (a *AdminsService) GetAdminList(page int, limit int) (adminSlice []models.Admin, count int) {
	models.DB().Offset((page - 1) * limit).Limit(limit).Find(&adminSlice).Offset(-1).Limit(-1).Count(&count)
	return
}
