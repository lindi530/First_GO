package cors

import (
	"fmt"
	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CorsMiddleware 处理跨域请求
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(">>> CorsMiddleware called on", c.Request.Method, c.Request.URL.Path)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080") // 允许访问的前端地址，开发时localhost:8080
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// 预检请求直接返回
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusNoContent)
			c.Abort()
			return
		}

		c.Next()
	}
}
