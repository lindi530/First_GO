package service

import (
	mysql_image "GO1/database/mysql/image"
	mysql_user "GO1/database/mysql/user"
	"GO1/global"
	models_upload "GO1/models/upload"
	"GO1/pkg/md5"
	"GO1/service/user"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func ModifyAvatar(c *gin.Context, userId int64, avatar *multipart.FileHeader) error {

	if !user.CheckUserId(c, userId) {
		global.Logger.Error("无权限")
		return errors.New("无权限")
	}

	if !inWriteList(filepath.Ext(avatar.Filename)) {
		return errors.New("格式不允许")
	}

	imageLimitSize := float64(global.Config.Upload.Image.Size)
	if float64(avatar.Size/1024/1024) > imageLimitSize {
		return errors.New("超出限制大小")
	}

	dirPath := global.Config.Upload.Image.Path
	ensureDir(dirPath)

	filePath := dirPath + "/" + avatar.Filename

	fileObj, err := avatar.Open()
	if err != nil {
		return errors.New("文件解析失败")
	}
	byteData, err := io.ReadAll(fileObj)
	if err != nil {
		return errors.New("文件解析失败")
	}
	md5Str := md5.MD5(byteData)
	if !mysql_image.CheckImage(md5Str) {
		err = mysql_image.CreateImage(md5Str, avatar.Filename, filePath)
		if err != nil {
			return errors.New("文件保存失败")
		}
		err = c.SaveUploadedFile(avatar, filePath)
		if err != nil {
			return errors.New("文件保存失败")
		}
	}

	mysql_user.ModifyAvatar(userId, md5Str)
	return nil
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
