package dto

import (
	"github.com/globalsign/mgo/bson"
)

//AuxQuestionDto aux questions request
type AuxQuestionDto struct {
	GroupCode     string        `json:"title"`
	SchoolID      bson.ObjectId `json:"schoolId"`
	ApproverID    bson.ObjectId `json:"approverId"`
	OwnerID       bson.ObjectId `json:"ownerId"`
	AuxQuestionID bson.ObjectId `json:"auxQuesId"`
}
