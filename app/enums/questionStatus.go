package enums

//QuestionStatus question status
type QuestionStatus = int

type qStatus struct {
	NEW      QuestionStatus // newly added question
	TRANSIT  QuestionStatus // approved ques under update request
	REMOVE   QuestionStatus // approved ques under delete request
	APPROVED QuestionStatus // approved ques
	REJECTED QuestionStatus // rejected request for above cases
	PENDING  QuestionStatus // comprises of NEW,TRANSIT,REMOVE : sent as param for teacher to view all sorts of pending requests
	OBSOLETE QuestionStatus // dup copy is approved & is mainstream ;thus original becomes OBSOLETE to be later cleared by scheduler
}

// CurrentQuestionStatus for public use
var CurrentQuestionStatus = &qStatus{
	NEW:      0,
	TRANSIT:  1,
	REMOVE:   2,
	APPROVED: 3,
	REJECTED: 4,
	PENDING:  5,
	OBSOLETE: 6,
}
