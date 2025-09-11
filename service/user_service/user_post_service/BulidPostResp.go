package user_post_service

import (
	"GO1/database/mysql/post_likes"
	models_post "GO1/models/post_model"
	"GO1/models/user_model"
)

func BuildPostResponse(userId int64, p models_post.Post) models_post.PostResponse {
	return models_post.PostResponse{
		PostID:    p.PostID,
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		Views:     p.Views,
		Likes:     p.Likes,
		Liked:     post_likes.PostLikeCheck(userId, p.PostID),
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Author: models_post.AuthorInfo{
			UserId:   p.Author.UserID,
			UserName: p.Author.UserName,
			Avatar:   user_model.GetAvatarPath(p.Author.Avatar),
		},
	}
}

func BuildPostsResponse(userId int64, posts []models_post.Post) []models_post.PostResponse {
	respPosts := make([]models_post.PostResponse, len(posts)) // 预分配长度
	for i, post := range posts {
		respPosts[i] = BuildPostResponse(userId, post)
	}
	return respPosts
}
