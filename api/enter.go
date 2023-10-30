package api

import "gin-box/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
