package admin

import "leeblog.com/app/models"

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}
func (us *UserService) GetUserList(page int, limit int) (userSlice []models.User, count int) {
	models.DB().Offset((page - 1) * limit).Limit(limit).Find(&userSlice)
	models.DB().Model(&models.User{}).Count(&count)
	return
}
