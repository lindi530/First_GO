package settings_api

import (
	response2 "GO1/middlewares/response"
	"github.com/gin-gonic/gin"
)

type SettingsAPI struct{}

func (SettingsAPI) SettingsInfoView(c *gin.Context) {
	response2.Ok(map[string]string{}, "Success", c)
	response2.OkWithData(map[string]string{}, c)
	response2.OkWithMessage("Success", c)

	response2.Fail(map[string]string{}, "Fail", c)
	response2.FailWithCode(response2.SettingsError, c)
	response2.FailWithMessage("Fail", c)
}
