package service

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/repo"

	"github.com/globalsign/mgo/bson"
)

//FetchOneQuestion fetches single ques based on quesID
func FetchOneQuestion(groupCode string, quesID string) *model.AppResponse {
	if !utility.IsValidGroupCode(groupCode) {
		return utility.GetErrorResponse(common.MSG_INVALID_GROUP)
	}
	if !utility.IsStringIDValid(quesID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	quesRepo := repo.NewQuestionRepository(groupCode)
	if quesRepo == nil { // panic - recover can be used here .....
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	ques, err := quesRepo.FindOne(bson.ObjectIdHex(quesID))
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REMOVE_ERROR)
	}
	return utility.GetSuccessResponse(ques)
}

//RemoveQuestion deletes rejected/removed status question from coll by teacher
func RemoveQuestion(question *model.Question) *model.AppResponse {
	if !utility.IsPrimaryIDValid(question.QuestionID) {
		return utility.GetErrorResponse(common.MSG_BAD_INPUT)
	}
	if question.Status != enums.CurrentQuestionStatus.REJECTED && // initially rejected while adding new ques
		question.Status != enums.CurrentQuestionStatus.REMOVE { // teacher requested to remove an approved ques
		return utility.GetErrorResponse(common.MSG_INVALID_STATE)
	}
	quesRepo := repo.NewQuestionRepository(question.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	err := quesRepo.Delete(question.QuestionID)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_QUES_REMOVE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_QUES_REMOVE_SUCCESS)
}

//FetchApprovedQuestions fetches all approved ques excluding requester's
func FetchApprovedQuestions(requestDto *dto.QuesRequestDto) *model.AppResponse {
	if !utility.IsValidGroupCode(requestDto.GroupCode) {
		return utility.GetErrorResponse(common.MSG_INVALID_GROUP)
	}
	if !utility.IsPrimaryIDValid(requestDto.UserID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	if requestDto.Std < common.MIN_VALID_STD || requestDto.Std > common.MAX_VALID_STD || requestDto.Subject == "" {
		return utility.GetErrorResponse(common.MSG_BAD_INPUT)
	}
	if requestDto.Limit <= 0 {
		requestDto.Limit = common.DEF_REQUESTS_LIMIT
	}
	quesRepo := repo.NewQuestionRepository(requestDto.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	questions, err := quesRepo.FindApprovedQuestions(requestDto)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(questions)
}
