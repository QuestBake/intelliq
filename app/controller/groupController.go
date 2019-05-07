package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"intelliq/app/cachestore"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/service"
)

//AddNewGroup adds new group
func AddNewGroup(ctx *gin.Context) {
	var group model.Group
	err := ctx.BindJSON(&group)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AddNewGroup(&group)
	ctx.JSON(http.StatusOK, res)
}

//UpdateGroup updates existing group
func UpdateGroup(ctx *gin.Context) {
	var group model.Group
	err := ctx.BindJSON(&group)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateGroup(&group)
	if res.Status == enums.Status.SUCCESS && res.Body != nil {
		if cachestore.CheckCache(ctx, group.Code) {
			cachestore.SetCache(ctx, group.Code, group,
				common.CACHE_OBJ_LONG_TIMEOUT, true)
		}
		if cachestore.CheckCache(ctx, group.GroupID.Hex()) {
			cachestore.SetCache(ctx, group.GroupID.Hex(), group,
				common.CACHE_OBJ_LONG_TIMEOUT, true)
		}
	}
	ctx.JSON(http.StatusOK, res)
}

//ListAllGroups fetches all groups
func ListAllGroups(ctx *gin.Context) {
	restrict := ctx.Param("restrict")
	ctr, err := strconv.Atoi(restrict)
	if err != nil || ctr < 0 || ctr > 1 {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.FetchAllGroups(ctr)
	ctx.JSON(http.StatusOK, res)
}

//ListGroupByCodeOrID List group by id or code
func ListGroupByCodeOrID(ctx *gin.Context) {
	key := ctx.Param("key")
	val := ctx.Param("val")
	if cachestore.CheckCache(ctx, val) {
		cacheVal := cachestore.GetCache(ctx, val).(string)
		var group model.Group
		json.Unmarshal([]byte(cacheVal), &group)
		fmt.Println("Reading group details from cache!!")
		ctx.JSON(http.StatusOK, utility.GetSuccessResponse(group))
	} else {
		res := service.FetchGroupByCodeOrID(key, val)
		if res.Status == enums.Status.SUCCESS && res.Body != nil {
			cachestore.SetCache(ctx, val, res.Body,
				common.CACHE_OBJ_LONG_TIMEOUT, true)
		}
		ctx.JSON(http.StatusOK, res)
	}
}
