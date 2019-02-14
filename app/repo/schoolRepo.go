package repo

import (
	db "intelliq/app/config"
	"intelliq/app/model"

	"github.com/globalsign/mgo/bson"

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

func (repo *schoolRepository) Save(school *model.School) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Insert(school)
	return err
}

func (repo *schoolRepository) Update(school *model.School) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Update(bson.M{"_id": school.SchoolID}, school)
	return err
}

func (repo *schoolRepository) FindAll(key string, val interface{}) (model.Schools, error) {
	defer db.CloseSession(repo.coll)
	var schools model.Schools
	filter := bson.M{
		key: val,
	}
	err := repo.coll.Find(filter).All(&schools)
	if err != nil {
		return nil, err
	}
	return schools, nil
}
