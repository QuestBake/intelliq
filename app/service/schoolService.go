package service

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/repo"
	"time"
)

//AddNewSchool adds new school
func AddNewSchool(school *model.School) *model.AppResponse {
	schoolRepo := repo.NewSchoolRepository()
	school.Code = school.ShortName + "_" + school.Address.Pincode
	school.CreateDate = time.Now()
	school.LastModifiedDate = time.Now()
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

//FetchAllSchools gets all schools
func FetchAllSchools() *model.AppResponse {
	schoolRepo := repo.NewSchoolRepository()
	schools, err := schoolRepo.FindAll()
	if err != nil {
		fmt.Println(err.Error())
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(schools)
}
