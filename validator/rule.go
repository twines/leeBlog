/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/19 10:12
*/
package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func mobile(fl validator.FieldLevel) bool {
	RegisterTagTranslation("mobile", "{0}不是一个可用的手机号码")
	ok, _ := regexp.MatchString(`^1[3-9][0-9]{9}$`, fl.Field().String())
	return ok
}
