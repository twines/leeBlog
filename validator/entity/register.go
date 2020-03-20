/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/17 10:19
*/
package entity

type User struct {
	Name            string `json:"name" form:"username" validate:"required,min=4,max=16"`
	Password        string `json:"password" form:"password" validate:"required,max=18,min=6"`
	ConfirmPassword string `gorm:"-" json:"confirmPassword" form:"password_confirm" validate:"required,max=18,min=6,eqfield=ConfirmPassword"`
	QQ              string `json:"qq" form:"qq"  validate:"required,numeric,min=6"`
	Mobile          string `json:"mobile" form:"mobile" validate:"required,mobile"`
	Code            string `gorm:"-" json:"code" form:"code" validate:"required,len=4"`
}
