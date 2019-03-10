package dto

import (
	"intelliq/app/enums"
	"intelliq/app/model"
)

//QuestionPaperDto question paper model
type QuestionPaperDto struct {
	Set      int       `json:"set"`
	Sections []Section `json:"sections"`
}

//Section QuestionPaperDto sections
type Section struct {
	Type      enums.QuesLength `json:"type"`
	Questions model.Questions  `json:"questions"`
	LevelMap  map[enums.QuesDifficulty][]model.Questions
}
