package service

import (
	"fmt"
	"skillbox-diploma/internal/models"
)

type VoiceService struct {
	path string
}

func (s VoiceService) Get() ([]models.Voice, error) {
	voiceList, err := models.GetVoice(s.path)
	if err != nil {
		return nil, fmt.Errorf("getVoice failed: %w", err)
	}

	return voiceList, nil
}

func NewVoiceService(path string) VoiceService {
	return VoiceService{
		path: path,
	}
}
