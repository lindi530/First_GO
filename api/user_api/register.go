package user_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models"
	service "GO1/service/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

func (UserAPI) Register(c *gin.Context) {
	// 数据校验
	register := models.ParamRegister{}
	if err := c.ShouldBindJSON(&register); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 验证错误类型是否是校验器错误
		if !ok {
			response.FailWithMessage(err.Error(), c)
			global.Logger.Error("注册数据错误！", err.Error())
			return
		}

		result := removeTopStruct(errs.Translate(global.Trans))
		global.Logger.Error(fmt.Sprintf("注册数据错误！ %s", result))
		response.FailWithMessage(result, c)
		return
	}
	// 调用服务
	result := service.Register(register)

	// 返回响应
	if !result {
		response.FailWithMessage("已存在该用户！", c)
		return
	}

	response.Ok(register, "注册成功！", c)
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
