package user

import (
	"StudyGin/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserViewSets struct{}

func (receiver *UserViewSets) GetUser(ctx *gin.Context) {

}

// CreateUser
// @Summary      CreateUser
// @Description  创建用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user body models.UserModel true "用户信息"
// @Success      200  {object}  models.UserModel
// @Failure      400
// @Router       /user [post]
func (receiver *UserViewSets) CreateUser(ctx *gin.Context) {
	var userObj models.UserModel
	//if err := ctx.ShouldBind(&userObj); err != nil {
	//	println("err ->", err.Error())
	//	return
	//}
	ctx.Bind(&userObj)
	ctx.JSON(http.StatusOK, userObj)
}

func InitUserGroup(apiGroup *gin.RouterGroup) {
	var userViewSets UserViewSets
	user := apiGroup.Group("/user")
	user.POST("", userViewSets.CreateUser)
}
