/*
@Author :   寒云
@Email  :   1355081829@qq.com
@Time : 2020/1/16 17:40
*/
package validator

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func init() {
	// NOTE: ommitting allot of error checking for brevity
	//
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ := uni.GetTranslator("zh")

	validate = validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})
	_ = validate.RegisterValidation("mobile", mobile)
	_ = zhtranslations.RegisterDefaultTranslations(validate, trans)
}
func RegisterTagTranslation(tag string, message string) {
	trans, _ := uni.GetTranslator("zh")
	_ = validate.RegisterTranslation(tag, trans, func(ut ut.Translator) error {
		return ut.Add(tag, message, false)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, err := ut.T(fe.Tag(), fe.Field())
		if err != nil {
			return fe.(error).Error()
		}
		return t
	})
}

func TranslateIndividual(data interface{}, keyMap map[string]string) (errMap map[string]string, ok bool) {
	trans, _ := uni.GetTranslator("zh")
	_ = zhtranslations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(data)
	errMap = map[string]string{}
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			if _, v := keyMap[e.Field()]; v == true {
				errMap[e.Field()] = strings.Replace(e.Translate(trans), e.Field(), keyMap[e.Field()], 1)
			} else {
				errMap[e.Field()] = e.Translate(trans)
			}
		}
		return errMap, false
	} else {
		return errMap, true
	}
}
