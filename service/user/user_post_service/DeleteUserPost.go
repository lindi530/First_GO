package user_post_service

import (
	service_user "GO1/service/user"
	"github.com/gin-gonic/gin"
)

func DeleteUserPost(c *gin.Context, postId int64, userId int64) bool {

	if service_user.AuthUser(c, userId) == false {
		return false
	}

	//post := CheckPostExit(postId)
	//if post.PostID == 0 {
	//	return false
	//}
	//
	//if post.UserID != userId {
	//	return false
	//}

	//mysql.DeletePost(post)
	return true
}
