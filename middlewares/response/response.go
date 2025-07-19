package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int `json:"code"`
	ErrCode int `json:"err_code"`
	Data    any `json:"data"`
	Message any `json:"message"`
}

const (
	SUCCESS = 0
	FAIL    = 1
)

func Result(code int, errCode int, data interface{}, message interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		ErrCode: errCode,
		Data:    data,
		Message: message,
	})
}

func Ok(data any, message any, c *gin.Context) {
	Result(SUCCESS, 0, data, message, c)
}
func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, 0, data, "成功", c)
}
func OkWithMessage(message any, c *gin.Context) {
	Result(SUCCESS, 0, map[string]any{}, message, c)
}

func Fail(data any, message any, c *gin.Context) {
	Result(FAIL, 0, data, message, c)
}
func FailWithMessage(message any, c *gin.Context) {
	Result(FAIL, 0, map[string]any{}, message, c)
}
func FailWithCode(errorCode ErrorCode, c *gin.Context) {
	message, ok := ErrorMap[errorCode]
	if !ok {
		Result(FAIL, int(errorCode), errorCode, "未知错误", c)
		return
	}
	Result(FAIL, int(errorCode), map[string]any{}, message, c)
}
