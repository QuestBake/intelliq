package model

import (
	"github.com/globalsign/mgo/bson"
)

//Standard std model
type Standard struct {
	Std      uint16    `json:"std" bson:"std"`
	Subjects []Subject `json:"subjects" bson:"subjects"`
}

//Subject subject model
type Subject struct {
	Title    string        `json:"title" bson:"title"`
	Approver bson.ObjectId `json:"approverId" bson:"approver_id,omitempty"`
	Topics   []string      `json:"topics" bson:"topics"`
}
