package model

// Role role model
type Role struct {
	RoleType string     `json:"roleType" bson:"roleType"`
	Std      []Standard `json:"std" bson:"std"`
}
