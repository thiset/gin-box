package router

import (
	"gin-box/api"
)

func (r RouterGroup) settingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	r.GET("/settings", settingsApi.SettingsInfoView)
}
