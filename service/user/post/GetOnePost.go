package post

import (
	mysql_post "GO1/database/mysql/post"
	"GO1/models"
	models_post "GO1/models/post"
)

func GetOnePost(postId int64) *models.HandleFuncResp {
	var Resp models.HandleFuncResp
	post := mysql_post.GetPostByPostId(postId)

	Resp.Ok = post.PostID > 0

	if Resp.Ok {
		Resp.Data = models_post.BuildPostResponse(post)
	} else {
		Resp.Msg = "获取帖子失败"
	}
	return &Resp
}
