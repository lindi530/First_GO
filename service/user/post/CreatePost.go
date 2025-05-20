package post

import (
	mysql "GO1/database/mysql/user"
	"GO1/models"
)

func CreatePost(p models.CreatePost) models.Post {
	post := CreatePostHandler(p)

	mysql.CreatePost(&post)

	return post
}
