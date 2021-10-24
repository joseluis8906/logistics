package transport

import (
	"context"

	"bitbucket.org/josecaceresatencora/logistics/pkg/geo"
)

var (
	cities = []City{
		{name: "Bogota", location: geo.Location{Lat: 0.0, Lng: 50.0}},
		{name: "Buenos Aires", location: geo.Location{Lat: 25.0, Lng: 25.0}},
		{name: "San Jose", location: geo.Location{Lat: 50.0, Lng: 100.0}},
		{name: "Lima", location: geo.Location{Lat: 100.0, Lng: 0.0}},
	}
)

func CityLocation(ctx context.Context, aCity City) geo.Location {
	for _, city := range cities {
		if city.Name() == aCity.Name() {
			return city.Location()
		}
	}

	return geo.NullLocation()
}
