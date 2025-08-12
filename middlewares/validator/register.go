package validator

import (
	"GO1/models/user_model"
	"github.com/go-playground/validator/v10"
)

// 把Password 变成 password
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(user_model.ParamRegister)
	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
