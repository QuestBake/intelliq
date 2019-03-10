package dto

import (
	"intelliq/app/model"
)

//TestDto request params
type TestDto struct {
	Template  *model.Template  `json:"template"`
	TestPaper *model.TestPaper `json:"testPaper"`
}
