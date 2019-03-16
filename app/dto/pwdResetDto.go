package dto

import (
	"github.com/globalsign/mgo/bson"
)

//PasswordDto reset password dto
type PasswordDto struct {
	UserID    bson.ObjectId `json:"userId"`
	OldPwd    string        `json:"oldPwd"`
	NewPwd    string        `json:"newPwd"`
	Mobile    string        `json:"mobile"`
	ForgotPwd bool          `json:"forgotPwd"`
}
