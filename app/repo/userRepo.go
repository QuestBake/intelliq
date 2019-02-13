package repo

import (
	db "intelliq/app/config"
	"intelliq/app/enums"
	"intelliq/app/model"
	"strconv"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
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
	err := repo.coll.Update(bson.M{"_id": user.UserID}, user)
	return err
}

func (repo *userRepository) FindAllSchoolAdmins(_id bson.ObjectId) (model.Users, error) {
	defer db.CloseSession(repo.coll)
	var users model.Users
	filter := bson.M{"$and": []bson.M{
		{
			"school.group._id": _id,
		},
		{
			"roles.roleType": enums.Role.SCHOOL,
		},
	},
	}
	err := repo.coll.Find(filter).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *userRepository) TransferRole(roleType enums.RoleType, fromUserID bson.ObjectId, toUserID bson.ObjectId) (string, error) {
	defer db.CloseSession(repo.coll)
	var users model.Users
	fromUserFilter := bson.M{"$and": []bson.M{
		{
			"_id": fromUserID,
		},
		{
			"roles.roleType": roleType,
		},
	},
	}
	toUserFilter := bson.M{"$and": []bson.M{
		{
			"_id": toUserID,
		},
		{
			"roles.roleType": bson.M{"$ne": roleType},
		},
	},
	}
	orFilter := bson.M{"$or": []bson.M{
		fromUserFilter,
		toUserFilter,
	},
	}
	cols := bson.M{"_id": 1, "roles": 1}

	err := repo.coll.Find(orFilter).Select(cols).All(&users)
	if err != nil {
		return "", err
	}
	count := len(users)
	if count < 2 {
		return "Expected 2 users, but found " + strconv.Itoa(count), nil
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
		updator := bson.M{"$set": bson.M{"roles": user.Roles, "lastModifiedDate": user.LastModifiedDate}}
		bulk.Update(selector, updator)
	}
	_, errs := bulk.Run()
	if errs != nil {
		return "", err
	}
	return "", nil
}
