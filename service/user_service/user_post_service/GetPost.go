package user_post_service

import (
	mysql_post "GO1/database/mysql/post_mysql"
	"GO1/database/mysql/user_mysql"
	"GO1/global"
	"GO1/middlewares/response"
	"GO1/models/post_model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserPosts(userId int64, page, pageSize int) (resp response.Response) {

	var posts []post_model.Post
	err := user_mysql.GetUserPost(userId, page, pageSize, &posts)
	if err != nil {
		resp.Code = 1
		resp.Message = "查询用户帖子失败"
		return
	}
	postResponse := BuildPostsResponse(userId, posts)

	var count int64
	err = mysql_post.GetUserPostCount(userId, &count)
	if err != nil {
		resp.Code = 1
		resp.Message = "获取用户帖子数量失败"
		return
	}
	resp.Code = 0
	resp.Data = gin.H{
		"posts": postResponse,
		"total": count,
	}

	return
}

func GetAllPost(userId int64, page, pageSize int) (resp response.Response) {
	offset := (page - 1) * pageSize
	limit := pageSize
	var posts []post_model.Post
	err := mysql_post.GetAllPost(&posts, offset, limit)
	if err != nil {
		resp.Code = 1
		resp.Message = "查询帖子信息失败"
		return
	}

	var Count int64
	mysql_post.GetAllPostCount(&Count)

	postResponse := BuildPostsResponse(userId, posts)
	resp.Code = 0
	resp.Data = gin.H{
		"list":  postResponse,
		"total": Count,
	}
	return
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
