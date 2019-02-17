package enums

//QuestionStatus question status
type QuestionStatus = int

type qStatus struct {
	PENDING  QuestionStatus
	REJECTED QuestionStatus
	APPROVED QuestionStatus
	TRANSIT  QuestionStatus
}

// CurrentStatus for public use
var CurrentStatus = &qStatus{
	PENDING:  0,
	REJECTED: 1,
	APPROVED: 2,
	TRANSIT:  3,
}
