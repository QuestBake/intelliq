package controller

import (
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//AddMetaData adds meta data
func AddMetaData(ctx *gin.Context) {
	var metaData model.Meta
	err := ctx.BindJSON(&metaData)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.AddNewData(&metaData)
	ctx.JSON(http.StatusOK, res)
}

//UpdateMetaData updates meta data
func UpdateMetaData(ctx *gin.Context) {
	var metaData model.Meta
	err := ctx.BindJSON(&metaData)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.UpdateMetaItems(&metaData)
	ctx.JSON(http.StatusOK, res)
}

//RemoveMetaData removes meta data fields
func RemoveMetaData(ctx *gin.Context) {
	var metaData model.Meta
	err := ctx.BindJSON(&metaData)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.RemoveMetaItems(&metaData)
	ctx.JSON(http.StatusOK, res)
}

//ReadMetaData fetches meta data
func ReadMetaData(ctx *gin.Context) {
	//service.SaveTestQuestions()
	// ctx.JSON(http.StatusOK, "done")
	res := service.ReadMetaData()
	ctx.JSON(http.StatusOK, res)

}
