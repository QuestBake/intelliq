package model

import "project/intelliq/app/enums"

//AppResponse appResponse
type AppResponse struct {
	Status enums.ResponseStatus `json:"status"`
	Body   interface{}          `json:"body"`
	Msg    string               `json:"msg"`
}
