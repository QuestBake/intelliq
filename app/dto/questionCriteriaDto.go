package dto

import (
	"intelliq/app/enums"
)

//QuestionCriteriaDto criteria for generating question paper
type QuestionCriteriaDto struct {
	GroupCode        string           `json:"groupCode"`
	Standard         int              `json:"std"`
	Subject          string           `json:"subject"`
	Topics           []string         `json:"topics"`
	Tags             []string         `json:"tags"`
	Sets             int              `json:"sets"`
	Length           []QuesLength     `json:"length"`
	Difficulty       []QuesDifficulty `json:"difficulty"`
	SearchTerm       string           `json:"searchTerm"`
	Page             int              `json:"page"`
	NativeLength     []enums.QuesLength
	NativeDifficulty []enums.QuesDifficulty
}

//QuesLength sectional length - OBJECTIVE,SHORT,BRIEF,LONG
type QuesLength struct {
	Type  enums.QuesLength `json:"type"`
	Count int              `json:"count"`
	Marks int              `json:"marks"`
}

//QuesDifficulty difficulty level - EASY,MEDIUM,HARD
type QuesDifficulty struct {
	Level   enums.QuesDifficulty `json:"level"`
	Percent int                  `json:"percent"`
}

//GenerateNatives simplifies length 'n' difficulty arrays for db query
func (criteria *QuestionCriteriaDto) GenerateNatives() {
	for _, length := range criteria.Length {
		criteria.NativeLength = append(criteria.NativeLength, length.Type)
	}
	for _, difficulty := range criteria.Difficulty {
		criteria.NativeDifficulty = append(criteria.NativeDifficulty, difficulty.Level)
	}
}
