package post

import (
	mysql_post "GO1/database/mysql/post"
	"GO1/database/mysql/user_mysql"
	"GO1/global"
	"GO1/models/post"
	"fmt"
)

func GetUserPosts(userId int64) ([]post.PostResponse, error) {

	posts, err := user_mysql.GetUserPost(userId)
	if err != nil {
		return nil, err
	}
	postResponse := BuildPostsResponse(userId, posts)

	return postResponse, err
}

func GetAllPost(userId int64) ([]post.PostResponse, error) {
	posts, err := mysql_post.GetAllPost()
	if err != nil {
		return nil, err
	}
	postResponse := BuildPostsResponse(userId, posts)
	return postResponse, err
}

func GetPostByPostId(postId int64) post.Post {
	post := mysql_post.GetPostByPostId(postId)
	return post
}

func GetThePagePost(userId int64, info post.PageInfo) (listRep []post.PostResponse, err error) {
	limit := info.Limit
	offset := (info.Page - 1) * info.Limit
	order := "created_at DESC"
	fmt.Println(limit, offset, order)

	list := []post.Post{}
	err = global.DB.Limit(limit).Offset(offset).Order(order).Find(&list).Error
	if err == nil {
		listRep = BuildPostsResponse(userId, list)
	}
	return listRep, err
}
