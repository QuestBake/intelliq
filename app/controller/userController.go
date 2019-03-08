package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/service"
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

//ListAllSchoolAdminsUnderGroup fetches all users with role schooladmin
func ListAllSchoolAdminsUnderGroup(ctx *gin.Context) {
	groupID := ctx.Param("groupId")
	res := service.FetchAllSchoolAdmins(groupID)
	ctx.JSON(http.StatusOK, res)
}

//ListAllTeachersUnderSchool fetches all users within school
func ListAllTeachersUnderSchool(ctx *gin.Context) {
	schoolID := ctx.Param("schoolId")
	res := service.FetchAllTeachers(schoolID)
	ctx.JSON(http.StatusOK, res)
}

//ListAllTeachersUnderReviewer fetches all users within school
func ListAllTeachersUnderReviewer(ctx *gin.Context) {
	schoolID := ctx.Param("schoolId")
	reviewerID := ctx.Param("reviewerId")
	res := service.FetchAllTeachersUnderReviewer(schoolID, reviewerID)
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

//AddBulkUsers adds new users
func AddBulkUsers(ctx *gin.Context) {
	var users model.Users
	err := ctx.BindJSON(&users)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AddBulkUser(users)
	ctx.JSON(http.StatusOK, res)
}

//UpdateBulkUsers updates bulk users
func UpdateBulkUsers(ctx *gin.Context) {
	var users model.Users
	err := ctx.BindJSON(&users)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateBulkUsers(users)
	ctx.JSON(http.StatusOK, res)
}

//AuthenticateUser authenticate and returns AppResponse object
func AuthenticateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AuthenticateUser(&user)
	ctx.JSON(http.StatusOK, res)
}

//Logout logs out user and clear sessions
func Logout(ctx *gin.Context) {
	userID := ctx.Param("userId")
	res := service.Logout(userID)
	ctx.JSON(http.StatusOK, res)
}

//ListUserByMobileOrID get user info by id or mobile number
func ListUserByMobileOrID(ctx *gin.Context) {
	key := ctx.Param("key")
	val := ctx.Param("val")
	if len(key) == 0 || len(val) == 0 ||
		(key != common.PARAM_KEY_ID && key != common.PARAM_KEY_MOBILE) {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.FetchUserByMobileOrID(key, val)
	ctx.JSON(http.StatusOK, res)
}
