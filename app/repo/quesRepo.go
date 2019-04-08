package repo

import (
	"fmt"
	"intelliq/app/common"
	db "intelliq/app/config"
	"intelliq/app/dto"
	"intelliq/app/enums"
	"intelliq/app/helper"
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

func (repo *questionRepository) Delete(questionID bson.ObjectId) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Remove(bson.M{"_id": questionID})
	return err
}

func (repo *questionRepository) UpdateLimitedCols(question *model.Question) error {
	defer db.CloseSession(repo.coll)
	selector := bson.M{"_id": question.QuestionID}
	updator := bson.M{"$set": bson.M{"status": question.Status,
		"rejectDesc": question.RejectDesc, "originId": question.OriginID,
		"lastModifiedDate": question.LastModifiedDate}}
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

func (repo *questionRepository) BulkSave(questions model.Questions) error {
	defer db.CloseSession(repo.coll)
	bulk := repo.coll.Bulk()
	for _, ques := range questions {
		bulk.Insert(ques)
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
		"status": bson.M{
			"$ne": enums.CurrentQuestionStatus.OBSOLETE,
		},
	}
	err := repo.coll.Find(filter).One(&question)
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (repo *questionRepository) FindApprovedQuestions(
	quesRequestDto *dto.QuesRequestDto) (model.Questions, error) {
	defer db.CloseSession(repo.coll)
	filter := bson.M{
		"status": enums.CurrentQuestionStatus.APPROVED,
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
		"owner._id":  quesRequestDto.UserID,
		"status": bson.M{
			"$in": status,
		},
	}
	return findAllRequests(repo, filter, quesRequestDto)
}

func (repo *questionRepository) FilterQuestionsForPaper(
	quesCriteriaDto *dto.QuestionCriteriaDto) (map[enums.QuesLength]map[enums.QuesDifficulty]model.Questions,
	error) {
	defer db.CloseSession(repo.coll)
	filter := bson.M{
		"std":     quesCriteriaDto.Standard,
		"subject": quesCriteriaDto.Subject,
		"status":  enums.CurrentQuestionStatus.APPROVED,
		"topic": bson.M{
			"$in": quesCriteriaDto.Topics,
		},
		"length": bson.M{
			"$in": quesCriteriaDto.NativeLength,
		},
		"difficulty": bson.M{
			"$in": quesCriteriaDto.NativeDifficulty,
		},
	}
	cols := bson.M{"_id": 1, "title": 1, "difficulty": 1, "length": 1,
		"topic": 1, "tags": 1, "imageUrl": 1}
	var ques model.Question
	itr := repo.coll.Find(filter).Batch(common.QUES_BATCH_SIZE).Select(cols).Iter()
	sectionQuesMap := make(map[enums.QuesLength]map[enums.QuesDifficulty]model.Questions)
	ctr := 0
	for itr.Next(&ques) {
		helper.PopulateSectionalQuestionMap(&ques, sectionQuesMap)
		ctr++
	}
	fmt.Println("Records: ", ctr)
	if err := itr.Close(); err != nil {
		return nil, err
	}
	return sectionQuesMap, nil
}

func (repo *questionRepository) FilterQuestionsPerCriteria(
	quesCriteriaDto *dto.QuestionCriteriaDto) (model.Questions, error) {
	defer db.CloseSession(repo.coll)
	var questions model.Questions
	filter := bson.M{
		"std":        quesCriteriaDto.Standard,
		"subject":    quesCriteriaDto.Subject,
		"status":     enums.CurrentQuestionStatus.APPROVED,
		"topic":      quesCriteriaDto.Topics[0],
		"length":     quesCriteriaDto.NativeLength[0],
		"difficulty": quesCriteriaDto.NativeDifficulty[0],
		"tags": bson.M{
			"$in": quesCriteriaDto.Tags,
		},
	}
	cols := bson.M{"_id": 1, "title": 1, "difficulty": 1, "length": 1,
		"topic": 1, "tags": 1, "imageUrl": 1}
	err := repo.coll.Find(filter).Select(cols).Limit(common.DEF_REQUESTS_PAGE_SIZE).
		Skip(quesCriteriaDto.Page).All(&questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (repo *questionRepository) FilterQuestionsPerSearchTerm(
	quesCriteriaDto *dto.QuestionCriteriaDto) (model.Questions, error) {
	defer db.CloseSession(repo.coll)
	var questions model.Questions
	filter := bson.M{
		"std":     quesCriteriaDto.Standard,
		"subject": quesCriteriaDto.Subject,
		"$text": bson.M{
			"$search": quesCriteriaDto.SearchTerm,
		},
	}
	cols := bson.M{"_id": 0, "title": 1}
	err := repo.coll.Find(filter).Select(cols).Sort("-_id").
		Limit(common.DEF_REQUESTS_PAGE_SIZE).Skip(quesCriteriaDto.Page *
		common.DEF_REQUESTS_PAGE_SIZE).All(&questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func findAllRequests(repo *questionRepository, filter bson.M,
	quesRequestDto *dto.QuesRequestDto) (model.Questions, error) {
	//	cols := bson.M{"_id": 1, "title": 1, "status": 1, "std": 1, "subject": 1, "lastModifiedDate": 1}
	filter = createStdSubjectFilter(filter, quesRequestDto.Standards)
	skip := quesRequestDto.Page * common.DEF_REQUESTS_PAGE_SIZE
	var questions model.Questions
	err := repo.coll.Find(filter).Sort("lastModifiedDate").
		Limit(common.DEF_REQUESTS_PAGE_SIZE).Skip(skip).All(&questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func createStdSubjectFilter(basicFilter bson.M, standards []model.Standard) bson.M {
	var stdFilters []bson.M
	for _, standard := range standards {
		subjects := getStringSubjects(standard.Subjects)
		if len(subjects) > 0 {
			subFilter := bson.M{"$and": []bson.M{
				{"std": standard.Std},
				{"subject": bson.M{
					"$in": subjects},
				},
			},
			}
			stdFilters = append(stdFilters, subFilter)
		}
	}
	orStdFilter := bson.M{"$or": stdFilters}
	mainFilter := bson.M{"$and": []bson.M{
		basicFilter,
		orStdFilter,
	},
	}
	fmt.Println(mainFilter)
	return mainFilter
}

func getStringSubjects(subjects []model.Subject) []string {
	var subs []string
	for _, subject := range subjects {
		subs = append(subs, subject.Title)
	}
	return subs
}
