package controller

import (
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//RequestAdd adds new aux question for approval
func RequestAdd(ctx *gin.Context) {
	var question model.Question
	err := ctx.BindJSON(&question)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.RequestAddNewQuestion(&question)
	ctx.JSON(http.StatusOK, res)
}

//RequestUpdate updates question for approval
func RequestUpdate(ctx *gin.Context) {
	var question model.Question
	err := ctx.BindJSON(&question)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if question.Status == enums.CurrentQuestionStatus.REJECTED &&
		question.OriginID == nil { //resubmit new ques rejected by reviewer initially
		res := service.RequestAddNewQuestion(&question)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := service.RequestApprovedQuestionUpdate(&question)
	ctx.JSON(http.StatusOK, res)
}

//RequestRemoval changes question status to REMOVE
func RequestRemoval(ctx *gin.Context) {
	var question model.Question
	err := ctx.BindJSON(&question)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if question.Status == enums.CurrentQuestionStatus.REJECTED {
		res := service.RemoveQuestion(&question)
		ctx.JSON(http.StatusOK, res)
		return
	}
	res := service.RequestApprovedQuesRemoval(&question)
	ctx.JSON(http.StatusOK, res)
}

//ApproveRequest aprroves question in collection
func ApproveRequest(ctx *gin.Context) {
	var question model.Question
	err := ctx.BindJSON(&question)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.ApproveRequest(&question)
	ctx.JSON(http.StatusOK, res)
}

//RejectRequest rejects question in collection
func RejectRequest(ctx *gin.Context) {
	var question model.Question
	err := ctx.BindJSON(&question)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.RejectRequest(&question)
	ctx.JSON(http.StatusOK, res)
}

//GetReviewerRequests gets all pending request for a reviewer
func GetReviewerRequests(ctx *gin.Context) {
	var requestDto dto.QuesRequestDto
	err := ctx.BindJSON(&requestDto)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.FetchReviewerRequests(&requestDto)
	ctx.JSON(http.StatusOK, res)
}

//GetTeacherRequests gets all requests based on status for a teacher
func GetTeacherRequests(ctx *gin.Context) {
	var requestDto dto.QuesRequestDto
	err := ctx.BindJSON(&requestDto)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.FetchTeacherRequests(&requestDto)
	ctx.JSON(http.StatusOK, res)
}
