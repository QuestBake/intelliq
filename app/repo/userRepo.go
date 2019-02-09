package repo

import (
	db "intelliq/app/config"

	"github.com/globalsign/mgo"
)

type userRepository struct {
	coll *mgo.Collection
}

//NewUserRepository repo struct
func NewUserRepository() *userRepository {
	coll := db.GetCollection(db.COLL_USER)
	if coll == nil {
		return nil
	}
	return &userRepository{
		coll,
	}
}
