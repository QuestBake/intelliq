package enums

//RoleType enum roletype
type RoleType int

type roletype struct {
	SUPER    RoleType
	GROUP    RoleType
	SCHOOL   RoleType
	APPROVER RoleType
	TEACHER  RoleType
}

// Role for public use
var Role = &roletype{
	SUPER:    0,
	GROUP:    1,
	SCHOOL:   2,
	APPROVER: 3,
	TEACHER:  4,
}
