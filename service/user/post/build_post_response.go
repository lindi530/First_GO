package post

import (
	"GO1/models"
)

func BuildPostResponse(p models.Post) models.PostResponse {
	return models.PostResponse{
		PostID:    p.PostID,
		UserID:    p.UserID,
		Title:     p.Title,
		Content:   p.Content,
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Author: models.AuthorInfo{
			UserName: p.Author.UserName,
			Email:    p.Author.Email,
		},
	}
}

func BuildPostsResponse(p []models.Post) []models.PostResponse {

	postsResponse := make([]models.PostResponse, 0, len(p))
	for _, post := range p {
		postsResponse = append(postsResponse, BuildPostResponse(post))
	}
	return postsResponse
}
