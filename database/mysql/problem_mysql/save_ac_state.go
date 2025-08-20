package problem_mysql

import (
	"GO1/global"
	"GO1/models/user_ac_problem"
	"gorm.io/gorm/clause"
)

func SaveAcState(userid int64, problemId uint) {
	global.DB.
		Clauses(clause.Insert{Modifier: "IGNORE"}). // 存在就忽略
		Create(&user_ac_problem.UserAcProblem{
			UserId:    userid,
			ProblemId: problemId,
		})
}
