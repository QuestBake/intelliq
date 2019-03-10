package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"

	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/model"
	"intelliq/app/repo"
)

//AddNewSchool adds new school
func AddNewSchool(school *model.School) *dto.AppResponseDto {
	school.ShortName = strings.ToUpper(school.ShortName)
	school.Code = school.ShortName + "_" + school.Address.Pincode
	school.CreateDate = time.Now().UTC()
	school.LastModifiedDate = time.Now().UTC()
	schoolRepo := repo.NewSchoolRepository()
	err := schoolRepo.Save(school)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_SAVE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_SAVE_SUCCESS)
}

//UpdateSchool updates existing school
func UpdateSchool(school *model.School) *dto.AppResponseDto {
	if !utility.IsPrimaryIDValid(school.SchoolID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	school.LastModifiedDate = time.Now().UTC()
	schoolRepo := repo.NewSchoolRepository()
	err := schoolRepo.Update(school)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_UPDATE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_UPDATE_SUCCESS)
}

//FetchAllSchools gets all schools under one group with either groupID or groupCode
func FetchAllSchools(key string, val string) *dto.AppResponseDto {
	var value interface{}
	switch key {
	case common.PARAM_KEY_ID: // key == _id
		if utility.IsStringIDValid(val) {
			value = bson.ObjectIdHex(val)
		} else {
			return utility.GetErrorResponse(common.MSG_INVALID_ID)
		}
		break
	case common.PARAM_KEY_CODE: // key == code
		val = strings.ToUpper(val)
		if !utility.IsValidGroupCode(val) {
			return utility.GetErrorResponse(common.MSG_INVALID_GROUP)
		}
		value = val
		break
	default:
		return utility.GetErrorResponse(common.MSG_BAD_INPUT)
	}
	schoolRepo := repo.NewSchoolRepository()
	schools, err := schoolRepo.FindAll("group."+key, value)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(schools)
}

//FetchSchoolByCodeOrID get school by Code or id
func FetchSchoolByCodeOrID(key string, val string) *dto.AppResponseDto {
	var value interface{}
	switch key {
	case common.PARAM_KEY_ID: // key == _id
		if utility.IsStringIDValid(val) {
			value = bson.ObjectIdHex(val)
		} else {
			return utility.GetErrorResponse(common.MSG_INVALID_ID)
		}
		break
	case common.PARAM_KEY_CODE: // key == code
		value = strings.ToUpper(val)
		break
	default:
		return utility.GetErrorResponse(common.MSG_BAD_INPUT)
	}
	schoolRepo := repo.NewSchoolRepository()
	school, err := schoolRepo.FindOne(key, value)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(school)
}
