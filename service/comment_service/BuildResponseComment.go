package comment_service

import (
	"GO1/models/Comment"
	"GO1/models/user_model"
)

func BuildResponseComment(comment *Comment.Comment, author *user_model.AuthorInfo, like bool) *Comment.ResponseComment {
	return &Comment.ResponseComment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		Likes:     comment.Likes,
		Like:      like,
		Author: user_model.AuthorInfo{
			UserID:   author.UserID,
			UserName: author.UserName,
			Avatar:   user_model.GetAvatarPath(author.Avatar),
		},
	}
}
