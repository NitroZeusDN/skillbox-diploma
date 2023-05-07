package service

import (
	"fmt"
	"skillbox-diploma/internal/models"
)

type Service struct {
	tmp string

	billing BillingService
	sms     SMSService
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

	res.SMS, err = s.sms.Get()
	if err != nil {
		return models.ResultSetT{}, fmt.Errorf("failed to get sms data: %w", err)
	}

	return res, nil
}

func New(tmp string) *Service {
	return &Service{
		tmp:     tmp,
		billing: NewBillingService(fmt.Sprintf("%s/%s", tmp, models.BillingFilename)),
		sms:     NewSMSService(fmt.Sprintf("%s/%s", tmp, models.SMSFilename)),
	}
}
