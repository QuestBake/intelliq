package repo

import (
	db "intelliq/app/config"
	"intelliq/app/enums"
	"intelliq/app/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type testPaperRepository struct {
	coll *mgo.Collection
}

//NewTestPaperRepository repo struct
func NewTestPaperRepository(groupCode string) *testPaperRepository {
	coll := db.GetCollection(groupCode + db.COLL_PAPER)
	if coll == nil {
		return nil
	}
	return &testPaperRepository{
		coll,
	}
}

func (repo *testPaperRepository) Save(testPaper *model.TestPaper) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Insert(testPaper)
	return err
}

func (repo *testPaperRepository) Update(testPaper *model.TestPaper) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Update(bson.M{"_id": testPaper.TestID}, testPaper)
	return err
}

func (repo *testPaperRepository) FindAll(teacherID bson.ObjectId) (model.TestPapers, error) {
	defer db.CloseSession(repo.coll)
	var drafts model.TestPapers
	filter := bson.M{
		"teacherId": teacherID,
		"status":    enums.CurrentTestStatus.DRAFT,
	}
	cols := bson.M{"_id": 1, "tag": 1, "lastModifiedDate": 1}
	err := repo.coll.Find(filter).Select(cols).All(&drafts)
	if err != nil {
		return nil, err
	}
	return drafts, nil
}

func (repo *testPaperRepository) FindOne(draftID bson.ObjectId) (*model.TestPaper, error) {
	defer db.CloseSession(repo.coll)
	var testPaper model.TestPaper
	filter := bson.M{
		"_id": draftID,
	}
	err := repo.coll.Find(filter).One(&testPaper)
	if err != nil {
		return nil, err
	}
	return &testPaper, nil
}
