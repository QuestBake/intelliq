package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

//Group group model
type Group struct {
	GroupID            bson.ObjectId `json:"groupId" bson:"_id,omitempty"`
	Code               string        `json:"code" bson:"code"`
	CollectionName     string        `json:"collName" bson:"collName"`
	QuestionCategories []string      `json:"quesCategories" bson:"quesCategories"`
	Admin              bson.ObjectId `json:"admin" bson:"admin"`
	Subjects           []Subject     `json:"subjects" bson:"subjects"`
	AuxQuestions       []Question    `json:"auxQuestions" bson:"auxQuestions"`
	CreateDate         time.Time     `json:"createDate" bson:"createDate"`
	LastModifiedDate   time.Time     `json:"lastModifiedDate" bson:"lastModifiedDate"`
}
