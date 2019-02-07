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
	Admin              admin         `json:"admin" bson:"admin"`
	Subjects           []Subject     `json:"subjects" bson:"subjects"`
	AuxQuestions       []Question    `json:"auxQuestions" bson:"auxQuestions"`
	CreateDate         time.Time     `json:"createDate" bson:"createDate"`
	LastModifiedDate   time.Time     `json:"lastModifiedDate" bson:"lastModifiedDate"`
}

type admin struct {
	UserID   bson.ObjectId `json:"userId" bson:"_id"`
	FullName string        `json:"name" bson:"name"`
	Gender   string        `json:"gender" bson:"gender"`
	Mobile   string        `json:"mobile" bson:"mobile"`
	Email    string        `json:"email" bson:"email"`
	School   string        `json:"schoolName" bson:"schoolName"`
}
