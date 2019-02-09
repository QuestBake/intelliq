package service

import (
	"fmt"
	"intelliq/app/common"
	"intelliq/app/enums"
	"intelliq/app/model"
	"intelliq/app/repo"
)

//AddNewData add meta data
func AddNewData(meta *model.Meta) *model.AppResponse {
	metaRepo := repo.NewMetaRepository()
	err := metaRepo.Save(meta)
	if err != nil {
		return &model.AppResponse{
			Status: enums.Status.ERROR,
			Msg:    common.MSG_SAVE_ERROR,
		}
	}
	return &model.AppResponse{
		Status: enums.Status.SUCCESS,
		Body:   common.MSG_SAVE_SUCCESS,
	}
}

//UpdateMetaData updates meta data
func UpdateMetaData(meta *model.Meta) *model.AppResponse {
	metaRepo := repo.NewMetaRepository()
	err := metaRepo.Update(meta)
	if err != nil {
		return &model.AppResponse{
			Status: enums.Status.ERROR,
			Msg:    common.MSG_UPDATE_ERROR,
		}
	}
	return &model.AppResponse{
		Status: enums.Status.SUCCESS,
		Body:   common.MSG_UPDATE_SUCCESS,
	}
}

//ReadMetaData reads data from db
func ReadMetaData() *model.AppResponse {
	metaRepo := repo.NewMetaRepository()
	metaData, err := metaRepo.Read()
	if err != nil {
		fmt.Println(err.Error())
		return &model.AppResponse{
			Status: enums.Status.ERROR,
			Msg:    common.MSG_REQUEST_FAILED,
		}
	}
	return &model.AppResponse{
		Status: enums.Status.SUCCESS,
		Body:   metaData,
	}
}
