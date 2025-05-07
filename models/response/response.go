package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

const (
	SUCCESS = 0
	FAIL    = 1
)

func Result(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func Ok(data any, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}
func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "成功", c)
}
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, message, c)
}

func Fail(data any, message string, c *gin.Context) {
	Result(FAIL, data, message, c)
}
func FailWithMessage(message string, c *gin.Context) {
	Result(FAIL, map[string]any{}, message, c)
}
func FailWithCode(errorCode ErrorCode, c *gin.Context) {
	message, ok := ErrorMap[errorCode]
	if !ok {
		Result(FAIL, errorCode, "未知错误", c)
		return
	}
	Result(FAIL, map[string]any{}, message, c)
}
