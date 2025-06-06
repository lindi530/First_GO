package post

import (
	mysql "GO1/database/mysql/user"
	models_post "GO1/models/post"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context, userId int64, p models_post.CreatePost) models_post.PostResponse {
	post := CreatePostHandler(p)

	mysql.CreatePost(&post)

	responsePost := BuildPostResponse(c, userId, GetPostByPostId(post.PostID))

	return responsePost
}
