package core

import (
	"gin-box/global"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitGorm 数据库初始化
func InitGorm() *gorm.DB {
	if global.Config.Pgsql.Host == "" {
		global.Log.Warnln("未配置pgsql，取消gorm连接！")
	}
	dsn := global.Config.Pgsql.Dsn()
	var pgsqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		// 开发环境显示所有sql
		pgsqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		pgsqlLogger = logger.Default.LogMode(logger.Error)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: pgsqlLogger,
	})
	if err != nil {
		global.Log.Fatalf("pgsql连接失败：%s", dsn)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
