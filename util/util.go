package util

func Clamp(num, min, max float64) float64 {
	if num < min {
		return min
	}

	if num > max {
		return max
	}

	return num
}
