package repo

import (
	"strconv"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"intelliq/app/common"
	db "intelliq/app/config"
	"intelliq/app/enums"
	"intelliq/app/model"
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

func (repo *userRepository) Save(user *model.User) error {
	defer db.CloseSession(repo.coll)
	err := repo.coll.Insert(user)
	return err
}

func (repo *userRepository) Update(user *model.User) error {
	defer db.CloseSession(repo.coll)
	selector := bson.M{"_id": user.UserID}
	updator := bson.M{"$set": bson.M{"name": user.FullName, "gender": user.Gender,
		"email": user.Email, "dob": user.DOB, "school": user.School,
		"roles": user.Roles, "lastModifiedDate": user.LastModifiedDate}}
	err := repo.coll.Update(selector, updator)
	return err
}

func (repo *userRepository) UpdateMobilePwd(selectorField string,
	updatorField string, selectorVal interface{}, updatorVal string) error {
	defer db.CloseSession(repo.coll)
	selector := bson.M{selectorField: selectorVal}
	updator := bson.M{"$set": bson.M{updatorField: updatorVal, "lastModifiedDate": time.Now().UTC()}}
	err := repo.coll.Update(selector, updator)
	return err
}

func (repo *userRepository) FindAllSchoolAdmins(groupID bson.ObjectId) (model.Users, error) {
	defer db.CloseSession(repo.coll)
	var users model.Users
	filter := bson.M{
		"school.group._id": groupID,
		"roles.roleType":   enums.Role.SCHOOL,
	}
	cols := bson.M{"password": 0, "prevSchools": 0, "days": 0, "lastModifiedDate": 0}
	err := repo.coll.Find(filter).Select(cols).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userRepository) FindAllSchoolTeachers(schoolID bson.ObjectId,
	roleType interface{}) (model.Users, error) {
	defer db.CloseSession(repo.coll)
	var users model.Users
	var filter interface{}
	if roleType == nil {
		filter = bson.M{"school._id": schoolID}
	} else {
		filter = bson.M{"school._id": schoolID, "roles.roleType": roleType}
	}
	cols := bson.M{"password": 0, "prevSchools": 0, "days": 0, "lastModifiedDate": 0, "school": 0}
	err := repo.coll.Find(filter).Select(cols).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userRepository) FindAllteachersUnderReviewer(schoolID bson.ObjectId,
	reviewerID bson.ObjectId) (model.Users, error) {
	defer db.CloseSession(repo.coll)
	var users model.Users
	filter := bson.M{
		"school._id":                     schoolID,
		"roles.roleType":                 enums.Role.TEACHER,
		"roles.std.subjects.approver_id": reviewerID,
	}
	cols := bson.M{"password": 0, "prevSchools": 0, "days": 0, "lastModifiedDate": 0, "school": 0}
	err := repo.coll.Find(filter).Select(cols).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userRepository) TransferRole(roleType enums.UserRole,
	fromUserID bson.ObjectId, toUserID bson.ObjectId) (string, []string, error) {
	defer db.CloseSession(repo.coll)
	var users model.Users
	fromUserFilter := bson.M{
		"_id":            fromUserID,
		"roles.roleType": roleType,
	}
	toUserFilter := bson.M{
		"_id":            toUserID,
		"roles.roleType": bson.M{"$ne": roleType},
	}
	orFilter := bson.M{"$or": []bson.M{
		fromUserFilter,
		toUserFilter,
	},
	}
	cols := bson.M{"_id": 1, "roles": 1, "mobile": 1}
	err := repo.coll.Find(orFilter).Select(cols).All(&users)
	if err != nil {
		return "", nil, err
	}
	count := len(users)
	if count < 2 {
		return common.MSG_INSUFFICIENT_USER_COUNT + strconv.Itoa(count), nil, nil
	}
	bulk := repo.coll.Bulk()
	for _, user := range users {
		if user.UserID == fromUserID { // remove the current role from role array
			for index, role := range user.Roles {
				if role.RoleType == roleType {
					user.Roles = append(user.Roles[:index], //appends records before this point
						user.Roles[index+1:]...) // appends records after this point
					break
				}
			}
		} else { //add new role to role array
			user.Roles = append(user.Roles, model.Role{
				RoleType: roleType,
			})
		}
		user.LastModifiedDate = time.Now().UTC()
		selector := bson.M{"_id": user.UserID}
		updator := bson.M{"$set": bson.M{"roles": user.Roles,
			"lastModifiedDate": user.LastModifiedDate}}
		bulk.Update(selector, updator)
	}
	_, errs := bulk.Run()
	if errs != nil {
		return "", nil, errs
	}
	mobiles := []string{users[0].Mobile, users[1].Mobile}
	return "", mobiles, nil
}

func (repo *userRepository) RemoveSchoolTeacher(schoolID bson.ObjectId,
	userID bson.ObjectId) error {
	defer db.CloseSession(repo.coll)
	var user model.User
	filter := bson.M{"_id": userID, "school._id": schoolID}
	cols := bson.M{"_id": 1, "roles": 1, "school": 1, "prevSchools": 1}
	err := repo.coll.Find(filter).Select(cols).One(&user)
	if err != nil {
		return err
	}
	user.School.PrevUserRoles = user.Roles
	user.PrevSchools = append(user.PrevSchools, user.School)
	user.School = model.School{}
	user.Roles = nil
	user.LastModifiedDate = time.Now().UTC()
	selector := bson.M{"_id": user.UserID}
	updator := bson.M{"$set": bson.M{"school": user.School, "prevSchools": user.PrevSchools,
		"roles": user.Roles, "lastModifiedDate": user.LastModifiedDate}}
	errs := repo.coll.Update(selector, updator)
	if errs != nil {
		return err
	}
	return nil
}

func (repo *userRepository) BulkSave(users []interface{}) error {
	defer db.CloseSession(repo.coll)
	bulk := repo.coll.Bulk()
	bulk.Insert(users...)
	_, err := bulk.Run()
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) BulkUpdate(users model.Users) error {
	defer db.CloseSession(repo.coll)
	bulk := repo.coll.Bulk()
	for _, user := range users {
		selector := bson.M{"_id": user.UserID}
		//	updator := bson.M{"$set": bson.M{"roles": user.Roles, "lastModifiedDate": time.Now().UTC()}}
		bulk.Update(selector, user)
	}
	_, err := bulk.Run()
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) FindOne(key string, val interface{}) (*model.User, error) {
	defer db.CloseSession(repo.coll)
	var user model.User
	filter := bson.M{
		key: val,
	}
	err := repo.coll.Find(filter).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepository) UpdateSchedule(user *model.User) error {
	defer db.CloseSession(repo.coll)
	selector := bson.M{"_id": user.UserID}
	updator := bson.M{"$set": bson.M{"days": user.Days}}
	err := repo.coll.Update(selector, updator)
	return err
}
