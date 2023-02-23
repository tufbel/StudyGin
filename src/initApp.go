package src

import (
	"StudyGin/src/projectSettings"
	"StudyGin/src/routeGroups/index"
	"StudyGin/src/routeGroups/middleware"
	"StudyGin/src/routeGroups/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() (rootGin *gin.Engine) {
	rootGin = gin.New()
	rootGin.Use(gin.Logger(), gin.Recovery())

	apiGroup := rootGin.Group(projectSettings.RootURL)

	apiGroup.Use(middleware.GormMiddleware())

	index.InitIndexGroup(apiGroup)
	user.InitUserGroup(apiGroup)

	rootGin.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(301, projectSettings.RootURL) // 301永久重定向
		//ctx.Request.URL.Path = projectSettings.RootURL // 路由重定向
		//rootGin.HandleContext(ctx)
	})

	return rootGin
}
