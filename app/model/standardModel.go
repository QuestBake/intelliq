package model

//Standard std model
type Standard struct {
	Std      uint16    `json:"std" bson:"std"`
	Subjects []Subject `json:"subjects" bson:"subjects"`
}

//Subject subject model
type Subject struct {
	Title    string   `json:"std" bson:"std"`
	Approver User     `json:"approver" bson:"approver"`
	Topics   []string `json:"topics" bson:"topics"`
}
