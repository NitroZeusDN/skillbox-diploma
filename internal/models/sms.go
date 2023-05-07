package models

import (
	"bufio"
	"io"
	"os"
	"skillbox-diploma/internal/utils"
	"strings"
)

// SMS данные о смс.
type SMS struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

// Индексы столбцов SMS в csv файле.
const (
	countrySMS = iota
	bandwidthSMS
	responseTimeSMS
	providerSMS
)

// SMSFilename название файла с данными SMS.
const SMSFilename = "sms.csv"

// GetSMS возвращает список данных о СМС из csv файла.
func GetSMS(path string) ([]SMS, error) {
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
	smsList := make([]SMS, 0)

	for scanner.Scan() {
		line := scanner.Text()
		sms, ok := parseSMS(line)
		if ok {
			smsList = append(smsList, sms)
		}
	}

	return smsList, nil
}

// parseSMS парсит CSV файл построчно и возвращает данные о SMS из строки.
// При этом проверяет корректное ли это СМС.
func parseSMS(line string) (SMS, bool) {
	sms := strings.Split(line, ";")

	switch {
	case len(sms) < 4:
		fallthrough
	case !utils.IsValidCountry(sms[countrySMS]):
		fallthrough
	case !utils.IsValidBandwidth(sms[bandwidthSMS]):
		fallthrough
	case !utils.IsValidResponseTime(sms[responseTimeSMS]):
		fallthrough
	case !utils.IsValidSMSProvider(sms[providerSMS]):
		return SMS{}, false
	}

	return SMS{
		Country:      sms[countrySMS],
		Bandwidth:    sms[bandwidthSMS],
		ResponseTime: sms[responseTimeSMS],
		Provider:     sms[providerSMS],
	}, true
}
