package controller

import (
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var res *model.AppResponse

//AddMetaData adds meta data
func AddMetaData(ctx *gin.Context) {
	var metaData model.Meta
	err := ctx.BindJSON(&metaData)
	if err != nil {
		res = utility.GetErrorResponse(common.MSG_BAD_INPUT)
	} else {
		res = service.AddNewData(&metaData)
	}
	ctx.JSON(http.StatusOK, res)
}

//UpdateMetaData updates meta data
func UpdateMetaData(ctx *gin.Context) {
	var metaData model.Meta
	err := ctx.BindJSON(&metaData)
	if err != nil {
		res = utility.GetErrorResponse(common.MSG_BAD_INPUT)
	} else {
		res = service.UpdateMetaData(&metaData)
	}
	ctx.JSON(http.StatusOK, res)
}

//ReadMetaData fetches meta data
func ReadMetaData(ctx *gin.Context) {
	res = service.ReadMetaData()
	ctx.JSON(http.StatusOK, res)
}
