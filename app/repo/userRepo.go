package repo

import (
	db "intelliq/app/config"
	"intelliq/app/enums"
	"intelliq/app/model"

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
