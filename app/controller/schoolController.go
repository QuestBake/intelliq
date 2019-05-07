package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"intelliq/app/cachestore"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/service"
)

//AddNewSchool adds new school
func AddNewSchool(ctx *gin.Context) {
	var school model.School
	err := ctx.BindJSON(&school)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AddNewSchool(&school)
	ctx.JSON(http.StatusOK, res)
}

//UpdateSchoolProfile updates existing school profile
func UpdateSchoolProfile(ctx *gin.Context) {
	var school model.School
	err := ctx.BindJSON(&school)
	if err != nil {
		fmt.Println(err)
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateSchool(&school)
	if res.Status == enums.Status.SUCCESS && res.Body != nil {
		if cachestore.CheckCache(ctx, school.Code) {
			cachestore.SetCache(ctx, school.Code, school,
				common.CACHE_OBJ_LONG_TIMEOUT, true)
		}
		if cachestore.CheckCache(ctx, school.SchoolID.Hex()) {
			cachestore.SetCache(ctx, school.SchoolID.Hex(), school,
				common.CACHE_OBJ_LONG_TIMEOUT, true)
		}
	}
	ctx.JSON(http.StatusOK, res)
}

//ListAllSchools fetches all schools under a group using groupCode or groupID
func ListAllSchools(ctx *gin.Context) {
	key := ctx.Param("key")
	val := ctx.Param("val")
	res := service.FetchAllSchools(key, val)
	ctx.JSON(http.StatusOK, res)
}

//ListSchoolByCodeOrID list school by id or code
func ListSchoolByCodeOrID(ctx *gin.Context) {
	key := ctx.Param("key")
	val := ctx.Param("val")
	if cachestore.CheckCache(ctx, val) {
		cacheVal := cachestore.GetCache(ctx, val).(string)
		var school model.School
		json.Unmarshal([]byte(cacheVal), &school)
		fmt.Println("Reading school details from cache!!")
		ctx.JSON(http.StatusOK, utility.GetSuccessResponse(school))
	} else {
		res := service.FetchSchoolByCodeOrID(key, val)
		if res.Status == enums.Status.SUCCESS && res.Body != nil {
			cachestore.SetCache(ctx, val, res.Body,
				common.CACHE_OBJ_LONG_TIMEOUT, true)
		}
		ctx.JSON(http.StatusOK, res)
	}
}
