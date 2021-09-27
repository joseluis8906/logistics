package geo

import (
	"math"
)

func CalculateDistance(x, y Location) float64 {
	return math.Sqrt(math.Pow(y.Lat-x.Lat, 2.0) + math.Pow(y.Lng-x.Lng, 2.0))
}
