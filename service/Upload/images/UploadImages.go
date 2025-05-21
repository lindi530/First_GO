package images

import (
	"GO1/global"
	models "GO1/models/Upload"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
)

func UploadImage(c *gin.Context, file []*multipart.FileHeader, rsp *[]*models.ResponseUploadImages) {
	dirPath := global.Config.Upload.Image.Path
	imageLimitSize := float64(global.Config.Upload.Image.Size)
	ensureDir(dirPath)

	for _, file := range file {
		size := float64(file.Size) / 1024 / 1024
		if size >= imageLimitSize {
			*rsp = append(*rsp, &models.ResponseUploadImages{
				FileName: file.Filename,
				Result:   "上传失败",
				Msg:      fmt.Sprintf("图片大小为：%dM，不能超过：%dM", file.Size, imageLimitSize),
			})
			continue
		}
		filePath := dirPath + "/" + file.Filename
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			*rsp = append(*rsp, &models.ResponseUploadImages{
				FileName: file.Filename,
				Result:   "保存失败",
				Msg:      "保存失败",
			})
			continue
		}
		*rsp = append(*rsp, &models.ResponseUploadImages{
			FileName: file.Filename,
			Result:   "上传成功",
			Msg:      "保存成功",
		})
	}
}

func ensureDir(dirPath string) {
	// 检查路径是否存在
	_, err := os.Stat(dirPath)
	if err == nil {
		return
	}
	// 如果错误是因为路径不存在
	if os.IsNotExist(err) {
		// 创建目录（包括所有必要的父目录）
		os.MkdirAll(dirPath, 0755)
	}
	// 其他类型的错误
}
