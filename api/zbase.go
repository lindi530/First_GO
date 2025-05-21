package api

import (
	"GO1/api/image_api"
	"GO1/api/post_api"
	"GO1/api/settings_api"
	"GO1/api/user_api"
)

type ApiGroup struct {
	SettingsAPI settings_api.SettingsAPI
	UserAPI     user_api.UserAPI
	PostAPI     post_api.PostAPI
	ImageAPI    image_api.ImageAPI
}

var ApiGroups = new(ApiGroup)
