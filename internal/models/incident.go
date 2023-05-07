package models

import (
	"encoding/json"
	"io"
	"net/http"
	"skillbox-diploma/internal/utils"
)

// Incident данные по инцеденту.
type Incident struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

// GetIncident получает данные по инцидентам по указанному URL.
func GetIncident(addr string) ([]Incident, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := make([]Incident, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	filteredData := make([]Incident, 0)
	for _, incident := range data {
		switch {
		case !utils.IsValidStatus(incident.Status):
		default:
			filteredData = append(filteredData, incident)
		}
	}

	return filteredData, nil
}
