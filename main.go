package main

import (
	_ "StudyGin/docs"
	"StudyGin/src"
	"StudyGin/src/models"
	"StudyGin/src/project_settings"
	"StudyGin/src/tools/myLog"
	swagFiles "github.com/swaggo/files"
	gSwag "github.com/swaggo/gin-swagger"
	"gorm.io/gorm/logger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server

// @host      localhost:20088
// @BasePath  /api/study_gin

// @schemes http https
func main() {
	{
		myLog.InitLogger()
	}
	webRouter := src.InitRouter()

	// 添加swagger文档
	{
		//docs.SwaggerInfo.Title = "StudyGin Example API"
		//docs.SwaggerInfo.Description = "This is a sample server."
		//docs.SwaggerInfo.Version = "1.0"
		//docs.SwaggerInfo.Host = "localhost:20088"
		//docs.SwaggerInfo.BasePath = "/api/study_gin"
		//docs.SwaggerInfo.Schemes = []string{"http", "https"}
		webRouter.GET(project_settings.RootURL+"/docs/*any", gSwag.WrapHandler(swagFiles.Handler))
		myLog.Logger.Debug("Docs: http://localhost:20088" + project_settings.RootURL + "/docs/index.html")
	}

	// startup
	{
		models.StartupDB(map[string]interface{}{"logLevel": logger.Info})
	}

	// Listen and Server in 0.0.0.0:20088
	webRouter.Run(":20088")

	// shutdown
	{
		defer models.ShutdownDB()
	}
}
