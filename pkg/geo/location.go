package geo

type (
	Location struct {
		Lat float64
		Lng float64
	}
)

func NullLocation() Location {
	return Location{}
}
