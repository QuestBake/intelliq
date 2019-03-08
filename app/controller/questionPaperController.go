package controller

import (
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GenerateQuestionPaper generates question paper as per criteria
func GenerateQuestionPaper(ctx *gin.Context) {
	var quesCriteriaDto dto.QuestionCriteriaDto
	err := ctx.BindJSON(&quesCriteriaDto)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.GenerateQuestionPaper(&quesCriteriaDto)
	ctx.JSON(http.StatusOK, res)
}
