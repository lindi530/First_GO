package image

import (
	"GO1/global"
	"GO1/models/upload_model"
	"time"
)

func CheckImage(imgMD5Str string) bool {
	var image upload_model.Image
	global.DB.Where("md5 = ?", imgMD5Str).First(&image)
	return image.Id != 0
}

func CreateImage(imgMD5Str string, fileName string, filePath string) error {
	time := time.Now()
	image := upload_model.Image{
		MD5:       imgMD5Str,
		Name:      fileName,
		Path:      filePath,
		CreatedAt: time,
	}
	err := global.DB.Create(&image)
	return err.Error
}
