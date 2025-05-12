package validator

import (
	"GO1/global"
	"GO1/middlewares/translator"
	"GO1/middlewares/validator"
	"GO1/models"
	"github.com/gin-gonic/gin/binding"
	v10 "github.com/go-playground/validator/v10"
)

func InitValidator() {
	// 创建一个校验器接口v
	if v, ok := binding.Validator.Engine().(*v10.Validate); ok {
		global.Validator = v
	} else {
		global.Validator = v10.New()
	}
}

func DefinedValidator() {
	v := global.Validator
	forTranslator(v)
	forUser(v)
}

func forTranslator(v *v10.Validate) {
	v.RegisterTranslation("password", global.Trans, translator.PwdRegisterFunc, translator.PwdFunc)
	// 错误时把struct的变量名改为json中的
	v.RegisterTagNameFunc(translator.TagName)
}

func forUser(v *v10.Validate) {
	v.RegisterValidation("password", validator.UserPassword)
	v.RegisterStructValidation(validator.SignUpParamStructLevelValidation, models.Register{})
}
