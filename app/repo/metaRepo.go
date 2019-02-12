package repo

import (
	db "intelliq/app/config"
	"intelliq/app/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type metaRepository struct {
	coll *mgo.Collection
}

//NewMetaRepository repo struct
func NewMetaRepository() *metaRepository {
	coll := db.GetCollection(db.COLL_META)
	if coll == nil {
		return nil
	}
	return &metaRepository{
		coll,
	}
}

func (repo *metaRepository) Save(meta *model.Meta) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Insert(meta)
	return err
}

func (repo *metaRepository) Update(meta *model.Meta) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Update(bson.M{"_id": meta.MetaID}, meta)
	return err
}

func (repo *metaRepository) Read() (*model.Meta, error) {
	defer db.CloseSession(repo.coll)
	var meta model.Meta
	err := repo.coll.Find(nil).One(&meta)
	if err != nil {
		return nil, err
	}
	return &meta, nil
}
