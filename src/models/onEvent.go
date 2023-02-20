package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	GormDB *gorm.DB
)

func StartupDB() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,  // DSN data source name
		DefaultStringSize:        256,  // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:   true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:  true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}

	poolDB, err := gormDB.DB()
	if err != nil {
		panic("failed to get poolDB on init:" + err.Error())
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	poolDB.SetMaxIdleConns(1)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	poolDB.SetMaxOpenConns(4)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	poolDB.SetConnMaxLifetime(time.Minute)
}

func ShutdownDB() {
	sqlDB, err := GormDB.DB()
	if err != nil {
		println("failed to get sqlDB on close: " + err.Error())
		return
	}
	sqlDB.Close()
}
