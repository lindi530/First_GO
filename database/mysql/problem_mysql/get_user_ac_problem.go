package problem_mysql

import "GO1/global"

func GetUserAcProblemID(userID int64) []uint {
	var problemIDs []uint
	err := global.DB.
		Table("user_ac_problems").
		Where("user_id = ?", userID).
		Pluck("problem_id", &problemIDs).Error
	if err != nil {
		return nil
	}
	return problemIDs
}
