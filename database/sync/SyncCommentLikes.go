package sync

import (
	"GO1/global"
	"GO1/models/Comment"
)

func SyncCommentLikes() {
	type LikeCount struct {
		CommentId int64
		Likes     int64
	}

	var results []LikeCount
	// 聚合查询
	global.DB.Table("comment_likes").
		Select("comment_id, COUNT(*) as likes").
		Group("comment_id").
		Scan(&results)

	// 遍历更新到 comments 表中
	for _, r := range results {
		global.DB.Model(&Comment.Comment{}).
			Where("id = ?", r.CommentId).
			Update("likes", r.Likes)
	}

	global.Logger.Info("评论点赞数同步完成，共更新", len(results), "条记录")
}
