package post

import (
	"GO1/global"
	"GO1/models/post"
)

func DeletePost(post post.Post) {
	global.DB.Delete(post)
}
