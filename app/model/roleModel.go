package model

import (
	"intelliq/app/enums"
)

// Role role model
type Role struct {
	RoleType enums.RoleType `json:"roleType" bson:"roleType"`
	Stds     []Standard     `json:"stds" bson:"std"`
}
