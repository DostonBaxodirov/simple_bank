package utils

const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
	AED = "AED"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, AED, CAD, EUR:
		return true
	}

	return false
}
