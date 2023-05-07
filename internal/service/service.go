package service

import (
	"fmt"
	"skillbox-diploma/internal/models"
)

type Service struct {
	tmp string

	sms      SMSService
	mms      MMSService
	billing  BillingService
	email    EmailService
	voice    VoiceService
	incident IncidentService
}

func (s *Service) Get() (models.ResultSetT, error) {
	var (
		res models.ResultSetT
		err error
	)

	res.SMS, err = s.sms.Get()
	if err != nil {
		return models.ResultSetT{}, fmt.Errorf("failed to get sms data: %w", err)
	}

	res.MMS, err = s.mms.Get()
	if err != nil {
		return models.ResultSetT{}, fmt.Errorf("failed to get mms data: %w", err)
	}

	res.Billing, err = s.billing.Get()
	if err != nil {
		return models.ResultSetT{}, fmt.Errorf("failed to get billing data: %w", err)
	}

	res.Email, err = s.email.Get()
	if err != nil {
		return models.ResultSetT{}, fmt.Errorf("failed to get email data: %w", err)
	}

	res.VoiceCall, err = s.voice.Get()
	if err != nil {
		return models.ResultSetT{}, fmt.Errorf("failed to get voice call data: %w", err)
	}

	res.Incidents, err = s.incident.Get()
	if err != nil {
		return models.ResultSetT{}, fmt.Errorf("failed to get incidents data: %w", err)
	}

	return res, nil
}

func New(tmp, host string) *Service {
	return &Service{
		tmp:      tmp,
		sms:      NewSMSService(fmt.Sprintf("%s/%s", tmp, models.SMSFilename)),
		mms:      NewMMSService(host),
		billing:  NewBillingService(fmt.Sprintf("%s/%s", tmp, models.BillingFilename)),
		email:    NewEmailService(fmt.Sprintf("%s/%s", tmp, models.EmailFilename)),
		voice:    NewVoiceService(fmt.Sprintf("%s/%s", tmp, models.VoiceFilename)),
		incident: NewIncidentService(host),
	}
}
