package admin

import (
	"leeblog.com/app/models"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}
func (us *UserService) GetUserList(page int, limit int, keyWord string) (userSlice []models.User, count int) {
	if keyWord != "" {
		models.DB().Offset((page-1)*limit).Limit(limit).Find(&userSlice, "name like ? or mobile like ? or qq like ", keyWord, keyWord, keyWord)
		count = len(userSlice)
	} else {
		models.DB().Offset((page - 1) * limit).Limit(limit).Find(&userSlice)
		models.DB().Model(&models.User{}).Count(&count)
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
