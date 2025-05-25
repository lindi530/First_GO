package post

import (
	mysql "GO1/database/mysql/post"
	service_user "GO1/service/user"
	"github.com/gin-gonic/gin"
)

func DeleteUserPost(c *gin.Context, postId int64, userId int64) bool {

	if service_user.AuthUser(c, userId) == false {
		return false
	}

	post := GetPostByPostId(postId)
	if post.PostID == 0 {
		return false
	}

	if post.UserID != userId {
		return false
	}

	mysql.DeletePost(post)
	return true
}
