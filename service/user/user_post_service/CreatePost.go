package user_post_service

import (
	"GO1/database/mysql/user_mysql"
	models_post "GO1/models/post_model"
)

func CreatePost(userId int64, p models_post.CreatePost) models_post.PostResponse {
	post := CreatePostHandler(p)

	user_mysql.CreatePost(&post)
	post.Author = user_mysql.FindUserByUserId(userId)
	
	responsePost := BuildPostResponse(userId, post)

	return responsePost
}
