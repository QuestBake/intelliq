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
	Reviewer bson.ObjectId `json:"reviewerId" bson:"reviewer_id,omitempty"`
	Topics   []string      `json:"topics" bson:"topics"`
	Tags     []string      `json:"tags" bson:"tags"`
}
