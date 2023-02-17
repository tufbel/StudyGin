package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// studyGinHome
// @Summary      home
// @Description  验活接口
// @Tags         index
// @Accept       json
// @Produce      json
// @Success      200  {string}  success
// @Failure      400
// @Router       / [get]
func studyGinHome(ctx *gin.Context) {
	ctx.String(http.StatusOK, "success")
}

func InitIndexGroup(apiGroup *gin.RouterGroup) {
	index := apiGroup.Group("/")
	index.GET("", studyGinHome)
}
