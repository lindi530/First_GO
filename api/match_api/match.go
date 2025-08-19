package match_api

import (
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/pkg/jwt"
	"GO1/service/match_service"
	"github.com/gin-gonic/gin"
)

func (MatchAPI) MatchRequest(c *gin.Context) {

	global.Logger.Info("host: ", c.Request.Host)
	userid := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))
	global.Logger.Infof("userid is %v", userid)
	err := match_service.SendMatchRequest(userid)
	if err != nil {
		response.FailWithMessage("发生匹配请求失败！", c)
	}
	response.OkWithMessage("发生匹配请求成功！", c)
}
