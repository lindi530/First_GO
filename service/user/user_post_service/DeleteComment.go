package user_post_service

import (
	"GO1/database/mysql/comment_mysql"
	"GO1/models"
	"GO1/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func DeleteComment(c *gin.Context, commentId int64) models.HandleFuncResp {
	userId := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	comment := comment_mysql.CheckCommentById(commentId)
	if comment.ID == 0 {
		return models.HandleFuncResp{
			Msg: "该评论不存在",
			Ok:  false,
		}
	}
	if comment.AuthorID != userId {
		return models.HandleFuncResp{
			Msg: "无权限删除该评论",
			Ok:  false,
		}
	}

	comment_mysql.DeleteComment(comment)

	return models.HandleFuncResp{
		Msg: "评论删除成功",
		Ok:  true,
	}
}
