package math

import stdmath "math"

func round(num float64) int {
	return int(num + stdmath.Copysign(0.5, num))
}

func Fixed(num float64, precision int) float64 {
	output := stdmath.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
