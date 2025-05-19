package post

import (
	mysql "GO1/database/mysql/user"
	"GO1/models"
)

func SavePost(p models.CreatePost) models.Post {
	post := CreatePostHandler(p)

	mysql.SavePost(post)

	return post
}
