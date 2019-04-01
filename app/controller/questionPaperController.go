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

//SaveTestPapers saves question paper and template
func SaveTestPapers(ctx *gin.Context) {
	var testDto dto.TestDto
	err := ctx.BindJSON(&testDto)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.SaveTestDetails(&testDto, false)
	ctx.JSON(http.StatusOK, res)
}

//DraftTestPapers saves question paper as draft and template
func DraftTestPapers(ctx *gin.Context) {
	var testDto dto.TestDto
	err := ctx.BindJSON(&testDto)
	if err != nil {
		res := utility.GetErrorResponse(common.MSG_BAD_INPUT)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := service.SaveTestDetails(&testDto, true)
	ctx.JSON(http.StatusOK, res)
}

//GetDraftSuggestions loads all drafts
func GetDraftSuggestions(ctx *gin.Context) {
	groupCode := ctx.Param("groupCode")
	teacherID := ctx.Param("teacherId")
	res := service.FetchTestPapers(groupCode, teacherID, true)
	ctx.JSON(http.StatusOK, res)
}

//GetReleasePaperSuggestions loads all released Testpapers
func GetReleasePaperSuggestions(ctx *gin.Context) {
	groupCode := ctx.Param("groupCode")
	teacherID := ctx.Param("teacherId")
	res := service.FetchTestPapers(groupCode, teacherID, false)
	ctx.JSON(http.StatusOK, res)
}

//GetTemplateSuggestions loads all templates
func GetTemplateSuggestions(ctx *gin.Context) {
	groupCode := ctx.Param("groupCode")
	teacherID := ctx.Param("teacherId")
	res := service.FetchAllTemplates(groupCode, teacherID)
	ctx.JSON(http.StatusOK, res)
}

//FindDraft retreives one draft
func FindDraft(ctx *gin.Context) {
	groupCode := ctx.Param("groupCode")
	testID := ctx.Param("testId")
	res := service.FetchSinglePaper(groupCode, testID)
	ctx.JSON(http.StatusOK, res)
}

//FindTemplate retreieves one template
func FindTemplate(ctx *gin.Context) {
	groupCode := ctx.Param("groupCode")
	templateID := ctx.Param("templateId")
	res := service.FetchSingleTemplate(groupCode, templateID)
	ctx.JSON(http.StatusOK, res)
}

//RemoveDraft removes test paper from coll
func RemoveDraft(ctx *gin.Context) {
	groupCode := ctx.Param("groupCode")
	testID := ctx.Param("testId")
	res := service.RemoveDraft(groupCode, testID)
	ctx.JSON(http.StatusOK, res)
}

//RemoveTemplate removes template from coll
func RemoveTemplate(ctx *gin.Context) {
	groupCode := ctx.Param("groupCode")
	templateID := ctx.Param("templateId")
	res := service.RemoveTemplate(groupCode, templateID)
	ctx.JSON(http.StatusOK, res)
}
