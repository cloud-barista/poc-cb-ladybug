package model

const (
	STATUS_FAIL    = 0
	STATUS_SUCCESS = 1
	STATUS_UNKNOWN = 2
	//STATUS_NOT_EXIST = 404
	// STATUS_OK        = 200
)

type Status struct {
	Kind    string `json:"kind"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewStatus() *Status {
	return &Status{
		Kind: KIND_STATUS,
	}
}
