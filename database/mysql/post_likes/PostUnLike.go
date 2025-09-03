package post_likes

import (
	"GO1/global"
	"GO1/models"
	"GO1/models/post_model"
	"gorm.io/gorm"
)

func PostUnLike(userId, postId int64) (resp models.HandleFuncResp) {

	err := global.DB.
		Where("user_id = ? AND post_id = ?", userId, postId).
		Delete(&post_model.PostLike{}).Error

	if err != nil {
		resp.Msg = "点赞数据删除失败"
		resp.Ok = false
		return
	}

	// 点赞数 -1（可选，防止负数）
	global.DB.Model(&post_model.Post{}).
		Where("id = ?", postId).
		UpdateColumn("likes", gorm.Expr("GREATEST(likes - 1, 0)")) // 防止减到负数

	resp.Msg = "点赞数据删除成功"
	resp.Ok = true
	return
}
