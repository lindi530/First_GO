package post

import (
	mysql_post "GO1/database/mysql/post"
	"GO1/models"
	"GO1/models/post"
)

func GetOnePost(userId, postId int64) *models.HandleFuncResp {
	var Resp models.HandleFuncResp

	var post post.Post
	err := mysql_post.GetPostByPostId(postId, &post)

	Resp.Ok = err == nil

	if Resp.Ok {
		mysql_post.PostViews(post.PostID)
		Resp.Data = BuildPostResponse(userId, post)
	} else {
		Resp.Msg = "获取帖子失败"
	}
	return &Resp
}
