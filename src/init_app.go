package src

import (
	"StudyGin/src/project_settings"
	"StudyGin/src/route_groups/index"
	"StudyGin/src/route_groups/middleware"
	"StudyGin/src/route_groups/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() (rootGin *gin.Engine) {
	rootGin = gin.New()
	rootGin.Use(gin.Logger(), gin.Recovery())

	apiGroup := rootGin.Group(project_settings.RootURL)

	apiGroup.Use(middleware.ExceptionCaptureMiddleware)

	index.InitIndexGroup(apiGroup)
	user.InitUserGroup(apiGroup)

	rootGin.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(301, project_settings.RootURL) // 301永久重定向
		//ctx.Request.URL.Path = project_settings.RootURL // 路由重定向
		//rootGin.HandleContext(ctx)
	})

	return rootGin
}
