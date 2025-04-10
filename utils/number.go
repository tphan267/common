package utils

import "strconv"

func ToInt(value any) int {
	return int(ToFloat64(value))
}

func ToInt64(value any) int64 {
	return int64(ToFloat64(value))
}

func ToFloat64(value any) float64 {
	switch v := value.(type) {
	case int:
		return float64(v)
	case float64:
		return v
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return 0
}
