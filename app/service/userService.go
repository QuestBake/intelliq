package service

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/repo"
	"strconv"
	"time"

	"github.com/globalsign/mgo/bson"
)

//AddNewUser adds new user
func AddNewUser(user *model.User) *model.AppResponse {
	userRepo := repo.NewUserRepository()
	user.Password = utility.EncryptData(common.TEMP_PWD_PREFIX + user.Mobile)
	user.CreateDate = time.Now().UTC()
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
	user.LastModifiedDate = time.Now().UTC()
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

//FetchAllTeachers gets all teachers within school
func FetchAllTeachers(schoolID string) *model.AppResponse {
	if utility.IsStringIDValid(schoolID) {
		userRepo := repo.NewUserRepository()
		users, err := userRepo.FindAllSchoolTeachers(bson.ObjectIdHex(schoolID))
		if err != nil {
			fmt.Println(err.Error())
			return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
		}
		return utility.GetSuccessResponse(users)
	}
	return utility.GetErrorResponse(common.MSG_INVALID_ID)
}

//TransferUserRole transfers user roles
func TransferUserRole(roleType string, fromUserID string, toUserID string) *model.AppResponse {
	role, errs := strconv.Atoi(roleType)
	if errs != nil || role < common.MIN_VALID_ROLE || role > common.MAX_VALID_ROLE {
		return utility.GetErrorResponse(common.MSG_NO_ROLE)
	}
	if !utility.IsStringIDValid(fromUserID) || !utility.IsStringIDValid(toUserID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	userRepo := repo.NewUserRepository()
	msg, err := userRepo.TransferRole(enums.RoleType(role), bson.ObjectIdHex(fromUserID), bson.ObjectIdHex(toUserID))
	if err != nil || len(msg) > 0 {
		if len(msg) > 0 {
			return utility.GetErrorResponse(msg)
		}
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_UPDATE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_UPDATE_SUCCESS)
}
