package response

type ErrorCode int

const (
	SettingsError ErrorCode = 1000 * iota
	BadRequest
	NeedLogin
	InvalidAccessToken
	InvalidRefreshToken
	InvalidLoginInfo
	Logout
	Login
	Register
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:       "系统错误",
		BadRequest:          "错误请求",
		NeedLogin:           "请登录",
		InvalidAccessToken:  "无效的assess_token",
		InvalidRefreshToken: "无效的refresh_token",
		InvalidLoginInfo:    "用户名或密码错误",
		Logout:              "已登出",
		Login:               "登录成功",
		Register:            "注册成功",
	}
)
