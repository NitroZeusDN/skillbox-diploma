package models

import (
	"encoding/json"
	"io"
	"net/http"
	"skillbox-diploma/internal/utils"
)

// MMS данные о ммс.
type MMS struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

// GetMMS получает данные о MMS по указанному адресу. Получает и проверяет данные.
func GetMMS(addr string) ([]MMS, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := make([]MMS, 0)
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	filteredData := make([]MMS, 0)
	for _, mms := range data {
		switch {
		case !utils.IsValidCountry(mms.Country):
			fallthrough
		case !utils.IsValidBandwidth(mms.Bandwidth):
			fallthrough
		case !utils.IsValidResponseTime(mms.ResponseTime):
			fallthrough
		case !utils.IsValidSMSProvider(mms.Provider):
		default:
			filteredData = append(filteredData, mms)
		}
	}

	return filteredData, nil
}
