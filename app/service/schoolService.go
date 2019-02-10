package service

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/repo"
	"time"

	"github.com/globalsign/mgo/bson"
)

//AddNewSchool adds new school
func AddNewSchool(school *model.School) *model.AppResponse {
	schoolRepo := repo.NewSchoolRepository()
	school.Code = school.ShortName + "_" + school.Address.Pincode
	school.CreateDate = time.Now().UTC()
	school.LastModifiedDate = time.Now().UTC()
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

//FetchAllSchools gets all schools under one group with either groupID or groupCode
func FetchAllSchools(key string, val string) *model.AppResponse {
	var value interface{} = val
	if key == common.PARAM_KEY_ID {
		if utility.IsStringIDValid(val) {
			value = bson.ObjectIdHex(val)
		} else {
			return utility.GetErrorResponse(common.MSG_INVALID_ID)
		}
	}
	schoolRepo := repo.NewSchoolRepository()
	schools, err := schoolRepo.FindAll("group."+key, value)
	if err != nil {
		fmt.Println(err.Error())
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(schools)
}
