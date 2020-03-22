package admin

import (
	"leeblog.com/app/models"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}
func (us *UserService) GetUserList(page int, limit int, keyword string) (userSlice []models.User, count int) {
	if keyword != "" {
		keyword = "%" + keyword + "%"
		models.DB().Offset((page-1)*limit).Limit(limit).Find(&userSlice, "name like ? or mobile like ? or qq like ?", keyword, keyword, keyword)
	} else {
		models.DB().Offset((page - 1) * limit).Limit(limit).Find(&userSlice).Offset(-1).Limit(-1).Count(&count)
	}
	return
}

func (us *UserService) DeleteUserById(userId uint) (success bool) {
	db := models.DB().Delete(&models.User{}, "id=?", userId)
	if db.RowsAffected > 0 {
		success = true
	} else {
		success = false
	}
	return
}
func (us *UserService) GetUserByName(username string) (user models.User) {
	models.DB().First(&user, "name =?", username)
	return
}
func (us *UserService) AddUser(user models.User) bool {
	res := models.DB().Create(&user)
	if res.RowsAffected > 0 {
		return true
	}
	return false
}
