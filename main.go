package main

import (
	"StudyGin/docs"
	"StudyGin/src"
	"StudyGin/src/models"
	"StudyGin/src/projectSettings"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	webRouter := src.InitRouter()

	// 添加swagger文档
	{
		docs.SwaggerInfo.Title = "StudyGin Example API"
		docs.SwaggerInfo.Description = "This is a sample server."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "localhost:20088"
		docs.SwaggerInfo.BasePath = "/api/study_gin"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		webRouter.GET(projectSettings.RootURL+"/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		fmt.Printf("Docs: http://localhost:20088%s", projectSettings.RootURL+"/docs/index.html")
	}

	// startup
	{
		models.StartupDB()
	}

	// Listen and Server in 0.0.0.0:20088
	webRouter.Run(":20088")

	// shutdown
	{
		defer models.StartupDB()
	}
}

func main() {

}