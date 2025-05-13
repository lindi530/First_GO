package response

type ErrorCode int

const (
	SettingsError ErrorCode = 1001
	BadRequest    ErrorCode = 400
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		BadRequest:    "错误请求",
	}
)
