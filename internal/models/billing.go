package models

import (
	"io"
	"math"
	"os"
)

// Индексы для данных из биллинга.
const (
	billingCreateCustomer = iota
	billingPurchase
	billingPayout
	billingRecurring
	billingFraudControl
	billingCheckoutPage
)

// BillingFilename название файла с данными по биллингу.
const BillingFilename = "billing.txt"

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

	content, err := io.ReadAll(file)
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
		CreateCustomer: mask>>billingCreateCustomer&1 == 1,
		Purchase:       mask>>billingPurchase&1 == 1,
		Payout:         mask>>billingPayout&1 == 1,
		Recurring:      mask>>billingRecurring&1 == 1,
		FraudControl:   mask>>billingFraudControl&1 == 1,
		CheckoutPage:   mask>>billingCheckoutPage&1 == 1,
	}
}
