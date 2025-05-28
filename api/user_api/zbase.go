package user_api

import (
	"GO1/api/user_api/account"
	"GO1/api/user_api/follow"
)

type UserAPI struct {
	UserAccountAPI account.UserAccountAPI
	UserFollowAPI  follow.UserFollowAPI
}
