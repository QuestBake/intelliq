package repo

import (
	db "intelliq/app/config"

	"github.com/globalsign/mgo"
)

type questionRepository struct {
	coll *mgo.Collection
}

//NewQuestionRepository repo struct
func NewQuestionRepository(quesCollName string) *questionRepository {
	coll := db.GetCollection(quesCollName)
	if coll == nil {
		return nil
	}
	return &questionRepository{
		coll,
	}
}
