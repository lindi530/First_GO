package post

import (
	"GO1/global"
	"GO1/models"
)

func DeletePost(post models.Post) {
	global.DB.Delete(post)
}
