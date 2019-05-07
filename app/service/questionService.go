package service

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/repo"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
)

//FetchOneQuestion fetches single ques based on quesID
func FetchOneQuestion(groupCode string, quesID string) *dto.AppResponseDto {
	if !utility.IsValidGroupCode(groupCode) {
		return utility.GetErrorResponse(common.MSG_INVALID_GROUP)
	}
	if !utility.IsStringIDValid(quesID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	quesRepo := repo.NewQuestionRepository(groupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	ques, err := quesRepo.FindOne(bson.ObjectIdHex(quesID))
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_NO_RECORD)
	}
	return utility.GetSuccessResponse(ques)
}

//RemoveQuestion deletes rejected/removed status question from coll by teacher
func RemoveQuestion(question *model.Question) *dto.AppResponseDto {
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

//FetchApprovedQuestions fetches all approved ques
func FetchApprovedQuestions(requestDto *dto.QuesRequestDto) *dto.AppResponseDto {
	if !utility.IsValidGroupCode(requestDto.GroupCode) {
		return utility.GetErrorResponse(common.MSG_INVALID_GROUP)
	}
	if !utility.IsPrimaryIDValid(requestDto.UserID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	quesRepo := repo.NewQuestionRepository(requestDto.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	responseDto, err := quesRepo.FindApprovedQuestions(requestDto)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(responseDto)
}

//FetchQuestionSuggestions fetch similar questions as per search term
func FetchQuestionSuggestions(criteriaDto *dto.QuestionCriteriaDto) *dto.AppResponseDto {
	errResponse := validateRequest(criteriaDto.GroupCode,
		criteriaDto.Subject, criteriaDto.Standard)
	if errResponse != nil {
		return errResponse
	}
	quesRepo := repo.NewQuestionRepository(criteriaDto.GroupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	questions, err := quesRepo.FilterQuestionsPerSearchTerm(criteriaDto)
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

//FilterQuestions filters questions as per criteria provided
func FilterQuestions(criteriaDto *dto.QuestionCriteriaDto) *dto.AppResponseDto {
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
	questions, err := quesRepo.FilterQuestionsPerCriteria(criteriaDto)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	if len(questions) == 0 {
		return utility.GetErrorResponse(common.MSG_NO_RECORD)
	}
	return utility.GetSuccessResponse(questions)
}

//RemoveObsoleteQuestions filters questions as per criteria provided
func RemoveObsoleteQuestions(groupCode string) *dto.AppResponseDto {
	if !utility.IsValidGroupCode(groupCode) {
		return utility.GetErrorResponse(common.MSG_INVALID_GROUP)
	}
	quesRepo := repo.NewQuestionRepository(groupCode)
	if quesRepo == nil {
		return utility.GetErrorResponse(common.MSG_UNATHORIZED_ACCESS)
	}
	info, err := quesRepo.DeleteAll()
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_QUES_REMOVE_ERROR)
	}
	return utility.GetSuccessResponse("Removed " + strconv.Itoa(info) + " from group " + groupCode)
}

func validateRequest(groupCode, subject string, std int) *dto.AppResponseDto {
	if !utility.IsValidGroupCode(groupCode) {
		return utility.GetErrorResponse(common.MSG_INVALID_GROUP)
	}
	if std < common.MIN_VALID_STD ||
		std > common.MAX_VALID_STD ||
		subject == "" {
		return utility.GetErrorResponse(common.MSG_BAD_INPUT)
	}
	return nil
}

func SaveTestQuestions() {
	quesRepo := repo.NewQuestionRepository("GP_DPS")
	var quesList model.Questions
	ctr := 1000
	topics := []string{"T1", "T2", "T3", "T4", "T5"}
	lengths := []enums.QuesLength{enums.Length.OBJECTIVE, enums.Length.SHORT, enums.Length.BRIEF,
		enums.Length.LONG}
	dificulties := []enums.QuesDifficulty{enums.DifficultyLvl.EASY, enums.DifficultyLvl.MEDIUM,
		enums.DifficultyLvl.HARD}
	tags := []string{"mtag1", "mtag2", "mtag3", "mtag4", "mtag5"}
	tagLen := len(tags)

	for _, topic := range topics {
		for _, length := range lengths {
			for _, difficulty := range dificulties {
				for i := 0; i < 5000; i++ {
					question := model.Question{
						GroupCode:        "GP_DPS",
						Std:              3,
						Subject:          "Mathematics",
						Title:            "T" + strconv.Itoa(ctr),
						Topic:            topic,
						Difficulty:       difficulty,
						Length:           length,
						Status:           enums.CurrentQuestionStatus.APPROVED,
						CreateDate:       time.Now().UTC(),
						LastModifiedDate: time.Now().UTC(),
						Owner: model.Contributor{
							UserID:   bson.ObjectIdHex("5c67cbea2ed7b04f1bdc1817"),
							UserName: "user@UT2_992",
						},
						Reviewer: model.Contributor{
							UserID:   bson.ObjectIdHex("5c67cbea2ed7b04f1bdc1810"),
							UserName: "user@UT3_998",
						},
						School: model.School{
							SchoolID:  bson.ObjectIdHex("5c5ee4cd7a4b5e31a340dccf"),
							ShortName: "DPS",
							Code:      "DPS_143001",
							Address: model.Address{
								City:    "Amritsar",
								State:   "PUNJAB",
								Pincode: "143001",
							},
						},
						Tags: []string{tags[ctr%tagLen], tags[(ctr*i)%tagLen]},
					}
					quesList = append(quesList, question)
					ctr++
				}
			}
		}
	}

	err := quesRepo.BulkSave(quesList)
	if err != nil {
		fmt.Println("BULK ERROR => ", err)
	} else {
		fmt.Println("BULK SUCCESS")
	}
}
