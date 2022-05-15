package roundfloat

import "math"

func RoundFloat(price float64) float64 {
	return math.Round(float64(price*100)) / 100
}
