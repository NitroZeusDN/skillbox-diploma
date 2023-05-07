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
	Incidents []Incident      `json:"incident"`
	Support   []int           `json:"support"`
}
