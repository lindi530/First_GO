package post

import (
	mysql_post "GO1/database/mysql/post"
	mysql_user "GO1/database/mysql/user"
	"GO1/global"
	"GO1/models"
	"fmt"
)

func GetUserPosts(userID int64) ([]models.PostResponse, error) {

	posts, err := mysql_user.GetUserPost(userID)
	if err != nil {
		return nil, err
	}
	postResponse := models.BuildPostsResponse(posts)

	return postResponse, err
}

func GetAllPost() ([]models.PostResponse, error) {
	posts, err := mysql_post.GetAllPost()
	if err != nil {
		return nil, err
	}
	postResponse := models.BuildPostsResponse(posts)
	return postResponse, err
}

func GetPostByPostId(postId int64) models.Post {
	post := mysql_post.GetPostByPostId(postId)
	return post
}

func GetThePagePost(info models.PageInfo) (listRep []models.PostResponse, err error) {
	limit := info.Limit
	offset := (info.Page - 1) * info.Limit
	order := "created_at DESC"
	fmt.Println(limit, offset, order)

	list := []models.Post{}
	err = global.DB.Limit(limit).Offset(offset).Order(order).Find(&list).Error
	if err == nil {
		listRep = models.BuildPostsResponse(list)
	}
	return listRep, err
}
