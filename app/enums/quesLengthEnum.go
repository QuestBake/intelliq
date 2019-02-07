package enums

//QuesLength enum question length
type QuesLength int

type length struct {
	OBJECTIVE QuesLength
	SHORT     QuesLength
	BRIEF     QuesLength
	LONG      QuesLength
}

// Length for public use
var Length = &length{
	OBJECTIVE: 0,
	SHORT:     1,
	BRIEF:     2,
	LONG:      3,
}
