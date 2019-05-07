package model

import (
	"strings"
	"time"

	"github.com/globalsign/mgo/bson"
)

//School school model
type School struct {
	SchoolID         bson.ObjectId `json:"schoolId" bson:"_id,omitempty"`
	ShortName        string        `json:"shortName" bson:"shortName"`
	FullName         string        `json:"fullName" bson:"fullName"`
	Code             string        `json:"code" bson:"code"`
	Address          Address       `json:"address" bson:"address"`
	Contact          Contact       `json:"contact" bson:"contact"`
	Board            string        `json:"board" bson:"board"`
	Group            Group         `json:"group" bson:"group"`
	PrevGroups       []Group       `json:"prevGroups" bson:"prevGroups"`
	Standards        []uint16      `json:"stds" bson:"stds"`
	CreateDate       time.Time     `json:"createDate" bson:"createDate"`
	LastModifiedDate time.Time     `json:"lastModifiedDate" bson:"lastModifiedDate"`
	RenewalDate      time.Time     `json:"renewalDate" bson:"renewalDate"`
	PrevUserRoles    []Role        `json:"prevUserRoles" bson:"prevUserRoles"` // to know teacher's role in previous school
	Schedule         schedule      `json:"schedule" bson:"schedule"`
}

//Address address model
type Address struct {
	Area      string `json:"area" bson:"area"`
	City      string `json:"city" bson:"city"`
	State     string `json:"state" bson:"state"`
	Pincode   string `json:"pincode" bson:"pincode"`
	Latitude  string `json:"latitude" bson:"latitude"`
	Longitude string `json:"longitude" bson:"longitude"`
}

//Contact contact model
type Contact struct {
	Landline []string `json:"landline" bson:"landline"`
	Mobile   []string `json:"mobile" bson:"mobile"`
	Email    string   `json:"email" bson:"email"`
	Website  string   `json:"website" bson:"website"`
}

type schedule struct {
	Days    int `json:"days" bson:"days"`
	Periods int `json:"periods" bson:"periods"`
}

//Schools school array
type Schools []School

//FormatAttributes formats school attributes
func (school *School) FormatAttributes() {
	school.ShortName = strings.ToUpper(school.ShortName)
	school.Code = school.ShortName + "_" + school.Address.Pincode
	school.Address.Area = strings.Title(school.Address.Area)
	school.Address.City = strings.Title(school.Address.City)
	school.Address.State = strings.ToUpper(school.Address.State)
	school.Board = strings.ToUpper(school.Board)
}
