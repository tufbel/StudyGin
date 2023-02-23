package main

import (
	"StudyGin/src/models"
	"StudyGin/src/tools/myLog"
	"gorm.io/gorm/logger"
)

func main() {

	myLog.Logger.Info("Start migrate models script ---------")
	models.StartupDB(map[string]interface{}{"logLevel": logger.Info})

	{
		models.GormDB.
			//Set("gorm:table_options", "ENGINE=InnoDB").
			Set("gorm:table_options", "CHARSET=utf8mb4").
			Set("gorm:table_options", "COLLATE=utf8mb4_unicode_ci").
			AutoMigrate(&models.User{})
	}

	myLog.Logger.Info("Migrate models ... success")
	defer models.ShutdownDB()
}
