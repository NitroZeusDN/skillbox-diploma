package service

import (
	"skillbox-diploma/internal/models"
	"sort"
)

const incidentURL = "/accendent"

type IncidentService struct {
	host string
}

func (s IncidentService) Get() ([]models.Incident, error) {
	data, err := models.GetIncident(s.host + incidentURL)
	if err != nil {
		return nil, nil
	}

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Status < data[j].Status
	})

	return data, nil
}

func NewIncidentService(host string) IncidentService {
	return IncidentService{
		host: host,
	}
}
