package translator

import (
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
)

func PwdRegisterFunc(tr ut.Translator) error {
	return tr.Add("password", "{0} 格式不正确", false)
}

func PwdFunc(tr ut.Translator, fe v10.FieldError) string {
	msg, _ := tr.T("password", fe.Field())
	return msg
}
