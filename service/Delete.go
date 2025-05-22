package service

import "GO1/database/mysql"

func DeleteList[T any](idList []int64, t T) error {
	err := mysql.DeleteList(idList, t)
	return err
}
