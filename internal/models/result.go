package models

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	Billing Billing `json:"billing"`
}

const (
	LOW_LOAD  = 1
	AVG_LOAD  = 2
	HIGH_LOAD = 3
)
