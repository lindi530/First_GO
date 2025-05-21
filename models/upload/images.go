package upload

import "time"

type Image struct {
	Id        int       `json:"id" gorm:"primaryKey"`
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
