package models

import (
	"encoding/json"
	"io"
	"net/http"
)

type Support struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

// GetSupport получает данные о поддержке по указанному URL.
func GetSupport(addr string) ([]Support, error) {
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

	data := make([]Support, 0)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
