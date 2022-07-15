package validate

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	//zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	vt    *validator.Validate
	trans ut.Translator
)

// 初始化验证器和翻译
func init() {
	registerDefault()
}

func registerDefault() {
	vt = validator.New()
	trans = registerTranslate(vt)
}

func registerTranslate(v *validator.Validate) ut.Translator {
	//zhT := zh.New()
	enT := en.New()

	uni := ut.New(enT, enT)
	tr, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(v, tr)
	return tr
}

// Struct 验证结构体
func Struct(data interface{}) error {
	return TransErr(vt.Struct(data))
}

func Trans(err error) string {
	return transErr(trans, err)
}

func TransErr(err error) error {
	if err == nil {
		return nil
	}
	return errors.New(Trans(err))
}

func TransGin(err error) string {
	return transErr(ginTrans, err)
}

func TransGinErr(err error) error {
	if err == nil {
		return nil
	}
	return errors.New(TransGin(err))
}

// Trans 翻译错误信息
func transErr(trans ut.Translator, err error) (msg string) {
	if err == nil {
		return
	}

	switch v := err.(type) {
	case validator.ValidationErrors:
		if len(v) > 0 {
			msg = v[0].Translate(trans)
		} else {
			msg = v.Error()
		}
	case validator.FieldError:
		msg = v.Translate(trans)
	default:
		msg = err.Error()
	}
	return
}
