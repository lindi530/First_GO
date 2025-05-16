package post

import (
	mysql "GO1/database/mysql/user"
	"GO1/models"
)

func GetPosts(userID int64) ([]models.PostResponse, error) {

	posts, err := mysql.GetPost(userID)
	if err != nil {
		return nil, err
	}
	postResponse := BuildPostsResponse(posts)

	return postResponse, err
}
