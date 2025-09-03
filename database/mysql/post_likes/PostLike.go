package post_likes

import (
	"GO1/global"
	"GO1/models"
	"GO1/models/post_model"
	"gorm.io/gorm"
	"time"
)

func PostLike(userId, postId int64) (resp models.HandleFuncResp) {
	postLike := post_model.PostLike{
		UserId:    userId,
		PostId:    postId,
		CreatedAt: time.Now(),
	}
	if err := global.DB.Create(&postLike).Error; err != nil {
		resp.Msg = "点赞数据保存失败"
		resp.Ok = false
		return
	}
	// 点赞数 +1
	global.DB.Model(&post_model.Post{}).
		Where("id = ?", postId).
		UpdateColumn("likes", gorm.Expr("GREATEST(likes + 1, 0)"))

	resp.Msg = "点赞数据保存成功"
	resp.Ok = true
	return
}
