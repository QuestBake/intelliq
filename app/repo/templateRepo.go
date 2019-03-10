package repo

import (
	db "intelliq/app/config"
	"intelliq/app/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type templateRepository struct {
	coll *mgo.Collection
}

//NewTemmplateRepository repo struct
func NewTemmplateRepository(groupCode string) *templateRepository {
	coll := db.GetCollection(groupCode + db.COLL_TEMPLATE)
	if coll == nil {
		return nil
	}
	return &templateRepository{
		coll,
	}
}

func (repo *templateRepository) Save(template *model.Template) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Insert(template)
	return err
}

func (repo *templateRepository) Update(template *model.Template) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Update(bson.M{"_id": template.TemplateID}, template)
	return err
}

func (repo *templateRepository) FindAll(teacherID bson.ObjectId) (model.Templates, error) {
	defer db.CloseSession(repo.coll)
	var templates model.Templates
	filter := bson.M{
		"teacherId": teacherID,
	}
	cols := bson.M{"_id": 1, "tag": 1, "lastModifiedDate": 1}
	err := repo.coll.Find(filter).Select(cols).All(&templates)
	if err != nil {
		return nil, err
	}
	return templates, nil
}

func (repo *templateRepository) FindOne(templateID bson.ObjectId) (*model.Template, error) {
	defer db.CloseSession(repo.coll)
	var template model.Template
	filter := bson.M{
		"_id": templateID,
	}
	err := repo.coll.Find(filter).One(&template)
	if err != nil {
		return nil, err
	}
	return &template, nil
}
