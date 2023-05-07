package service

import (
	"fmt"
	"skillbox-diploma/internal/models"
	"skillbox-diploma/internal/utils"
)

type EmailService struct {
	path string
}

func (s EmailService) Get() (models.EmailCollection, error) {
	emails, err := models.GetEmail(s.path)
	if err != nil {
		return nil, fmt.Errorf("failed to get emails ")
	}

	resultEmail := make(models.EmailCollection, 0)

	for code := range utils.CountryCode {
		slowProviders, fastProviders := models.GetEmailProvider(emails, code)
		if len(slowProviders) > 2 && len(fastProviders) > 2 {
			resultEmail[code] = [][]models.Email{slowProviders, fastProviders}
		}
	}

	return resultEmail, nil
}

func NewEmailService(path string) EmailService {
	return EmailService{
		path: path,
	}
}
