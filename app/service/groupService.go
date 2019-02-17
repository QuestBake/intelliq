package service

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"

	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/repo"
)

//AddNewGroup adds new group
func AddNewGroup(group *model.Group) *model.AppResponse {
	groupRepo := repo.NewGroupRepository()
	group.CreateDate = time.Now().UTC()
	group.LastModifiedDate = time.Now().UTC()
	err := groupRepo.Save(group)
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

//UpdateGroup updates existing group
func UpdateGroup(group *model.Group) *model.AppResponse {
	if !utility.IsPrimaryIDValid(group.GroupID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	groupRepo := repo.NewGroupRepository()
	group.LastModifiedDate = time.Now().UTC()
	err := groupRepo.Update(group)
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

//FetchAllGroups gets all groups
func FetchAllGroups(restrict int) *model.AppResponse {
	groupRepo := repo.NewGroupRepository()
	groups, err := groupRepo.FindAll(restrict)
	if err != nil {
		fmt.Println(err.Error())
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(groups)
}

//FetchGroupByCodeOrID get group by code or id
func FetchGroupByCodeOrID(key string, val string) *model.AppResponse {
	var value interface{} = val
	if key == common.PARAM_KEY_ID {
		if utility.IsStringIDValid(val) {
			value = bson.ObjectIdHex(val)
		} else {
			return utility.GetErrorResponse(common.MSG_INVALID_ID)
		}
	}
	groupRepo := repo.NewGroupRepository()
	group, err := groupRepo.FindOne(key, value)
	if err != nil {
		fmt.Println(err.Error())
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(group)
}
