package models

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]SMS         `json:"sms"`
	MMS       [][]MMS         `json:"mms"`
	Billing   Billing         `json:"billing"`
	Email     EmailCollection `json:"email"`
	VoiceCall []Voice         `json:"voice_call"`
}

const (
	LOW_LOAD  = 1
	AVG_LOAD  = 2
	HIGH_LOAD = 3
)
