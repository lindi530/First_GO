package problem_mysql

import (
	"GO1/global"
	"GO1/models/user_ac_problem"
)

func SaveAcState(userID int64, problemID uint) {
	global.Logger.Info("SaveAcState:", userID, problemID)
	userAcProblemState := user_ac_problem.UserAcProblem{
		UserId:    userID,
		ProblemId: problemID,
	}
	global.DB.Create(&userAcProblemState)
}
