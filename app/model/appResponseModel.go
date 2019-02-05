package model

//AppResponse appResponse
type AppResponse struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
	Msg    string      `json:"msg"`
}
