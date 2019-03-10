package enums

//TestPaperStatus question status
type TestPaperStatus = int

type tStatus struct {
	DRAFT   TestPaperStatus // draft paper
	RELEASE TestPaperStatus // saved/downloaded paper
}

// CurrentTestStatus for public use
var CurrentTestStatus = &tStatus{
	DRAFT:   0,
	RELEASE: 1,
}
