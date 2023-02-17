package user

import (
	"StudyGin/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserViewSets struct{}

// allUsers
// @Summary      allUsers
// @Description  获取用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.UserModel
// @Failure      400
// @Router       /user/all [get]
func (receiver *UserViewSets) allUsers(ctx *gin.Context) {
	userObjSlice := []models.UserModel{
		{Name: "Tuffy", Age: 18, MyType: "mouse"},
		{Name: "Tom", Age: 20, MyType: "cat"},
	}
	ctx.JSON(http.StatusOK, userObjSlice)
}

// createUser
// @Summary      createUser
// @Description  创建用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user body models.UserModel true "用户信息"
// @Success      200  {object}  models.UserModel
// @Failure      400
// @Router       /user [post]
func (receiver *UserViewSets) createUser(ctx *gin.Context) {
	var userObj models.UserModel
	//if err := ctx.ShouldBind(&userObj); err != nil {
	//	println("err ->", err.Error())
	//	return
	//}
	ctx.Bind(&userObj)
	ctx.JSON(http.StatusOK, userObj)
}

// getUser
// @Summary      getUser
// @Description  获取用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        name path string true "用户ID"
// @Success      200  {object}  models.UserModel
// @Failure      400
// @Router       /user/{name} [get]
func (receiver *UserViewSets) getUser(ctx *gin.Context) {
	name := ctx.Param("name")
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		println("Get id Error ->", err.Error(), ctx.Param("id"), ctx.Param("name"))
	} else {
		println("Get id ->", id)
	}
	userObj := models.UserModel{Name: name, Age: 18, MyType: "mouse"}
	ctx.JSON(http.StatusOK, userObj)
}

// deleteUser
// @Summary      deleteUser
// @Description  删除用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        name path string true "用户ID"
// @Success      200  {string} string
// @Failure      400
// @Router       /user/{name} [delete]
func (receiver *UserViewSets) deleteUser(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.String(http.StatusOK, "delete user "+name)
}

func InitUserGroup(apiGroup *gin.RouterGroup) {
	var userViewSetObj UserViewSets
	user := apiGroup.Group("/user")
	{
		user.GET("/all", userViewSetObj.allUsers)
		user.POST("", userViewSetObj.createUser)
		user.GET("/:name", userViewSetObj.getUser)
		user.DELETE("/:name", userViewSetObj.deleteUser)
	}
}
