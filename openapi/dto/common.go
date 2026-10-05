package openapidto

type (
	Amount struct {
		Value    string `json:"value" validate:"required"`
		Currency string `json:"currency" validate:"required"`
	}
)

const (
	CurrencyCodeIDR = "IDR"
)

var ValidCurrencyCodes = map[string]bool{
	CurrencyCodeIDR: true,
}
