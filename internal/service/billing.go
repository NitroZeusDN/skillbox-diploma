package service

import "skillbox-diploma/internal/models"

type BillingService struct {
	path string
}

func (s BillingService) Get() (models.Billing, error) {
	billing, err := models.GetBilling(s.path)
	if err != nil {
		return models.Billing{}, err
	}

	return billing, nil
}

func NewBillingService(path string) BillingService {
	return BillingService{
		path: path,
	}
}
