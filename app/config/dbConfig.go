package config

import (
	"fmt"
	"intelliq/app/common"
	"strings"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	url    = "localhost"
	port   = "27017"
	dbName = "intelliQ"
)

const (
	COLL_META     = "meta"
	COLL_USER     = "users"
	COLL_SCHOOL   = "schools"
	COLL_GROUP    = "groups"
	COLL_QUES     = "_questions"
	COLL_TEMPLATE = "_templates"
	COLL_PAPER    = "_papers"
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
	createIndices(session.Copy())
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
	if strings.HasPrefix(collName, common.GROUP_CODE_PREFIX) {
		collNames, err := db.CollectionNames()
		if err != nil {
			log.Error("Failed to get coll names:", err)
			return nil
		}
		collFound := false
		for _, name := range collNames {
			collFound = collName == name
			if collFound {
				break
			}
		}
		if !collFound {
			log.Error("collection with name: ", collName, "  does not exist ..")
			return nil
		}
	}
	coll := db.C(collName)
	if coll == nil {
		return nil
	}
	return coll
}

//CloseSession closes session
func CloseSession(coll *mgo.Collection) {
	if coll != nil {
		coll.Database.Session.Close()
	}
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
	addUniqueIndex(db, COLL_GROUP, []string{"code"})
	addUniqueIndex(db, COLL_SCHOOL, []string{"code"})
	addUniqueIndex(db, "GP_DPS"+COLL_TEMPLATE, []string{"criteria512Hash"})
	var searchFields []searchField
	searchFields = append(searchFields, searchField{field: "title", weight: 4})
	searchFields = append(searchFields, searchField{field: "topic", weight: 2})
	searchFields = append(searchFields, searchField{field: "tags", weight: 1})
	addSearchIndex(db, "GP_DPS"+COLL_QUES, searchFields)
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
	log.Info("Creating search index for " + collName)
	err := coll.EnsureIndex(index)
	if err != nil {
		panic("Could not create search index for " + collName + err.Error())
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
		fmt.Println("Creating unique index on := ", key, " for coll := ", collName)
		if err := coll.EnsureIndex(index); err != nil {
			panic("Could not create unique index for " + collName)
		}
	}
}

func addIndex(db *mgo.Database, collName string, fields []string) {
	coll := db.C(collName)
	if coll == nil {
		panic("No such Collection in DB" + collName)
	}
	for _, key := range fields {
		index := mgo.Index{
			Key:    []string{key},
			Sparse: true,
		}
		if err := coll.EnsureIndex(index); err != nil {
			panic("Could not create index for " + collName)
		}
	}
}
