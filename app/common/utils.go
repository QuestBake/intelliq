package common

import (
	"fmt"
	"intelliq/app/enums"
	"intelliq/app/model"
	"log"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Printf("error type: %T", err)
	if err == mgo.ErrNotFound {
		return MSG_NO_RECORD
	}
	switch err.(type) {
	case *mgo.LastError:
		errorCode := (err.(*mgo.LastError)).Code
		switch errorCode {
		case ERR_CODE_DUPLICATE:
			return MSG_DUPLICATE_RECORD
		default:
			return ""
		}
	case *mgo.QueryError:
		errorCode := (err.(*mgo.QueryError)).Code
		fmt.Println("QUERY ERROR CODE => ", errorCode)
		return err.Error()
	default:
		return ""
	}
}

//FormatDateToString formats date to readable string
func FormatDateToString(date time.Time) string {
	return date.Format(DATE_TIME_FORMAT)
}

//FormatStringToDate formats string to time
func FormatStringToDate(date string) time.Time {
	t, _ := time.Parse(DATE_TIME_FORMAT, date)
	return t
}

//EncryptData data encryption
func EncryptData(pwd string) string {
	pwdHash := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(pwdHash, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//ComparePasswords compare hashed and plain text
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	plainHash := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainHash)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
