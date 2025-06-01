package post

import (
	mysql_post "GO1/database/mysql/post"
	mysql_user "GO1/database/mysql/user"
	"GO1/global"
	"GO1/models/post"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserPosts(c *gin.Context, userID int64) ([]post.PostResponse, error) {

	posts, err := mysql_user.GetUserPost(userID)
	if err != nil {
		return nil, err
	}
	postResponse := post.BuildPostsResponse(c, posts)

	return postResponse, err
}

func GetAllPost(c *gin.Context) ([]post.PostResponse, error) {
	posts, err := mysql_post.GetAllPost()
	if err != nil {
		return nil, err
	}
	postResponse := post.BuildPostsResponse(c, posts)
	return postResponse, err
}

func GetPostByPostId(postId int64) post.Post {
	post := mysql_post.GetPostByPostId(postId)
	return post
}

func GetThePagePost(c *gin.Context, info post.PageInfo) (listRep []post.PostResponse, err error) {
	limit := info.Limit
	offset := (info.Page - 1) * info.Limit
	order := "created_at DESC"
	fmt.Println(limit, offset, order)

	list := []post.Post{}
	err = global.DB.Limit(limit).Offset(offset).Order(order).Find(&list).Error
	if err == nil {
		listRep = post.BuildPostsResponse(c, list)
	}
	return listRep, err
}
