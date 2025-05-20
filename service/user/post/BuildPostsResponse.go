package post

import (
	"GO1/models"
)

func BuildPostsResponse(p []models.Post) []models.PostResponse {
	postsResponse := make([]models.PostResponse, 0, len(p))
	for _, post := range p {
		postsResponse = append(postsResponse, models.BuildPostResponse(post))
	}
	return postsResponse
}
