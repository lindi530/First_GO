package post

import (
	mysql "GO1/database/mysql/post"
)

func DeleteUserPost(postId int64, userId int64) bool {
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
