package transport

import "bitbucket.org/josecaceresatencora/logistics/pkg/geo"

type (
	City struct {
		name     string
		location geo.Location
	}
)

func (c City) Name() string {
	return c.name
}

func (c City) Location() geo.Location {
	return c.location
}
