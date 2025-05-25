package user

import (
	"GO1/service/context"
	"github.com/gin-gonic/gin"
)

func CheckUserId(c *gin.Context, userId int64) bool {

	ctxUserId := context.GetContext(c, context.CtxUserIdKey)
	//global.Logger.Info(fmt.Sprintf("userId: %d   ctxUserId: %d", userId, ctxUserId))
	return userId == ctxUserId
}
