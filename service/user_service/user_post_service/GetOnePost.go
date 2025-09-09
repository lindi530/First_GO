package user_post_service

import (
	"GO1/database/mysql/post_mysql"
	"GO1/models"
	"GO1/models/post_model"
)

func GetOnePost(userId, postId int64) *models.HandleFuncResp {
	var Resp models.HandleFuncResp

	var post post_model.Post
	err := post_mysql.GetPostByPostId(postId, &post)

	Resp.Ok = err == nil

	if Resp.Ok {
		post_mysql.PostViews(post.PostID)
		Resp.Data = BuildPostResponse(userId, post)
	} else {
		Resp.Msg = "获取帖子失败"
	}
	return &Resp
}
