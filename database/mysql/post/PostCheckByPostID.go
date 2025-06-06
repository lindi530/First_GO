package post

import (
	"GO1/global"
	"GO1/models/post"
)

func PostCheckByPostID(postId int64) bool {
	var p post.Post
	err := global.DB.Select("id").First(&p, postId).Error //First(&p, postId)：等价于 WHERE id = postId LIMIT 1；
	return err == nil
}
