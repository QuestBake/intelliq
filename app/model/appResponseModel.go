package model

import "intelliq/app/enums"

//AppResponse appResponse
type AppResponse struct {
	Status enums.ResponseStatus `json:"status"`
	Body   interface{}          `json:"body"`
	Msg    string               `json:"msg"`
}
