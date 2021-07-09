package global

import (
	"go-service/pkg/logger"
	"go-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettings
	DatabaseSetting *setting.DatabaseSettings
	Logger          *logger.Logger
)
