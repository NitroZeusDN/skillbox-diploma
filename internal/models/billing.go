package models

import (
	"io/ioutil"
	"math"
	"os"
)

// Индексы для данных из биллинга.
const (
	BILLING_CREATE_CUSTOMER = iota
	BILLING_PURCHASE
	BILLING_PAYOUT
	BILLING_RECURRING
	BILLING_FRAUD_CONTROL
	BILLING_CHECKOUT_PAGE
)

// BillingFilename название файла с данными по биллингу.
const BillingFilename = "billing.csv"

// Billing данные по биллингу.
type Billing struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

// GetBilling возвращает список данных по биллингу из csv файла.
func GetBilling(path string) (Billing, error) {
	file, err := os.Open(path)
	if err != nil {
		return Billing{}, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return Billing{}, err
	}

	return parseBilling(content), nil

}

// parseBilling обрабатывает байты файла и возвращает готовую структуру Billing.
func parseBilling(data []byte) Billing {
	if len(data) != 6 {
		return Billing{}
	}

	var mask int8
	for i, bit := range data {
		if bit == '1' {
			mask += int8(math.Pow(2, float64(len(data)-i-1)))
		}
	}

	return Billing{
		CreateCustomer: mask>>BILLING_CREATE_CUSTOMER&1 == 1,
		Purchase:       mask>>BILLING_PURCHASE&1 == 1,
		Payout:         mask>>BILLING_PAYOUT&1 == 1,
		Recurring:      mask>>BILLING_RECURRING&1 == 1,
		FraudControl:   mask>>BILLING_FRAUD_CONTROL&1 == 1,
		CheckoutPage:   mask>>BILLING_CHECKOUT_PAGE&1 == 1,
	}
}
