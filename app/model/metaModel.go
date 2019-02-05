package model

//Meta meta data model
type Meta struct {
	Subjects  []string `json:"subjects" bson:"subjects"`
	Standards []uint16 `json:"standards" bson:"standards"`
}
