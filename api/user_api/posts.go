package user_api

import (
	mysql "GO1/database/mysql/user"
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models"
	service "GO1/service/user/post"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func (UserAPI) GetUserPosts(c *gin.Context) {
	// 得到此时登录的userID
	userId, _ := strconv.Atoi(c.Param("id"))

	// 通过server获得帖子信息
	posts, err := service.GetPosts(int64(userId))
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

func (UserAPI) SaveUserPost(c *gin.Context) {
	// 数据校验
	post := models.CreatePost{}
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&post); err != nil || int64(userId) != post.UserID {
		global.Logger.Error("<UNK>", zap.Error(err))
		response.FailWithCode(response.BadRequest, c)
		return
	}
	result := mysql.CheckUser(mysql.UserIdParam(userId))
	if result == false {
		response.FailWithCode(response.InvalidUser, c)
		return
	}

	newPost := service.SavePost(post)
	if !result {
		response.FailWithCode(response.ServiceError, c)
		return
	}
	response.OkWithData(gin.H{
		"post": newPost,
	}, c)
}
