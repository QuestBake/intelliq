package service

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/dto"
	"intelliq/app/model"
	"intelliq/app/repo"
)

//AddNewData add meta data
func AddNewData(meta *model.Meta) *dto.AppResponseDto {
	metaRepo := repo.NewMetaRepository()
	err := metaRepo.Save(meta)
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

//UpdateMetaItems updates meta data
func UpdateMetaItems(meta *model.Meta) *dto.AppResponseDto {
	if !utility.IsPrimaryIDValid(meta.MetaID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	metaRepo := repo.NewMetaRepository()
	err := metaRepo.Update(meta)
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

//RemoveMetaItems removes meta data items
func RemoveMetaItems(meta *model.Meta) *dto.AppResponseDto {
	if !utility.IsPrimaryIDValid(meta.MetaID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	metaRepo := repo.NewMetaRepository()
	err := metaRepo.Remove(meta)
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_DELETE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_DELETE_SUCCESS)
}

//ReadMetaData reads data from db
func ReadMetaData() *dto.AppResponseDto {
	metaRepo := repo.NewMetaRepository()
	metaData, err := metaRepo.Read()
	if err != nil {
		fmt.Println(err.Error())
		errorMsg := utility.GetErrorMsg(err)
		if len(errorMsg) > 0 {
			return utility.GetErrorResponse(errorMsg)
		}
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(metaData)
}
