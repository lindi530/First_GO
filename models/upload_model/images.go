package upload_model

import (
	"GO1/global"
	"gorm.io/gorm"
	"os"
	"time"
)

type Image struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	MD5       string    `json:"md5"`
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
type ResponseUploadImages struct {
	FileName string
	Result   string
	Msg      string
}

var WriteImageList = []string{
	".jpg",
	".png",
	".gif",
}

func (I *Image) BeforeDelete(db *gorm.DB) error {
	err := os.Remove(I.Path)
	if err != nil {
		global.Logger.Error("删除失败！")
		return err
	}
	return nil
}
