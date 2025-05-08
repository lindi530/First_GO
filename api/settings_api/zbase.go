package settings_api

import (
	"GO1/models/response"
	"github.com/gin-gonic/gin"
)

type SettingsAPI struct{}

func (SettingsAPI) SettingsInfoView(c *gin.Context) {
	response.Ok(map[string]string{}, "Success", c)
	response.OkWithData(map[string]string{}, c)
	response.OkWithMessage("Success", c)

	response.Fail(map[string]string{}, "Fail", c)
	response.FailWithCode(response.SettingsError, c)
	response.FailWithMessage("Fail", c)
}
