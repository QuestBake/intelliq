package model

import "github.com/globalsign/mgo/bson"

//Meta meta data model
type Meta struct {
	MetaID    bson.ObjectId `json:"metaId" bson:"_id,omitempty"`
	Subjects  []string      `json:"subjects" bson:"subjects"`
	Standards []uint16      `json:"standards" bson:"standards"`
}
