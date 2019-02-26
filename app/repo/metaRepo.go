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
	bulk := repo.coll.Bulk()
	selector := bson.M{"_id": meta.MetaID}
	topicUpdator := bson.M{"$addToSet": bson.M{"subjects": bson.M{"$each": meta.Subjects}}}
	tagUpdator := bson.M{"$addToSet": bson.M{"standards": bson.M{"$each": meta.Standards}}}
	bulk.Update(selector, topicUpdator)
	bulk.Update(selector, tagUpdator)
	_, err := bulk.Run()
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

func (repo *metaRepository) Remove(meta *model.Meta) error {
	defer db.CloseSession(repo.coll)
	bulk := repo.coll.Bulk()
	selector := bson.M{"_id": meta.MetaID}
	subjectUpdator := bson.M{"$pull": bson.M{"subjects": bson.M{"$in": meta.Subjects}}}
	standardUpdator := bson.M{"$pull": bson.M{"standards": bson.M{"$in": meta.Standards}}}
	bulk.Update(selector, subjectUpdator)
	bulk.Update(selector, standardUpdator)
	_, err := bulk.Run()
	if err != nil {
		return err
	}
	return nil
}
