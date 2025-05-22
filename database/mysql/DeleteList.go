package mysql

import "GO1/global"

func DeleteList[T any](idList []int64, t T) error {
	var deleteRes []T
	err := global.DB.Find(&deleteRes, idList).Error
	if err != nil {
		return err
	}
	err = global.DB.Delete(deleteRes).Error
	return err
}
