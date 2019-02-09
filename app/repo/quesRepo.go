package repo

import (
	db "intelliq/app/config"

	"github.com/globalsign/mgo"
)

type questionRepository struct {
	coll *mgo.Collection
}

//NewQuestionRepository repo struct
func NewQuestionRepository(groupCode string) *questionRepository {
	coll := db.GetCollection(groupCode + db.COLL_QUES)
	if coll == nil {
		return nil
	}
	return &questionRepository{
		coll,
	}
}
