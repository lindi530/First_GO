package mapping

import (
	"GO1/global"
	"github.com/gin-gonic/gin"
)

var MappingPath = map[string]string{}

func MappingRouter(router *gin.Engine) {
	AddMappingRouter("/images", global.Config.Upload.Image.Path, router)
}

func AddMappingRouter(relativePath string, root string, router *gin.Engine) {
	MappingPath[root] = relativePath
	router.Static(relativePath, root)
}

func GetRealPath(path string) string {
	return MappingPath[path]
}
