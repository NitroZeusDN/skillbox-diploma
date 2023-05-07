package service

import (
	"fmt"
	"skillbox-diploma/internal/models"
)

const supportURL = "/support"

const (
	lowLoad  = 1
	avgLoad  = 2
	highLoad = 3
)

type SupportService struct {
	host string
}

func (s SupportService) Get() ([]int, error) {
	supports, err := models.GetSupport(s.host + supportURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get support supports: %w", err)
	}

	if len(supports) == 0 {
		return nil, fmt.Errorf("empty supports array")
	}

	activeTickets := 0
	for _, support := range supports {
		activeTickets += support.ActiveTickets
	}

	avgTime := activeTickets * 60 / 18
	load := lowLoad
	switch {
	case activeTickets >= 9 && activeTickets <= 16:
		load = avgLoad
	case activeTickets > 16:
		load = highLoad
	}

	return []int{load, avgTime}, nil
}

func NewSupportService(host string) SupportService {
	return SupportService{
		host: host,
	}
}
