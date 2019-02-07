package model

import (
	"project/intelliq/app/enums"
)

// Role role model
type Role struct {
	RoleType enums.RoleType `json:"roleType" bson:"roleType"`
	Std      []Standard     `json:"std" bson:"std"`
}
