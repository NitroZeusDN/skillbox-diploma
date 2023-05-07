package models

import (
	"bufio"
	"io/ioutil"
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
	COUNTRY_VOICE = iota
	LOAD_VOICE
	RESPONSE_TIME_VOICE
	PROVIDER_VOICE
	STABILITY_VOICE
	TTFB_VOICE
	PURITY_VOICE
	MEDIAN_DURATION_VOICE
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

	content, err := ioutil.ReadAll(file)
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
	case !utils.IsValidCountry(voice[COUNTRY_VOICE]):
		fallthrough
	case !utils.IsValidLoad(voice[LOAD_VOICE]):
		fallthrough
	case !utils.IsValidResponseTime(voice[RESPONSE_TIME_VOICE]):
		fallthrough
	case !utils.IsValidVoiceProvider(voice[PROVIDER_VOICE]):
		fallthrough
	case !utils.IsValidStability(voice[STABILITY_VOICE]):
		fallthrough
	case !utils.IsValidPurity(voice[PURITY_VOICE]):
		fallthrough
	case !utils.IsValidTTFB(voice[TTFB_VOICE]):
		fallthrough
	case !utils.IsMedianDuration(voice[MEDIAN_DURATION_VOICE]):
		return Voice{}, false
	}

	load := voice[LOAD_VOICE]
	responseTime := voice[RESPONSE_TIME_VOICE]
	stability64, _ := strconv.ParseFloat(voice[STABILITY_VOICE], 32)
	ttfb, _ := strconv.Atoi(voice[RESPONSE_TIME_VOICE])
	purity, _ := strconv.Atoi(voice[PURITY_VOICE])
	medianDuration, _ := strconv.Atoi(voice[MEDIAN_DURATION_VOICE])

	return Voice{
		Country:        voice[COUNTRY_VOICE],
		Load:           load,
		ResponseTime:   responseTime,
		Provider:       voice[PROVIDER_VOICE],
		Stability:      float32(stability64),
		TTFB:           ttfb,
		Purity:         purity,
		MedianDuration: medianDuration,
	}, true
}
