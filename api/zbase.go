package api

import (
	"GO1/api/auth_api"
	"GO1/api/chat_api"
	"GO1/api/comment_api"
	"GO1/api/image_api"
	"GO1/api/match_api"
	"GO1/api/message_api"
	"GO1/api/post_api"
	"GO1/api/problem_api"
	"GO1/api/saber_api"
	"GO1/api/settings_api"
	"GO1/api/upload_api"
	"GO1/api/user_api"
)

type ApiGroup struct {
	SettingsAPI settings_api.SettingsAPI
	UserAPI     user_api.UserAPI
	PostAPI     post_api.PostAPI
	ImageAPI    image_api.ImageAPI
	CommentAPI  comment_api.CommentAPI
	AuthAPI     auth_api.AuthAPI
	ChatAPI     chat_api.ChatAPI
	MessageAPI  message_api.MessageAPI
	ProblemAPI  problem_api.ProblemAPI
	MatchAPI    match_api.MatchAPI
	SaberAPI    saber_api.SaberAPI
	UploadAPI   upload_api.UploadAPI
}

var ApiGroups = new(ApiGroup)
