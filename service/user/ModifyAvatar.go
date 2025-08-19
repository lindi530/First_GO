package user

import (
	mysql_image "GO1/database/mysql/image"
	"GO1/database/mysql/user_mysql"
	"GO1/global"
	models_upload "GO1/models/upload"
	"GO1/models/user_model"
	"GO1/pkg/md5"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func ModifyAvatar(c *gin.Context, userId int64, avatar *multipart.FileHeader) gin.H {
	if !AuthUser(c, userId) {
		global.Logger.Error("无权限")
		return gin.H{
			"avatarPath": "",
			"msg":        "无权限",
		}
	}

	if !inWriteList(filepath.Ext(avatar.Filename)) {
		return gin.H{
			"avatarPath": "",
			"msg":        "格式不允许",
		}
	}

	imageLimitSize := float64(global.Config.Upload.Image.Size)
	if float64(avatar.Size/1024/1024) > imageLimitSize {
		return gin.H{
			"avatarPath": "",
			"msg":        "超出限制大小",
		}
	}

	dirPath := global.Config.Upload.Image.Path
	ensureDir(dirPath)

	filePath := dirPath + "/" + avatar.Filename

	fileObj, err := avatar.Open()
	if err != nil {
		return gin.H{
			"avatarPath": "",
			"msg":        "文件解析失败",
		}
	}
	byteData, err := io.ReadAll(fileObj)
	if err != nil {
		return gin.H{
			"avatarPath": "",
			"msg":        "文件解析失败",
		}
	}
	fileMD5 := md5.MD5(byteData)
	if !mysql_image.CheckImage(fileMD5) {
		err = mysql_image.CreateImage(fileMD5, avatar.Filename, filePath)
		if err != nil {
			return gin.H{
				"avatarPath": "",
				"msg":        "文件保存失败",
			}
		}
		err = c.SaveUploadedFile(avatar, filePath)
		if err != nil {
			return gin.H{
				"avatarPath": "",
				"msg":        "文件保存失败",
			}
		}
	}
	user_mysql.ModifyAvatar(userId, fileMD5)
	avatarPath := user_model.GetAvatarPath(fileMD5)
	return gin.H{
		"avatar": avatarPath,
		"msg":    "上传成功",
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
