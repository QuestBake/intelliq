package enums

//ResponseStatus codes
type ResponseStatus int

type status struct {
	SUCCESS ResponseStatus
	ERROR   ResponseStatus
}

//Status AppResonseStatus codes
var Status = &status{
	SUCCESS: 204,
	ERROR:   402,
}
