package post

import (
	"GO1/database/mysql/post_likes"
	models_post "GO1/models/post"
	"GO1/models/user_model"
	"github.com/gin-gonic/gin"
)

func BuildPostResponse(c *gin.Context, userId int64, p models_post.Post) models_post.PostResponse {
	return models_post.PostResponse{
		PostID:    p.PostID,
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		Views:     p.Views,
		Likes:     p.Likes,
		Like:      post_likes.PostLikeCheck(userId, p.PostID),
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Author: models_post.AuthorInfo{
			UserId:   p.Author.UserID,
			UserName: p.Author.UserName,
			Avatar:   user_model.GetAvatarPath(c, p.Author.Avatar),
		},
	}
}

func BuildPostsResponse(c *gin.Context, userId int64, p []models_post.Post) []models_post.PostResponse {
	postsResponse := make([]models_post.PostResponse, 0, len(p))
	for _, post := range p {
		postsResponse = append(postsResponse, BuildPostResponse(c, userId, post))
	}
	return postsResponse
}
