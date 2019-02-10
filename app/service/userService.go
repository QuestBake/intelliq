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

//AddNewUser adds new user
func AddNewUser(user *model.User) *model.AppResponse {
	userRepo := repo.NewUserRepository()
	user.Password = "Temp_" + user.Mobile
	user.CreateDate = time.Now()
	user.LastModifiedDate = time.Now()
	err := userRepo.Save(user)
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

//UpdateUser updates existing user
func UpdateUser(user *model.User) *model.AppResponse {
	if !utility.IsPrimaryIDValid(user.UserID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	userRepo := repo.NewUserRepository()
	user.LastModifiedDate = time.Now()
	err := userRepo.Update(user)
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

//FetchAllSchoolAdmins gets all users with role school admin for a group
func FetchAllSchoolAdmins(groupID string) *model.AppResponse {
	if utility.IsStringIDValid(groupID) {
		userRepo := repo.NewUserRepository()
		users, err := userRepo.FindAllSchoolAdmins(bson.ObjectIdHex(groupID))
		if err != nil {
			fmt.Println(err.Error())
			return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
		}
		return utility.GetSuccessResponse(users)
	}
	return utility.GetErrorResponse(common.MSG_INVALID_ID)
}
