package settings_api

import "github.com/gin-gonic/gin"

func (SettingsAPI) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "settings_info_view",
	})
}
