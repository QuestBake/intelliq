package dto

import (
	"intelliq/app/model"
)

//QuesResponseDto question response dto
type QuesResponseDto struct {
	Records   int             `json:"records"`   // total count of records
	Questions model.Questions `json:"questions"` // resultant questions
}
