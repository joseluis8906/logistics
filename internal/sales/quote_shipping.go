package sales

import (
	"context"
	"time"

	"bitbucket.org/josecaceresatencora/logistics/pkg/bus"
	"bitbucket.org/josecaceresatencora/logistics/pkg/geo"
	"bitbucket.org/josecaceresatencora/logistics/pkg/math"
	log "github.com/sirupsen/logrus"
)

type (
	Weight float64

	Shipping struct {
		From      geo.Location
		To        geo.Location
		When      time.Time
		Dimension Dimension
		Weight    Weight
	}

	ShippingQuoted struct {
		Distance float64 `json:"distance_fee"`
		Weight   float64 `json:"weight_fee"`
		Size     float64 `json:"volume_fee"`
		Total    float64 `json:"total_fee"`
	}
)

func QuoteShipping(_ context.Context, shipping Shipping) (QuoteBrief, error) {
	const (
		AmountByKm        = Currency(0.01)
		AmountByLbr       = Currency(0.2)
		AmountByCubicInch = Currency(0.01)
	)

	brief := QuoteBrief{
		Distance: Currency(geo.CalculateDistance(shipping.From, shipping.To)) * AmountByKm,
		Weight:   Currency(shipping.Weight) * AmountByLbr,
		Size:     Currency(shipping.Dimension.Volume()) * AmountByCubicInch,
	}

	err := bus.Emit("shipping.quoted", ShippingQuoted{
		Distance: math.Fixed(float64(brief.Distance), 2),
		Weight:   math.Fixed(float64(brief.Weight), 2),
		Size:     math.Fixed(float64(brief.Size), 2),
		Total:    math.Fixed(float64(brief.Total()), 2),
	})
	if err != nil {
		log.Error(err)
	}

	return brief, nil
}
