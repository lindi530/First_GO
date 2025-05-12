package translator

import (
	"GO1/global"
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 定义一个全局翻译器T

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {

	v := global.Validator

	// 注册一个获取json tag的自定义方法

	//v.RegisterValidation("UserPassword", middlewares.UserPassword)   // 注册一个校验器
	//v.RegisterStructValidation(middlewares.UserList, models.Users{}) // 校验器用于哪一种接口

	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器

	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni := ut.New(enT, zhT, enT)

	// locale 通常取决于 http 请求头的 'Accept-Language'
	var ok bool
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	global.Trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
	}

	// 注册翻译器
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, global.Trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
	}
	return
}
