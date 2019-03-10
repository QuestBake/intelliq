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

//AddNewGroup adds new group
func AddNewGroup(group *model.Group) *dto.AppResponseDto {
	group.Code = common.GROUP_CODE_PREFIX + strings.ToUpper(group.Code)
	group.CreateDate = time.Now().UTC()
	group.LastModifiedDate = time.Now().UTC()
	groupRepo := repo.NewGroupRepository()
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
func UpdateGroup(group *model.Group) *dto.AppResponseDto {
	if !utility.IsPrimaryIDValid(group.GroupID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	group.LastModifiedDate = time.Now().UTC()
	groupRepo := repo.NewGroupRepository()
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
func FetchAllGroups(restrict int) *dto.AppResponseDto {
	groupRepo := repo.NewGroupRepository()
	groups, err := groupRepo.FindAll(restrict)
	if err != nil {
		fmt.Println(err.Error())
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(groups)
}

//FetchGroupByCodeOrID get group by code or id
func FetchGroupByCodeOrID(key string, val string) *dto.AppResponseDto {
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
	groupRepo := repo.NewGroupRepository()
	group, err := groupRepo.FindOne(key, value)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(group)
}

// AddSubjectTopicTags adds new tags n topics to group collection for specific subject
func AddSubjectTopicTags(question *model.Question) {
	groupRepo := repo.NewGroupRepository()
	err := groupRepo.AddTopicTags(question)
	if err != nil {
		fmt.Println("TOPIC ADD ERROR => ", err.Error())
	}
}
