package comment_mysql

import (
	"GO1/global"
	"GO1/models"
	"GO1/models/Comment"
)

func CreateComment(comment *Comment.Comment) models.HandleFuncResp {
	if err := global.DB.Create(comment).Error; err != nil {
		return models.HandleFuncResp{
			Ok:  false,
			Msg: "上传失败",
		}
	}
	return models.HandleFuncResp{
		Ok:  true,
		Msg: "上传成功",
	}
}
