package images

import (
	"GO1/database/mysql/image"
	"GO1/global"
	models_upload "GO1/models/upload"
	"GO1/pkg/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func UploadImage(c *gin.Context, file []*multipart.FileHeader, rsp *[]models_upload.ResponseUploadImages) {
	dirPath := global.Config.Upload.Image.Path
	imageLimitSize := float64(global.Config.Upload.Image.Size)
	ensureDir(dirPath)

	for _, file := range file {
		filename := file.Filename
		if !inWriteList(filepath.Ext(filename)) {
			addResponse(rsp, filename, "上传失败", "文件格式错误")
			continue
		}

		size := float64(file.Size) / 1024 / 1024
		if size >= imageLimitSize {
			addResponse(rsp, filename, "上传失败", fmt.Sprintf("图片大小为：%dM，不能超过：%dM", file.Size, imageLimitSize))
			continue
		}
		filePath := dirPath + "/" + filename

		fileObj, err := file.Open()
		if err != nil {
			global.Logger.Error(err)
			addResponse(rsp, filename, "上传失败", "解析错误1")
			continue
		}
		byteData, err := io.ReadAll(fileObj)
		if err != nil {
			addResponse(rsp, filename, "上传失败", "解析错误2")
			continue
		}
		md5Str := md5.MD5(byteData)

		if image.CheckImage(md5Str) {
			addResponse(rsp, filename, "上传失败", "该图片已存在")
			continue
		}

		err = image.CreateImage(md5Str, filename, filePath)
		if err != nil {
			addResponse(rsp, filename, "上传失败", "保存失败")
			continue
		}

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			addResponse(rsp, filename, "上传失败", "保存失败")
			continue
		}

		addResponse(rsp, filename, "上传成功", "保存成功")
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

func inWriteList(ext string) bool {
	ext = strings.ToLower(ext)
	for _, EXT := range models_upload.WriteImageList {
		if EXT == ext {
			return true
		}
	}
	return false
}

func addResponse(rsp *[]models_upload.ResponseUploadImages, filename string, result string, msg string) {
	*rsp = append(*rsp, models_upload.ResponseUploadImages{
		FileName: filename,
		Result:   result,
		Msg:      msg,
	})
}
