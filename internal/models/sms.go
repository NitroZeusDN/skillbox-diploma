package models

import (
	"bufio"
	"io/ioutil"
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
	COUNTRY_SMS = iota
	BANDWIDTH_SMS
	RESPONSE_TIME_SMS
	PROVIDER_SMS
)

const SMSFilename = "sms.csv"

// GetSMS возвращает список данных о СМС из csv файла.
func GetSMS(path string) ([]SMS, error) {
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
	case !utils.IsValidCountry(sms[COUNTRY_SMS]):
		fallthrough
	case !utils.IsValidBandwidth(sms[BANDWIDTH_SMS]):
		fallthrough
	case !utils.IsValidResponseTime(sms[RESPONSE_TIME_SMS]):
		fallthrough
	case !utils.IsValidSMSProvider(sms[PROVIDER_SMS]):
		return SMS{}, false
	}

	return SMS{
		Country:      sms[COUNTRY_SMS],
		Bandwidth:    sms[BANDWIDTH_SMS],
		ResponseTime: sms[RESPONSE_TIME_SMS],
		Provider:     sms[PROVIDER_SMS],
	}, true
}
