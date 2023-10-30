package global

import (
	"gin-box/conf"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// Config 用于保存配置文件
	Config *conf.Config
	// DB 用于保存数据库实例
	DB *gorm.DB
	// Log 全局Log
	Log *logrus.Logger
)
