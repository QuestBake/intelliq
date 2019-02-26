package controller

import (
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindQuestion get question based on quesId
func FindQuestion(ctx *gin.Context) {
	groupCode := ctx.Param("groupCode")
	quesID := ctx.Param("quesId")
	res := service.FetchOneQuestion(groupCode, quesID)
	ctx.JSON(http.StatusOK, res)
}

//GetQuestionsFromBank get all approved questions from bank
func GetQuestionsFromBank(ctx *gin.Context) {
	var requestDto dto.QuesRequestDto
	err := ctx.BindJSON(&requestDto)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.FetchApprovedQuestions(&requestDto)
	ctx.JSON(http.StatusOK, res)
}
