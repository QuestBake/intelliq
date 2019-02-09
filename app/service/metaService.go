package service

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"intelliq/app/model"
	"intelliq/app/repo"
)

//AddNewData add meta data
func AddNewData(meta *model.Meta) *model.AppResponse {
	metaRepo := repo.NewMetaRepository()
	err := metaRepo.Save(meta)
	if err != nil {
		fmt.Println(err.Error())
		return utility.GetErrorResponse(common.MSG_SAVE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_SAVE_SUCCESS)
}

//UpdateMetaData updates meta data
func UpdateMetaData(meta *model.Meta) *model.AppResponse {
	if !utility.IsPrimaryIDValid(meta.MetaID) {
		return utility.GetErrorResponse(common.MSG_INVALID_ID)
	}
	metaRepo := repo.NewMetaRepository()
	err := metaRepo.Update(meta)
	if err != nil {
		fmt.Println(err.Error())
		return utility.GetErrorResponse(common.MSG_UPDATE_ERROR)
	}
	return utility.GetSuccessResponse(common.MSG_UPDATE_SUCCESS)
}

//ReadMetaData reads data from db
func ReadMetaData() *model.AppResponse {
	metaRepo := repo.NewMetaRepository()
	metaData, err := metaRepo.Read()
	if err != nil {
		fmt.Println(err.Error())
		return utility.GetErrorResponse(common.MSG_REQUEST_FAILED)
	}
	return utility.GetSuccessResponse(metaData)
}
