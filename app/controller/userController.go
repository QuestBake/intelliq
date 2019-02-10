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
		res = utility.GetErrorResponse(common.MSG_BAD_INPUT)
	} else {
		res = service.AddNewUser(&user)
	}
	ctx.JSON(http.StatusOK, res)
}

//UpdateUserProfile updates user profile
func UpdateUserProfile(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		res = utility.GetErrorResponse(common.MSG_BAD_INPUT)
	} else {
		res = service.UpdateUser(&user)
	}
	ctx.JSON(http.StatusOK, res)
}

//ListAllSchoolAdmins fetches all users with role schooladmin
func ListAllSchoolAdmins(ctx *gin.Context) {
	groupID := ctx.Param("groupId")
	res = service.FetchAllSchoolAdmins(groupID)
	ctx.JSON(http.StatusOK, res)
}
