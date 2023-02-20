package main

import (
	"StudyGin/src/models"
	"StudyGin/src/tools"
)

func main() {
	tools.Logger.Info("Run migrate models script ---------")
	models.StartupDB()
	models.GormDB.AutoMigrate(&models.User{})

	defer models.ShutdownDB()
}
