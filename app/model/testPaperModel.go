package model

import (
	"intelliq/app/enums"
	"time"

	"github.com/globalsign/mgo/bson"
)

//TestPaper test paper model
type TestPaper struct {
	TestID           bson.ObjectId         `json:"testId" bson:"_id,omitempty"`
	TeacherID        bson.ObjectId         `json:"teacherId" bson:"teacherId,omitempty"`
	SchoolID         bson.ObjectId         `json:"schoolId" bson:"schoolId,omitempty"`
	GroupCode        string                `json:"groupCode" bson:"groupCode"`
	Standard         int                   `json:"std" bson:"std"`
	Subject          string                `json:"subject" bson:"subject"`
	Tag              string                `json:"tag" bson:"tag"`
	Sets             []set                 `json:"sets" bson:"sets"`
	Status           enums.TestPaperStatus `json:"status" bson:"status"`
	CreateDate       time.Time             `json:"createDate" bson:"createDate"`
	LastModifiedDate time.Time             `json:"lastModifiedDate" bson:"lastModifiedDate"`
}

type set struct {
	Set      int       `json:"set" bson:"set"`
	Sections []section `json:"sections" bson:"sections"`
}

type section struct {
	Type      enums.QuesLength `json:"type" bson:"type"`
	Marks     int              `json:"marks" bson:"marks"`
	Questions Questions        `json:"questions" bson:"questions"`
}

//TestPapers testPaper array
type TestPapers []TestPaper
