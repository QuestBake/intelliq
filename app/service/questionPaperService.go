package service

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/helper"
	"intelliq/app/model"
	"intelliq/app/repo"
	"time"
)

//GenerateQuestionPaper generated question paper as per criteria provided
func GenerateQuestionPaper(criteriaDto *dto.QuestionCriteriaDto) *model.AppResponse {
	errResponse := validateRequest(criteriaDto.GroupCode,
		criteriaDto.Subject, criteriaDto.Standard)
	if errResponse != nil {
		return errResponse
	}
	quesRepo := repo.NewQuestionRepository(criteriaDto.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	criteriaDto.GenerateNatives()
	dbstart := time.Now()
	sectionQuesMap, err := quesRepo.FilterQuestionsForPaper(criteriaDto)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	if len(sectionQuesMap) == 0 {
		return utility.GetErrorResponse(common.MSG_NO_RECORD)
	}
	fmt.Println("DB QUERY TIME := ", time.Since(dbstart))
	start := time.Now()
	helper.PrioritiseDifficultyList(criteriaDto.Difficulty)
	sectionCountMap := helper.GetSectionCountMap(criteriaDto.Length)
	sectionChannel := make(chan *model.Section)
	for section, lvlMap := range sectionQuesMap { // go routine
		go helper.GetResultSectionQuesList(lvlMap, criteriaDto.Difficulty,
			section, sectionCountMap[section], criteriaDto.Sets, sectionChannel)
	}
	var quesSectionList []model.Section
	for i := 0; i < len(sectionQuesMap); i++ {
		quesSectionList = append(quesSectionList, *<-sectionChannel)
	}
	close(sectionChannel)
	paperChannel := make(chan *model.QuestionPaper)
	for currSet := 0; currSet < criteriaDto.Sets; currSet++ { // go routine
		go helper.GenerateQuestionPaper(quesSectionList,
			currSet, criteriaDto.Difficulty, paperChannel)
	}
	var questionPapers []model.QuestionPaper
	for currSet := 0; currSet < criteriaDto.Sets; currSet++ {
		questionPapers = append(questionPapers, *<-paperChannel)
	}
	close(paperChannel)
	fmt.Println("TOTAL ALGO TIME := ", time.Since(start))
	return utility.GetSuccessResponse(questionPapers)
}
