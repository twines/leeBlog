package web

import "leeblog.com/app/models"

type User struct {
}

func NewUser() *User {
	return &User{}
}
func (u User) AddUser(user User) {
	models.DB().Create(&user)
}
func (u User) GetUser(user models.User) (userModel models.User) {
	models.DB().Where(user).First(&userModel)
	return
}
