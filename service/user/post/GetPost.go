package post

import (
	mysql_post "GO1/database/mysql/post"
	mysql_user "GO1/database/mysql/user"
	"GO1/models"
)

func GetUserPosts(userID int64) ([]models.PostResponse, error) {

	posts, err := mysql_user.GetUserPost(userID)
	if err != nil {
		return nil, err
	}
	postResponse := BuildPostsResponse(posts)

	return postResponse, err
}

func GetAllPost() ([]models.PostResponse, error) {
	posts, err := mysql_post.GetAllPost()
	if err != nil {
		return nil, err
	}
	postResponse := BuildPostsResponse(posts)
	return postResponse, err
}

func GetPostByPostId(postId int64) models.Post {
	post := mysql_post.GetPostByPostId(postId)
	return post
}
