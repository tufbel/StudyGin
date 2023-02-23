package user

import (
	"StudyGin/src/models"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"net/http"
)

type UserViewSets struct{}

// allUsers
// @Summary      allUsers
// @Description  获取用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.User
// @Failure      400
// @Router       /user/all [get]
func (receiver *UserViewSets) allUsers(ctx *gin.Context) {
	var userObjSlice []models.User
	result := models.GormDB.Find(&userObjSlice)
	if result.Error != nil {
		panic(result.Error)
	}
	ctx.JSON(http.StatusOK, userObjSlice)
}

// createUser
// @Summary      createUser
// @Description  创建用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user body models.CreateUserSchema true "用户信息"
// @Success      200  {object}  models.User
// @Failure      400
// @Router       /user [post]
func (receiver *UserViewSets) createUser(ctx *gin.Context) {
	var userForm_ models.CreateUserSchema
	if validationErr := ctx.ShouldBind(&userForm_); validationErr != nil {
		panic(validationErr)
	}
	userObj := models.User{
		Uid:      ksuid.New().String(),
		Name:     userForm_.Name,
		Email:    userForm_.Email,
		Password: userForm_.Password,
	}
	if result_ := models.GormDB.Create(&userObj); result_.Error != nil {
		panic(result_.Error)
	}
	ctx.JSON(http.StatusOK, userObj)
}

// getUser
// @Summary      getUser
// @Description  获取用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uid path string true "用户ID"
// @Success      200  {object}  models.User
// @Failure      400
// @Router       /user/{uid} [get]
func (receiver *UserViewSets) getUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	userObj := models.User{Uid: uid}
	if result_ := models.GormDB.First(&userObj); result_.Error != nil {
		panic(result_.Error)
	}
	ctx.JSON(http.StatusOK, userObj)
}

// updateUser
// @Summary      updateUser
// @Description  更新用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uid path string true "用户ID"
// @Param        user body models.UpdateUserSchema true "用户信息"
// @Success      200  {object}  models.User
// @Failure      400
// @Router       /user/{uid} [patch]
func (receiver *UserViewSets) updateUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	var userForm models.UpdateUserSchema
	if validationErr := ctx.ShouldBind(&userForm); validationErr != nil {
		panic(validationErr)
	}
	userObj := models.User{Uid: uid}
	if result_ := models.GormDB.First(&userObj); result_.Error != nil {
		panic(result_.Error)
	}
	updateUser := models.User{
		Name:     userForm.Name,
		Email:    userForm.Email,
		Password: userForm.Password,
	}
	if result_ := models.GormDB.Model(&userObj).Updates(updateUser); result_.Error != nil {
		panic(result_.Error)
	}
	ctx.JSON(http.StatusOK, userObj)
}

// deleteUser
// @Summary      deleteUser
// @Description  删除用户
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        uid path string true "用户ID"
// @Success      200  {string} string
// @Failure      400
// @Router       /user/{uid} [delete]
func (receiver *UserViewSets) deleteUser(ctx *gin.Context) {
	uid := ctx.Param("uid")
	userObj := models.User{Uid: uid}
	if result_ := models.GormDB.First(&userObj); result_.Error != nil {
		panic(result_.Error)
	}
	models.GormDB.Delete(&userObj)
	ctx.JSON(http.StatusOK, userObj)
}

func InitUserGroup(apiGroup *gin.RouterGroup) {
	var userViewSetObj UserViewSets
	user := apiGroup.Group("/user")
	{
		user.GET("/all", userViewSetObj.allUsers)
		user.POST("", userViewSetObj.createUser)
		user.GET("/:uid", userViewSetObj.getUser)
		user.PATCH("/:uid", userViewSetObj.updateUser)
		user.DELETE("/:uid", userViewSetObj.deleteUser)
	}
}
