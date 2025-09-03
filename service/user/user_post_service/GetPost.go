package user_post_service

import (
	mysql_post "GO1/database/mysql/post_mysql"
	"GO1/database/mysql/user_mysql"
	"GO1/global"
	"GO1/models/post_model"
	"fmt"
)

func GetUserPosts(userId int64) ([]post_model.PostResponse, error) {

	posts, err := user_mysql.GetUserPost(userId)
	if err != nil {
		return nil, err
	}
	postResponse := BuildPostsResponse(userId, posts)

	return postResponse, err
}

func GetAllPost(userId int64) ([]post_model.PostResponse, error) {
	posts, err := mysql_post.GetAllPost()
	if err != nil {
		return nil, err
	}
	postResponse := BuildPostsResponse(userId, posts)
	return postResponse, err
}

//func GetPostByPostId(userId, postId int64) post_model.Post {
//	var post post_model.Post
//	err := mysql_post.GetPostByPostId(userId, postId, &post)
//	if err != nil {
//	}
//	return post
//}

func GetThePagePost(userId int64, info post_model.PageInfo) (listRep []post_model.PostResponse, err error) {
	limit := info.Limit
	offset := (info.Page - 1) * info.Limit
	order := "created_at DESC"
	fmt.Println(limit, offset, order)

	list := []post_model.Post{}
	err = global.DB.Limit(limit).Offset(offset).Order(order).Find(&list).Error
	if err == nil {
		listRep = BuildPostsResponse(userId, list)
	}
	return listRep, err
}
