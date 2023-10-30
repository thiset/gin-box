package settings_api

import (
	"gin-box/models/res"

	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(res.SystemError, c)
}
