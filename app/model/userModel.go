package model

import (
	"intelliq/app/enums"
	"time"

	"github.com/globalsign/mgo/bson"
)

//User user model
type User struct {
	UserID           bson.ObjectId `json:"userId" bson:"_id,omitempty"`
	FullName         string        `json:"name" bson:"name"`
	UserName         string        `json:"userName" bson:"userName"`
	Gender           string        `json:"gender" bson:"gender"`
	Mobile           string        `json:"mobile" bson:"mobile"`
	Email            string        `json:"email" bson:"email"`
	Password         string        `json:"password" bson:"password"`
	DOB              time.Time     `json:"dob" bson:"dob"`
	CreateDate       time.Time     `json:"createDate" bson:"createDate"`
	LastModifiedDate time.Time     `json:"lastModifiedDate" bson:"lastModifiedDate"`
	School           School        `json:"school" bson:"school"`
	PrevSchools      []School      `json:"prevSchools" bson:"prevSchools"`
	Roles            []Role        `json:"roles" bson:"roles"`
}

//Users user array
type Users []User

// Role role model
type Role struct {
	RoleType enums.UserRole `json:"roleType" bson:"roleType"`
	Stds     []Standard     `json:"stds" bson:"std"`
}

//Standard std model
type Standard struct {
	Std      uint16    `json:"std" bson:"std"`
	Subjects []Subject `json:"subjects" bson:"subjects"`
}

//Subject subject model
type Subject struct {
	Title    string      `json:"title" bson:"title"`
	Reviewer Contributor `json:"reviewer" bson:"reviewer"`
	Topics   []string    `json:"topics" bson:"topics"`
	Tags     []string    `json:"tags" bson:"tags"`
}
