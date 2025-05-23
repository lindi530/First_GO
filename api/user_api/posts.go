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
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)

	// 通过server获得帖子信息
	posts, err := service.GetUserPosts(userId)
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

func (UserAPI) CreateUserPost(c *gin.Context) {
	// 数据校验
	post := models.CreatePost{}
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err := c.ShouldBindJSON(&post); err != nil || userId != post.UserID {
		global.Logger.Error("<UNK>", zap.Error(err))
		response.FailWithCode(response.BadRequest, c)
		return
	}
	result := mysql.CheckUser(mysql.UserIdParam(userId))
	if result == false {
		response.FailWithCode(response.InvalidUser, c)
		return
	}

	newPost := service.CreatePost(post)
	if !result {
		response.FailWithCode(response.ServiceError, c)
		return
	}
	response.OkWithData(gin.H{
		"post": newPost,
	}, c)
}

func (UserAPI) DeleteUserPost(c *gin.Context) {
	// 前端数据的提取，校验
	userId, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	postId, _ := strconv.ParseInt(c.Param("post_id"), 10, 64)

	// 服务
	res := service.DeleteUserPost(c, postId, userId)
	if res == false {
		response.FailWithCode(response.NoAuthority, c)
		return
	}

	response.OkWithData(gin.H{
		"message": "删除成功",
	}, c)
}
