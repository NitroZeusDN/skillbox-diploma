package service

import (
	"fmt"
	"skillbox-diploma/internal/models"
	"skillbox-diploma/internal/utils"
	"sort"
)

const mmsURL = "/mms"

type MMSService struct {
	host string
}

func (s MMSService) Get() ([][]models.MMS, error) {
	data, err := models.GetMMS(s.host + mmsURL)
	if err != nil {
		return nil, fmt.Errorf("getMMS failed: %w", err)
	}

	for i := range data {
		data[i].Country = utils.CodeToCountry(data[i].Country)
	}

	mmsCountrySort := data
	mmsProviderSort := make([]models.MMS, len(data))
	copy(mmsProviderSort, mmsCountrySort)

	sort.SliceStable(mmsCountrySort, func(a, b int) bool {
		return mmsCountrySort[a].Country < mmsCountrySort[b].Country
	})
	sort.SliceStable(mmsProviderSort, func(a, b int) bool {
		return mmsProviderSort[a].Provider < mmsProviderSort[b].Provider
	})

	return [][]models.MMS{
		mmsProviderSort,
		mmsCountrySort,
	}, nil
}

func NewMMSService(host string) MMSService {
	return MMSService{
		host: host,
	}
}
