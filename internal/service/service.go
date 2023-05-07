package service

import (
	"fmt"
	"skillbox-diploma/internal/models"
)

type Service struct {
	tmp string

	billing BillingService
}

func (s *Service) Get() (models.ResultSetT, error) {
	var (
		res models.ResultSetT
		err error
	)

	res.Billing, err = s.billing.Get()
	if err != nil {
		return models.ResultSetT{}, fmt.Errorf("failed to get billing data: %w", err)
	}

	return res, nil
}

func New(tmp string) *Service {
	return &Service{
		tmp:     tmp,
		billing: NewBillingService(fmt.Sprintf("%s/%s", tmp, models.BillingFilename)),
	}
}
