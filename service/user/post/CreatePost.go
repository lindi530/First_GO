package post

import (
	mysql "GO1/database/mysql/user"
	"GO1/models/post"
)

func CreatePost(p post.CreatePost) post.Post {
	post := CreatePostHandler(p)

	mysql.CreatePost(&post)

	return post
}
