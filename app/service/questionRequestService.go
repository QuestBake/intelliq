package service

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/repo"
	"strings"
	"time"
)

func isQuestionInfoValid(question *model.Question) bool {
	return strings.HasPrefix(question.GroupCode, common.GROUP_CODE_PREFIX) &&
		utility.IsPrimaryIDValid(question.Reviewer.UserID) &&
		utility.IsPrimaryIDValid(question.Owner.UserID) &&
		utility.IsPrimaryIDValid(question.School.SchoolID)
}

//RequestAddNewQuestion adds new aux question by teacher
func RequestAddNewQuestion(question *model.Question) *model.AppResponse {
	if !isQuestionInfoValid(question) {
		return utility.GetErrorResponse(common.MSG_BAD_INPUT)
	}
	updateQuestionAttributes(question, enums.CurrentQuestionStatus.NEW, true, true)
	question.FormatTopicTags()
	question.CreateDate = time.Now().UTC()
	quesRepo := repo.NewQuestionRepository(question.GroupCode)
	if quesRepo == nil { // panic - recover can be used here .....
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	var err error
	if utility.IsPrimaryIDValid(question.QuestionID) {
		err = quesRepo.Update(question) //  update resubmitted rejected ques as request
	} else {
		err = quesRepo.Save(question) //  save new ques request
	}
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(common.MSG_QUES_SUBMIT_SUCCESS)
}

//RequestApprovedQuestionUpdate create updated verion of approved question by teacher
func RequestApprovedQuestionUpdate(question *model.Question) *model.AppResponse {
	if !isQuestionInfoValid(question) || !utility.IsPrimaryIDValid(question.QuestionID) {
		return utility.GetErrorResponse(common.MSG_BAD_INPUT)
	}
	if question.Status != enums.CurrentQuestionStatus.APPROVED {
		return utility.GetErrorResponse(common.MSG_INVALID_STATE)
	}
	_id := question.QuestionID
	question.OriginID = &_id
	question.QuestionID = ""

	updateQuestionAttributes(question, enums.CurrentQuestionStatus.TRANSIT, true, false)
	question.CreateDate = time.Now().UTC()
	question.LastModifiedDate = question.CreateDate
	question.FormatTopicTags()
	quesRepo := repo.NewQuestionRepository(question.GroupCode)
	if quesRepo == nil { // panic - recover can be used here .....
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	err := quesRepo.Save(question) //  creates a duplicate copy of original ques with original id tagged
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(common.MSG_QUES_SUBMIT_SUCCESS)
}

//RequestApprovedQuesRemoval changes status to REMOVE ; raised for approved ques by teacher
func RequestApprovedQuesRemoval(question *model.Question) *model.AppResponse {
	if !utility.IsPrimaryIDValid(question.QuestionID) {
		return utility.GetErrorResponse(common.MSG_BAD_INPUT)
	}
	if question.Status != enums.CurrentQuestionStatus.APPROVED {
		return utility.GetErrorResponse(common.MSG_INVALID_STATE)
	}
	quesRepo := repo.NewQuestionRepository(question.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	updateQuestionAttributes(question, enums.CurrentQuestionStatus.REMOVE, false, false)
	err := quesRepo.UpdateLimitedCols(question)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(common.MSG_QUES_SUBMIT_SUCCESS)
}

//ApproveRequest updates existing question status by approver
func ApproveRequest(question *model.Question) *model.AppResponse {
	var quesList model.Questions
	switch question.Status {
	case enums.CurrentQuestionStatus.REMOVE: // remove request raised by teacher
		return RemoveQuestion(question)
	case enums.CurrentQuestionStatus.NEW: // add new ques by teacher
		break
	case enums.CurrentQuestionStatus.TRANSIT: // update existing approved question request by teacher
		if !utility.IsPrimaryIDValid(*question.OriginID) { //checks for validity of original question for this updated version
			return utility.GetErrorResponse(common.MSG_BAD_INPUT)
		}
		ques := model.Question{
			QuestionID: *question.OriginID,
			Status:     enums.CurrentQuestionStatus.OBSOLETE, // original ques updated to OBSOLETE ; cleaned up by scheduler
		}
		quesList = append(quesList, ques)
		break
	default: // no other status processed
		return utility.GetErrorResponse(common.MSG_INVALID_STATE)
	}
	updateQuestionAttributes(question, enums.CurrentQuestionStatus.APPROVED, true, true)
	quesList = append(quesList, *question)
	quesRepo := repo.NewQuestionRepository(question.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	err := quesRepo.BulkUpdate(quesList)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_QUES_STATUS_ERROR)
	}
	AddSubjectTopicTags(question) // inserts topic/tags in group based on subject title
	return utility.GetSuccessResponse(common.MSG_QUES_STATUS_SUCCESS)
}

//RejectRequest updates existing question status by approver
func RejectRequest(question *model.Question) *model.AppResponse {
	switch question.Status {
	case enums.CurrentQuestionStatus.REMOVE: // remove request raised by teacher
		updateQuestionAttributes(question, enums.CurrentQuestionStatus.APPROVED, true, true) // switch status back to APPROVED
		break
	case enums.CurrentQuestionStatus.NEW: // add new ques by teacher
		updateQuestionAttributes(question, enums.CurrentQuestionStatus.REJECTED, false, true) // reject status
		break
	case enums.CurrentQuestionStatus.TRANSIT: // update existing approved question request by teacher
		updateQuestionAttributes(question, enums.CurrentQuestionStatus.REJECTED, false, false) // reject status ; retain originID
		break
	default: // no other status processed
		return utility.GetErrorResponse(common.MSG_INVALID_STATE)
	}
	quesRepo := repo.NewQuestionRepository(question.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	err := quesRepo.UpdateLimitedCols(question)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_QUES_STATUS_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_QUES_STATUS_SUCCESS)
}

func updateQuestionAttributes(question *model.Question, status enums.QuestionStatus, clearRejectReason bool, clearOriginID bool) {
	question.Status = status
	question.LastModifiedDate = time.Now().UTC()
	if clearRejectReason {
		question.RejectDesc = ""
	}
	if clearOriginID {
		question.OriginID = nil
	}
}

//FetchReviewerRequests fetches all ques with status NEW,TRANSIT,REMOVE for a reviewer
func FetchReviewerRequests(requestDto *dto.QuesRequestDto) *model.AppResponse {
	if !utility.IsValidGroupCode(requestDto.GroupCode) {
		return utility.GetErrorResponse(common.MSG_INVALID_GROUP)
	}
	if !utility.IsPrimaryIDValid(requestDto.SchoolID) ||
		!utility.IsPrimaryIDValid(requestDto.UserID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	if requestDto.Limit <= 0 { // default limit on record result
		requestDto.Limit = common.DEF_REQUESTS_LIMIT
	}
	status := []enums.QuestionStatus{enums.CurrentQuestionStatus.NEW,
		enums.CurrentQuestionStatus.TRANSIT, enums.CurrentQuestionStatus.REMOVE}
	quesRepo := repo.NewQuestionRepository(requestDto.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	questions, err := quesRepo.FindReviewersRequests(requestDto, status)
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

//FetchTeacherRequests fetches all ques with either od status : APPROVED / REJECTED/ PENDING for a teacher
func FetchTeacherRequests(requestDto *dto.QuesRequestDto) *model.AppResponse {
	if !utility.IsValidGroupCode(requestDto.GroupCode) {
		return utility.GetErrorResponse(common.MSG_INVALID_GROUP)
	}
	if !utility.IsPrimaryIDValid(requestDto.SchoolID) ||
		!utility.IsPrimaryIDValid(requestDto.UserID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	if requestDto.Std < common.MIN_VALID_STD || requestDto.Std > common.MAX_VALID_STD || requestDto.Subject == "" {
		return utility.GetErrorResponse(common.MSG_BAD_INPUT)
	}
	var status []enums.QuestionStatus
	switch requestDto.Status {
	case enums.CurrentQuestionStatus.APPROVED:
		status = []enums.QuestionStatus{enums.CurrentQuestionStatus.APPROVED}
		break
	case enums.CurrentQuestionStatus.REJECTED:
		status = []enums.QuestionStatus{enums.CurrentQuestionStatus.REJECTED}
		break
	case enums.CurrentQuestionStatus.PENDING:
		status = []enums.QuestionStatus{enums.CurrentQuestionStatus.NEW,
			enums.CurrentQuestionStatus.TRANSIT, enums.CurrentQuestionStatus.REMOVE}
		break
	default:
		return utility.GetErrorResponse(common.MSG_NO_STATUS)
	}
	if requestDto.Limit <= 0 {
		requestDto.Limit = common.DEF_REQUESTS_LIMIT
	}
	quesRepo := repo.NewQuestionRepository(requestDto.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	questions, err := quesRepo.FindTeachersRequests(requestDto, status)
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
