package enums

//UserRole enum roletype
type UserRole int

type roletype struct {
	SUPER    UserRole
	GROUP    UserRole
	SCHOOL   UserRole
	REVIEWER UserRole
	TEACHER  UserRole
}

// Role for public use
var Role = &roletype{
	SUPER:    0,
	GROUP:    1,
	SCHOOL:   2,
	REVIEWER: 3,
	TEACHER:  4,
}
