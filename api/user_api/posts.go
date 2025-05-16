package user_api

import (
	"GO1/middlewares/response"
	"GO1/service/context"
	"GO1/service/user/post"
	"github.com/gin-gonic/gin"
)

func (UserAPI) UserPosts(c *gin.Context) {

	// 得到此时登录的userID
	userId, _ := context.GetContext(c, context.CUserIdKey).(int64)

	// 通过server获得帖子信息
	posts, err := post.GetPosts(userId)
	if err != nil {
		response.FailWithMessage("出错", c)
		return
	}

	// 返回信息
	response.OkWithData(gin.H{
		"len":   len(posts),
		"posts": posts,
	}, c)
}
