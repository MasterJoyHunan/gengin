package translator

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

const labelName = "label"

func init() {
	validate := binding.Validator.Engine().(*validator.Validate)
	uni := ut.New(en.New(), zh.New())
	trans, _ = uni.GetTranslator("zh")

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get(labelName)
		if label == "" {
			return field.Name
		}
		return label
	})
	customError(validate)
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
}

// Translate 校验异常
func Translate(errs error) string {
	errors, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs.Error()
	}
	return errors[0].Translate(trans)
}

// customError 自定义异常
func customError(validate *validator.Validate) {
	validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
		if len(fl.Field().String()) != 11 {
			return false
		}
		return true
	})
	validate.RegisterTranslation(
		"phone",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("phone", "{0}输入错误", false)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(fe.Tag(), fe.Field())
			return t
		})
}
