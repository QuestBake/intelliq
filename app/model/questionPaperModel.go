package model

import (
	"intelliq/app/enums"
)

//QuestionPaper question paper model
type QuestionPaper struct {
	Set      int       `json:"set"`
	Sections []Section `json:"sections"`
}

//Section QuestionPaper sections
type Section struct {
	Type      enums.QuesLength `json:"type"`
	Questions Questions        `json:"questions"`
	LevelMap  map[enums.Difficulty][]Questions
}
