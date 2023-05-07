package models

import (
	"bufio"
	"io"
	"os"
	"skillbox-diploma/internal/utils"
	"strconv"
	"strings"
)

type Voice struct {
	Country        string  `json:"country"`
	Load           string  `json:"bandwidth"`
	ResponseTime   string  `json:"response_time"`
	Provider       string  `json:"provider"`
	Stability      float32 `json:"connection_stability"`
	TTFB           int     `json:"ttfb"`
	Purity         int     `json:"voice_purity"`
	MedianDuration int     `json:"median_of_call_time"`
}

// Индексы столбцов из csv файла с Voice.
const (
	countryVoice = iota
	loadVoice
	responseTimeVoice
	providerVoice
	stabilityVoice
	ttfbVoice
	purityVoice
	medianDurationVoice
)

// VoiceFilename название csv файла с данными о Voice.
const VoiceFilename = "voice.csv"

// GetVoice возвращает список Voice из csv файла VoiceFilename
func GetVoice(path string) ([]Voice, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(string(content))
	scanner := bufio.NewScanner(reader)
	voiceList := make([]Voice, 0)

	for scanner.Scan() {
		line := scanner.Text()
		voice, ok := parseVoiceData(line)

		if ok {
			voiceList = append(voiceList, voice)
		}
	}

	return voiceList, nil
}

// parseVoiceData парсит строку из csv-файла. Проверяет данные на валидность.
func parseVoiceData(line string) (Voice, bool) {
	voice := strings.Split(line, ";")

	switch {
	case len(voice) != 8:
		fallthrough
	case !utils.IsValidCountry(voice[countryVoice]):
		fallthrough
	case !utils.IsValidLoad(voice[loadVoice]):
		fallthrough
	case !utils.IsValidResponseTime(voice[responseTimeVoice]):
		fallthrough
	case !utils.IsValidVoiceProvider(voice[providerVoice]):
		fallthrough
	case !utils.IsValidStability(voice[stabilityVoice]):
		fallthrough
	case !utils.IsValidPurity(voice[purityVoice]):
		fallthrough
	case !utils.IsValidTTFB(voice[ttfbVoice]):
		fallthrough
	case !utils.IsMedianDuration(voice[medianDurationVoice]):
		return Voice{}, false
	}

	load := voice[loadVoice]
	responseTime := voice[responseTimeVoice]
	stability64, _ := strconv.ParseFloat(voice[stabilityVoice], 32)
	ttfb, _ := strconv.Atoi(voice[responseTimeVoice])
	purity, _ := strconv.Atoi(voice[purityVoice])
	medianDuration, _ := strconv.Atoi(voice[medianDurationVoice])

	return Voice{
		Country:        voice[countryVoice],
		Load:           load,
		ResponseTime:   responseTime,
		Provider:       voice[providerVoice],
		Stability:      float32(stability64),
		TTFB:           ttfb,
		Purity:         purity,
		MedianDuration: medianDuration,
	}, true
}
