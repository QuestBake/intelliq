package controller

import (
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AddNewUser adds new user
func AddNewUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AddNewUser(&user)
	ctx.JSON(http.StatusOK, res)
}

//UpdateUserProfile updates user profile
func UpdateUserProfile(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateUser(&user)
	ctx.JSON(http.StatusOK, res)
}

//ListAllSchoolAdmins fetches all users with role schooladmin
func ListAllSchoolAdmins(ctx *gin.Context) {
	groupID := ctx.Param("groupId")
	res := service.FetchAllSchoolAdmins(groupID)
	ctx.JSON(http.StatusOK, res)
}

//ListAllTeachers fetches all users within school
func ListAllTeachers(ctx *gin.Context) {
	schoolID := ctx.Param("schoolId")
	res := service.FetchAllTeachers(schoolID)
	ctx.JSON(http.StatusOK, res)
}

//ListSelectedTeachers fetches all users within school for specific role (teacher/apporver)
func ListSelectedTeachers(ctx *gin.Context) {
	schoolID := ctx.Param("schoolId")
	roleType := ctx.Param("roleType")
	res := service.FetchSelectedTeachers(schoolID, roleType)
	ctx.JSON(http.StatusOK, res)
}

//TransferRole transfers role from one user to another
func TransferRole(ctx *gin.Context) {
	roleType := ctx.Param("roleType")
	fromUserID := ctx.Param("fromUser")
	toUserID := ctx.Param("toUser")
	if len(roleType) == 0 || len(fromUserID) == 0 || len(toUserID) == 0 {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.TransferUserRole(roleType, fromUserID, toUserID)
	ctx.JSON(http.StatusOK, res)
}

//RemoveUserFromSchool removes user from current school
func RemoveUserFromSchool(ctx *gin.Context) {
	schoolID := ctx.Param("schoolId")
	userID := ctx.Param("userId")
	res := service.RemoveUserFromSchool(schoolID, userID)
	ctx.JSON(http.StatusOK, res)
}
