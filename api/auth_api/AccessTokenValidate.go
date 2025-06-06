package auth_api

import (
	"GO1/middlewares/response"
	"github.com/gin-gonic/gin"
)

func (AuthAPI) AccessTokenValidate(c *gin.Context) {
	response.OkWithMessage("成功", c)
}
