package dto

import "intelliq/app/enums"

//AppResponseDto appResponse
type AppResponseDto struct {
	Status enums.ResponseStatus `json:"status"`
	Body   interface{}          `json:"body"`
	Msg    string               `json:"msg"`
}
