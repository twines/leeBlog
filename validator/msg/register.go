/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/17 10:22
*/
package msg

//Name            string `json:"name" form:"name" validate:"required,min=3"`
//Password        string `json:"password" form:"password" validate:"required,max=20,min=6"`
//ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" validate:"required,max=20,min=6"`
//QQ              string `json:"qq" form:"qq"  validate:"required,min=3"`
//Mobile          string `json:"mobile" form:"mobile" validate:"require"`
//Email           string `json:"email" form:"email"`
var (
	Login = map[string]string{"name": "用户名", "password": "密码", "confirmPassword": "确认密码", "qq": "QQ", "mobile": "手机号", "email": "邮箱", "code": "验证码"}
)
