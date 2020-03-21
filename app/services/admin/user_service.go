package admin

import (
	"crypto/md5"
	"fmt"
	"leeblog.com/app/models"
)

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

func (us *UserService) DeleteUserById(userId uint) (success bool) {

	db := models.DB().Delete(&models.User{}, "id=?", userId)
	if db.RowsAffected > 0 {
		success = true
	} else {
		success = false
	}
	return
}
func (us *UserService) AddUser(userName string, password string) (userModel models.User, err error) {
	md5Pwd := fmt.Sprint(md5.Sum([]byte(password)))
	user := models.User{Name: userName, Password: md5Pwd}
	tmpUser := &models.User{}
	db := models.DB().First(tmpUser, "name=", userName)
	if tmpUser.Name == user.Name {
		userModel = user
		err = nil
		return
	} else {
		err = db.Error
		return
	}

}
