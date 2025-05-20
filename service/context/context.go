package context

import "github.com/gin-gonic/gin"

const CtxUserIdKey = "user_id"

func SetContext(c *gin.Context, key string, value interface{}) {
	c.Set(key, value)
}

func GetContext(c *gin.Context, key string) any {
	data, _ := c.Get(key)
	return data
}
