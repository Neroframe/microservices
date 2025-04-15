package utils

func DefaultIfNilString(val *string, fallback string) string {
	if val != nil {
		return *val
	}
	return fallback
}

func DefaultIfNilFloat(val *float64, fallback float64) float64 {
	if val != nil {
		return *val
	}
	return fallback
}

func DefaultIfNilInt(val *int, fallback int) int {
	if val != nil {
		return *val
	}
	return fallback
}
