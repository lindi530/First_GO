package user_post_service

import (
	"GO1/database/mysql/comment_like"
	"GO1/database/mysql/comment_mysql"
	"GO1/database/mysql/user_mysql"
	"GO1/models"
	"GO1/models/Comment"
	"GO1/models/user_model"
	"GO1/service/comment_service"
	"github.com/gin-gonic/gin"
)

func GetComments(userId, postId int64) models.HandleFuncResp {
	var comments []Comment.Comment
	comment_mysql.GetComments(&comments, postId)

	authorIDs := make([]int64, len(comments))
	for i, comment := range comments {
		authorIDs[i] = comment.AuthorID
	}

	var authors []user_model.AuthorInfo
	err := user_mysql.FindAuthorsInfo(&authors, authorIDs)
	if err != nil {
		return models.HandleFuncResp{
			Ok: false,
		}
	}
	authorMap := make(map[int64]user_model.AuthorInfo, len(authors))
	for _, author := range authors {
		authorMap[author.UserID] = author
	}

	var userLikeCommentIDs []int64
	var commentIDs []int64
	for _, comment := range comments {
		commentIDs = append(commentIDs, comment.ID)
	}
	err = comment_like.GetUserLikeCommentIDs(&userLikeCommentIDs, &commentIDs, userId)
	if err != nil {
		return models.HandleFuncResp{
			Ok: false,
		}
	}
	userLikeCommentsMap := make(map[int64]bool, len(comments))
	for _, ID := range userLikeCommentIDs {
		userLikeCommentsMap[ID] = true
	}

	responseComments := make([]*Comment.ResponseComment, len(comments))
	for i, comment := range comments {
		author := authorMap[comment.AuthorID]
		res := comment_service.BuildResponseComment(&comments[i], &author, userLikeCommentsMap[comment.ID])
		responseComments[i] = res
	}

	//
	////authors := user_mysql.FindAuthorInfo(comment_mysql.AuthorID)
	//
	//for idx, comment_mysql := range comments {
	//	author := user_mysql.FindAuthorInfo(comment_mysql.AuthorID)
	//res := comment_service.BuildResponseComment(userId, &comments[idx], author)
	//	responseComments[idx] = res
	//}

	return models.HandleFuncResp{
		Data: gin.H{
			"comments": responseComments,
			"length":   len(responseComments),
		},
		Ok: true,
	}
}
