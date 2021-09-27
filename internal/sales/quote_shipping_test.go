package sales_test

import (
	"io/ioutil"
	"log"
	"testing"

	"bitbucket.org/josecaceresatencora/logistics/internal/sales"
	"bitbucket.org/josecaceresatencora/logistics/pkg/geo"
	"bitbucket.org/josecaceresatencora/logistics/pkg/math"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestQuoteShipping(t *testing.T) {
	log.SetOutput(ioutil.Discard)

	want := 12.75
	brief, err := sales.QuoteShipping(context.Background(), sales.Shipping{
		From: geo.Location{
			Lat: 0,
			Lng: 0,
		},
		To: geo.Location{
			Lat: 1000,
			Lng: 0,
		},
		Dimension: sales.Dimension{
			Width:  5,
			Height: 5,
			Length: 5,
		},
		Weight: sales.Weight(7.5),
	})

	got := math.Fixed(float64(brief.Total()), 2)

	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
