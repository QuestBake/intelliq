package repo

import (
	db "project/intelliq/app/config"

	"github.com/globalsign/mgo"
)

type metaRepository struct {
	coll *mgo.Collection
}

//NewMetaRepository repo struct
func NewMetaRepository() *metaRepository {
	coll := db.GetCollection("meta")
	if coll == nil {
		return nil
	}
	return &metaRepository{
		coll,
	}
}
