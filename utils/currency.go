package utils

const (
	USD = "USD"
	UZS = "UZS"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case EUR, UZS, USD:
		return true
	default:
		return false

	}
}
