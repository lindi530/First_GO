package sync

import (
	"GO1/global"
	"GO1/models/post_model"
)

func SyncPostLikes() {
	type LikeCount struct {
		PostId int64
		Likes  int64
	}

	var results []LikeCount
	// 聚合查询
	global.DB.Table("post_likes").
		Select("post_id, COUNT(*) as likes").
		Group("post_id").
		Scan(&results)

	// 遍历更新到 comments 表中
	for _, r := range results {
		global.DB.Model(&post_model.Post{}).
			Where("id = ?", r.PostId).
			UpdateColumn("likes", r.Likes)
	}

	global.Logger.Info("帖子点赞数同步完成，共更新", len(results), "条记录")
}
