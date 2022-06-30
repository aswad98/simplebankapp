package util

const (
	USD   = "USD"
	EUR   = "EUR"
	RIYAL = "RIYAL"
	CAD   = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RIYAL, CAD:
		return true
	}
	return false
}
