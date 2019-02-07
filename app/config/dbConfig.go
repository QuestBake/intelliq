package config

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	url    = "localhost"
	port   = "27017"
	dbName = "intelliQ"
)

var dbSession *mgo.Session

//Connect db conn
func Connect() (*mgo.Session, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Info("Successfully connected to DB at ", url)
	dbSession = session
	//	createIndices(session.Copy())
	return session, nil
}

//GetCollection copy of original session
func GetCollection(collName string) *mgo.Collection {
	if dbSession == nil {
		return nil
	}
	session := dbSession.Copy()
	db := session.DB(dbName)
	if db == nil {
		return nil
	}
	coll := db.C(collName)
	if coll == nil {
		return nil
	}
	return coll
}

type searchField struct {
	field  string
	weight int
}

func createIndices(session *mgo.Session) {
	db := session.DB(dbName)
	if db == nil {
		panic("No DB session")
	}

	var searchFields []searchField
	searchFields = append(searchFields, searchField{field: "city", weight: 4})
	searchFields = append(searchFields, searchField{field: "state", weight: 2})

	addSearchIndex(db, "addresses", searchFields)
	addUniqueIndex(db, "addresses", []string{"city"})
	db.Session.Close()
}

func addSearchIndex(db *mgo.Database, collName string, searchFields []searchField) {
	coll := db.C(collName)
	if coll == nil {
		panic("No such Collection in DB" + collName)
	}
	var fields []string
	weights := make(map[string]int)

	for _, val := range searchFields {
		fields = append(fields, "$text:"+val.field)
		weights[val.field] = val.weight
	}
	index := mgo.Index{
		Key:     fields,
		Weights: weights,
		Name:    "textIndex",
	}
	err := coll.EnsureIndex(index)
	if err != nil {
		panic("Hi Could not create search index for " + collName + err.Error())
	}
}

func addUniqueIndex(db *mgo.Database, collName string, fields []string) {
	coll := db.C(collName)
	if coll == nil {
		panic("No such Collection in DB" + collName)
	}
	for _, key := range fields {
		index := mgo.Index{
			Key:    []string{key},
			Unique: true,
		}
		if err := coll.EnsureIndex(index); err != nil {
			panic("Could not create unique index for " + collName)
		}
	}
}
