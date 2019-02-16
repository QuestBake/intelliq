package repo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	db "intelliq/app/config"
	"intelliq/app/model"
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

func (repo *groupRepository) Save(group *model.Group) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Insert(group)
	return err
}

func (repo *groupRepository) Update(group *model.Group) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Update(bson.M{"_id": group.GroupID}, group)
	return err
}

func (repo *groupRepository) FindAll(restrict int) (model.Groups, error) {
	defer db.CloseSession(repo.coll)
	var groups model.Groups
	var err error
	if restrict > 0 {
		cols := bson.M{"_id": 1, "code": 1}
		err = repo.coll.Find(nil).Select(cols).All(&groups)
	} else {
		err = repo.coll.Find(nil).All(&groups)
	}
	if err != nil {
		return nil, err
	}
	return groups, nil
}

func (repo *groupRepository) FindOne(key string, val interface{}) (*model.Group, error) {
	defer db.CloseSession(repo.coll)
	var group model.Group
	filter := bson.M{
		key: val,
	}
	err := repo.coll.Find(filter).One(&group)
	if err != nil {
		return nil, err
	}
	return &group, nil
}
