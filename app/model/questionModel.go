package model

import (
	"project/intelliq/app/enums"
	"time"

	"github.com/globalsign/mgo/bson"
)

//Question question model
type Question struct {
	QuestionID       bson.ObjectId    `json:"quesId" bson:"_id,omitempty"`
	Title            string           `json:"title" bson:"title"`
	Std              uint16           `json:"std" bson:"std"`
	Subject          string           `json:"subject" bson:"subject"`
	Topic            string           `json:"topic" bson:"topic"`
	NewTopic         bool             `json:"newTopic" bson:"newTopic"`
	Difficulty       enums.Difficulty `json:"difficulty" bson:"difficulty"`
	Length           enums.QuesLength `json:"length" bson:"length"`
	Tags             []string         `json:"tags" bson:"tags"`
	Category         string           `json:"category" bson:"category"`
	ImageURL         string           `json:"imageUrl" bson:"imageUrl"`
	Owner            User             `json:"owner" bson:"owner"`
	Approver         User             `json:"approver" bson:"approver"`
	School           School           `json:"school" bson:"school"`
	Group            Group            `json:"group" bson:"group"`
	CreateDate       time.Time        `json:"createDate" bson:"createDate"`
	LastModifiedDate time.Time        `json:"lastModifiedDate" bson:"lastModifiedDate"`
}

//AuxQuestionRequest request to view aux questions
type AuxQuestionRequest struct {
	Group      Group         `json:"group" bson:"group"`
	SchoolID   bson.ObjectId `json:"schoolId" bson:"schoolId"`
	TeacherID  bson.ObjectId `json:"teacherId" bson:"teacherId"`
	ApproverID bson.ObjectId `json:"approverId" bson:"approverId"`
	AuxQuesID  bson.ObjectId `json:"auxQuesId" bson:"auxQuesId"`
}
