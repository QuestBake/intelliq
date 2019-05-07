package dto

import (
	"intelliq/app/enums"
	"intelliq/app/model"

	"github.com/globalsign/mgo/bson"
)

//QuesRequestDto request params
type QuesRequestDto struct {
	GroupCode string               `json:"groupCode"`
	SchoolID  bson.ObjectId        `json:"schoolId"`
	UserID    bson.ObjectId        `json:"userId"` // reviewer || teacher
	Status    enums.QuestionStatus `json:"status"` // quest status
	Page      int                  `json:"page"`   // for pagination
	Standards []model.Standard     `json:"standards"`
	GetCount  bool                 `json:"getCount"` // return total records count along with data
}
