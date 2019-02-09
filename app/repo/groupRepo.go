package repo

import (
	db "intelliq/app/config"

	"github.com/globalsign/mgo"
)

type groupRepository struct {
	coll *mgo.Collection
}

//NewGroupRepository repo struct
func NewGroupRepository() *groupRepository {
	coll := db.GetCollection(db.COLL_GROUP)
	if coll == nil {
		return nil
	}
	return &groupRepository{
		coll,
	}
}
