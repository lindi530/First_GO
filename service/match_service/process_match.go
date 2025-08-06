package match_service

import (
	"GO1/global"
	"GO1/models/match_model"
	"sync"
)

var (
	matchLock   sync.Mutex
	waitingUser *match_model.MatchUser = nil
)

func processMatch(user *match_model.MatchUser) { // 匹配逻辑
	matchLock.Lock()
	defer matchLock.Unlock()

	if waitingUser == nil {
		global.Logger.Errorf("User %s waiting for match_model", user.Username)
		waitingUser = user
	} else {
		global.Logger.Errorf("Matched %s vs %s!\n", waitingUser.Username, user.Username)
		startGame(waitingUser, user)
		waitingUser = nil
	}
}
