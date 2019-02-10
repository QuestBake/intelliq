package controller

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//AddNewGroup adds new group
func AddNewGroup(ctx *gin.Context) {
	var group model.Group
	err := ctx.BindJSON(&group)
	if err != nil {
		res = utility.GetErrorResponse(common.MSG_BAD_INPUT)
	} else {
		res = service.AddNewGroup(&group)
	}
	ctx.JSON(http.StatusOK, res)
}

//UpdateGroup updates existing group
func UpdateGroup(ctx *gin.Context) {
	var group model.Group
	err := ctx.BindJSON(&group)
	if err != nil {
		res = utility.GetErrorResponse(common.MSG_BAD_INPUT)
	} else {
		res = service.UpdateGroup(&group)
	}
	ctx.JSON(http.StatusOK, res)
}

//ListAllGroups fetches all groups
func ListAllGroups(ctx *gin.Context) {
	restrict := ctx.Param("restrict")
	ctr, err := strconv.Atoi(restrict)
	if err != nil || ctr < 0 {
		res = utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	fmt.Println("Still PRinting")
	res = service.FetchAllGroups(ctr)
	ctx.JSON(http.StatusOK, res)
}
