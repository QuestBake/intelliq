package common

import (
	"intelliq/app/enums"
	"intelliq/app/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//GetErrorResponse prepares error response with msg
func GetErrorResponse(msg string) *model.AppResponse {
	return &model.AppResponse{
		Status: enums.Status.ERROR,
		Msg:    msg,
	}
}

//GetSuccessResponse prepares success response with body
func GetSuccessResponse(body interface{}) *model.AppResponse {
	return &model.AppResponse{
		Status: enums.Status.SUCCESS,
		Body:   body,
	}
}

//IsPrimaryIDValid checks if bson id is vaild
func IsPrimaryIDValid(_id bson.ObjectId) bool {
	return _id.Valid()
}

//IsStringIDValid checks if string id is vaild bson id
func IsStringIDValid(_id string) bool {
	return bson.IsObjectIdHex(_id)
}

//GetErrorMsg specific db errors
func GetErrorMsg(err error) string {
	switch err.(type) {
	case *mgo.LastError:
		errorCode := (err.(*mgo.LastError)).Code
		switch errorCode {
		case ERR_CODE_DUPLICATE:
			return MSG_DUPLICATE_RECORD
		default:
			return ""
		}
	default:
		return ""
	}
}
