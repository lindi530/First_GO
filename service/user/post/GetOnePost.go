package post

import (
	mysql_post "GO1/database/mysql/post"
	"GO1/models"
	"github.com/gin-gonic/gin"
)

func GetOnePost(c *gin.Context, userId, postId int64) *models.HandleFuncResp {
	var Resp models.HandleFuncResp
	post := mysql_post.GetPostByPostId(postId)

	Resp.Ok = post.PostID > 0

	if Resp.Ok {
		Resp.Data = BuildPostResponse(c, userId, post)
		mysql_post.PostViews(post.PostID)
	} else {
		Resp.Msg = "获取帖子失败"
	}
	return &Resp
}
