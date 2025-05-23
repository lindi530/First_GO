package routers

import (
	"GO1/global"
	service_router "GO1/mapping"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)

	router := gin.New()
	service_router.MappingRouter(router)
	router.RemoveExtraSlash = true
	router.Use(func(c *gin.Context) {
		fmt.Printf(">> CORS applied on %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // 允许前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	AddRouter(router.Group(""))
	router.NoRoute(func(c *gin.Context) {
		global.Logger.Infof("404 <UNK>")
		c.JSON(404, gin.H{
			"code":    404,
			"message": "Not Found",
		})
	})
	return router
}

func AddRouter(api *gin.RouterGroup) {
	SettingsRouter(api)
	UsersRouter(api)
	PostsRouter(api)
	ImagesRouter(api)
}
