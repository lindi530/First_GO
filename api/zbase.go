package api

import (
	"GO1/api/settings_api"
	"GO1/api/users_api"
)

type ApiGroup struct {
	SettingsAPI settings_api.SettingsAPI
	UsersAPI    users_api.UsersAPI
}

var ApiGroups = new(ApiGroup)
