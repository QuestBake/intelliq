package repo

import (
	db "intelliq/app/config"

	"github.com/globalsign/mgo"
)

type schoolRepository struct {
	coll *mgo.Collection
}

//NewSchoolRepository repo struct
func NewSchoolRepository() *schoolRepository {
	coll := db.GetCollection(db.COLL_SCHOOL)
	if coll == nil {
		return nil
	}
	return &schoolRepository{
		coll,
	}
}
