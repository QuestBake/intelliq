package model

import (
	"intelliq/app/enums"
	"time"

	"github.com/globalsign/mgo/bson"
)

//Template test paper template
type Template struct {
	TemplateID       bson.ObjectId `json:"templateId" bson:"_id,omitempty"`
	TeacherID        bson.ObjectId `json:"teacherId" bson:"teacherId,omitempty"`
	GroupCode        string        `json:"groupCode" bson:"groupCode"`
	Tag              string        `json:"tag" bson:"tag"`
	Criteria         criteria      `json:"criteria" bson:"criteria"`
	Criteria512Hash  string        `json:"criteriaHash" bson:"criteria512Hash"`
	CreateDate       time.Time     `json:"createDate" bson:"createDate"`
	LastModifiedDate time.Time     `json:"lastModifiedDate" bson:"lastModifiedDate"`
}

type criteria struct {
	Standard   int          `json:"std" bson:"std"`
	Subject    string       `json:"subject" bson:"subject"`
	Topics     []string     `json:"topics" bson:"topics"`
	Tags       []string     `json:"tags" bson:"tags"`
	Sets       int          `json:"sets" bson:"sets"`
	Marks      int          `json:"totalMarks" bson:"marks"`
	Length     []length     `json:"length" bson:"length"`
	Difficulty []difficulty `json:"difficulty" bson:"difficulty"`
	TeacherID  bson.ObjectId
}

type length struct {
	Type  enums.QuesLength `json:"type" bson:"type"`
	Count int              `json:"count" bson:"count"`
	Marks int              `json:"marks" bson:"marks"`
}

type difficulty struct {
	Level   enums.QuesDifficulty `json:"level" bson:"level"`
	Percent int                  `json:"percent" bson:"percent"`
}

//Templates template array
type Templates []Template
