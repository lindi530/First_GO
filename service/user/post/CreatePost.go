package post

import (
	mysql "GO1/database/mysql/user"
	models_post "GO1/models/post"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context, p models_post.CreatePost) models_post.PostResponse {
	post := CreatePostHandler(p)

	mysql.CreatePost(&post)

	responsePost := models_post.BuildPostResponse(c, GetPostByPostId(post.PostID))

	return responsePost
}
