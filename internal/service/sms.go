package service

import (
	"fmt"
	"skillbox-diploma/internal/models"
	"skillbox-diploma/internal/utils"
	"sort"
)

type SMSService struct {
	path string
}

func (s SMSService) Get() ([][]models.SMS, error) {
	data, err := models.GetSMS(s.path)
	if err != nil {
		return nil, fmt.Errorf("getSMS failed: %w", err)
	}

	for i := range data {
		data[i].Country = utils.CodeToCountry(data[i].Country)
	}

	smsCountrySort := data
	smsProviderSort := make([]models.SMS, len(data))
	copy(smsProviderSort, smsCountrySort)

	sort.SliceStable(smsCountrySort, func(i, j int) bool {
		return smsCountrySort[i].Country < smsCountrySort[j].Country
	})
	sort.SliceStable(smsProviderSort, func(i, j int) bool {
		return smsProviderSort[i].Provider < smsProviderSort[j].Provider
	})

	return [][]models.SMS{
		smsProviderSort,
		smsCountrySort,
	}, nil
}

func NewSMSService(path string) SMSService {
	return SMSService{
		path: path,
	}
}
