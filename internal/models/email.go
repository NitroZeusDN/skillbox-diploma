package models

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"skillbox-diploma/internal/utils"
	"sort"
	"strconv"
	"strings"
)

// Email данные по email.
type Email struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type EmailCollection = map[string][][]Email

// Индексы столбцов в csv файле
const (
	COUNTRY_EMAIL = iota
	PROVIDER_EMAIL
	DELIVERY_TIME_EMAIL
)

// EmailFilename название файла с данными Email.
const EmailFilename = "email.csv"

// GetEmail возвращает список Email из csv файла EmailFilename.
func GetEmail(path string) ([]Email, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	reader := strings.NewReader(string(content))
	scanner := bufio.NewScanner(reader)
	emailList := make([]Email, 0)

	for scanner.Scan() {
		line := scanner.Text()
		email, ok := parseEmail(line)

		if ok {
			emailList = append(emailList, email)
		}
	}

	return emailList, nil
}

// GetEmailProvider сортирует всех поставщиков услуг в стране в соответствии со средним временем доставки писем.
// Возвращает два фрагмента: первый содержит трех самых быстрых провайдеров, второй - трех самых медленных.
func GetEmailProvider(data []Email, code string) (slow []Email, fast []Email) {
	emailsByCountry := make([]Email, 0)
	for _, email := range data {
		if email.Country == code {
			emailsByCountry = append(emailsByCountry, email)
		}
	}

	sort.SliceStable(emailsByCountry, func(i, j int) bool {
		return emailsByCountry[i].DeliveryTime < emailsByCountry[j].DeliveryTime
	})

	if len(emailsByCountry) < 3 {
		return emailsByCountry, emailsByCountry
	}

	return emailsByCountry[len(emailsByCountry)-3:], emailsByCountry[:3]
}

// parseEmail парсит CSV файл построчно и возвращает данные о Email из строки.
// При этом проверяет корректная ли это информация.
func parseEmail(line string) (Email, bool) {
	email := strings.Split(line, ";")

	switch {
	case len(email) != 3:
		fallthrough
	case !utils.IsValidCountry(email[COUNTRY_EMAIL]):
		fallthrough
	case !utils.IsValidEmailProvider(email[PROVIDER_EMAIL]):
		fallthrough
	case !utils.IsValidDeliveryTime(email[DELIVERY_TIME_EMAIL]):
		return Email{}, false
	}

	deliveryTime, _ := strconv.Atoi(email[DELIVERY_TIME_EMAIL])

	return Email{
		Country:      email[COUNTRY_EMAIL],
		Provider:     email[PROVIDER_EMAIL],
		DeliveryTime: deliveryTime,
	}, true
}
