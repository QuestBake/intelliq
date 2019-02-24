package repo

import (
	db "intelliq/app/config"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type questionRepository struct {
	coll *mgo.Collection
}

//NewQuestionRepository repo struct
func NewQuestionRepository(groupCode string) *questionRepository {
	coll := db.GetCollection(groupCode + db.COLL_QUES)
	if coll == nil {
		return nil
	}
	return &questionRepository{
		coll,
	}
}

func (repo *questionRepository) Save(question *model.Question) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Insert(question)
	return err
}

func (repo *questionRepository) Update(question *model.Question) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Update(bson.M{"_id": question.QuestionID}, question)
	return err
}

func (repo *questionRepository) Upsert(question *model.Question) error {
	defer db.CloseSession(repo.coll)
	_, err := repo.coll.Upsert(bson.M{"_id": question.QuestionID}, question)
	return err
}

func (repo *questionRepository) Delete(questionID bson.ObjectId) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Remove(bson.M{"_id": questionID})
	return err
}

func (repo *questionRepository) UpdateLimitedCols(question *model.Question) error {
	defer db.CloseSession(repo.coll)
	selector := bson.M{"_id": question.QuestionID}
	updator := bson.M{"$set": bson.M{"status": question.Status, "rejectDesc": question.RejectDesc,
		"lastModifiedDate": question.LastModifiedDate, "originId": question.OriginID}}
	err := repo.coll.Update(selector, updator)
	return err
}

func (repo *questionRepository) BulkUpdate(questions model.Questions) error {
	defer db.CloseSession(repo.coll)
	bulk := repo.coll.Bulk()
	for _, question := range questions {
		selector := bson.M{"_id": question.QuestionID}
		bulk.Update(selector, question)
	}
	_, err := bulk.Run()
	if err != nil {
		return err
	}
	return nil
}

func (repo *questionRepository) FindOne(quesID bson.ObjectId) (*model.Question, error) {
	defer db.CloseSession(repo.coll)
	var question model.Question
	filter := bson.M{
		"_id": quesID,
	}
	err := repo.coll.Find(filter).One(&question)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (repo *questionRepository) FindApprovedQuestions(quesRequestDto *dto.QuesRequestDto) (model.Questions, error) {
	defer db.CloseSession(repo.coll)
	filter := bson.M{
		"std":       quesRequestDto.Std,
		"subject":   quesRequestDto.Subject,
		"status":    enums.CurrentQuestionStatus.APPROVED,
		"owner._id": bson.M{"$ne": quesRequestDto.UserID},
	}
	return findAllRequests(repo, filter, quesRequestDto)
}

func (repo *questionRepository) FindReviewersRequests(quesRequestDto *dto.QuesRequestDto,
	status []enums.QuestionStatus) (model.Questions, error) {
	defer db.CloseSession(repo.coll)
	filter := bson.M{
		"school._id":   quesRequestDto.SchoolID,
		"reviewer._id": quesRequestDto.UserID,
		"status": bson.M{
			"$in": status,
		},
	}
	return findAllRequests(repo, filter, quesRequestDto)
}

func (repo *questionRepository) FindTeachersRequests(quesRequestDto *dto.QuesRequestDto,
	status []enums.QuestionStatus) (model.Questions, error) {
	defer db.CloseSession(repo.coll)
	filter := bson.M{
		"school._id": quesRequestDto.SchoolID,
		"std":        quesRequestDto.Std,
		"subject":    quesRequestDto.Subject,
		"owner._id":  quesRequestDto.UserID,
		"status": bson.M{
			"$in": status,
		},
	}
	return findAllRequests(repo, filter, quesRequestDto)
}

func findAllRequests(repo *questionRepository, filter bson.M, quesRequestDto *dto.QuesRequestDto) (model.Questions, error) {
	//	cols := bson.M{"_id": 1, "title": 1, "status": 1, "std": 1, "subject": 1, "lastModifiedDate": 1}
	skip := quesRequestDto.Page * quesRequestDto.Limit
	var questions model.Questions
	err := repo.coll.Find(filter).Sort("lastModifiedDate").
		Limit(quesRequestDto.Limit).Skip(skip).All(&questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}
