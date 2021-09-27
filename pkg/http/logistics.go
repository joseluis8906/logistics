package http

import (
	"context"
	"net/http"
	"strconv"

	"bitbucket.org/josecaceresatencora/logistics/internal/sales"
	"bitbucket.org/josecaceresatencora/logistics/pkg/geo"
	"bitbucket.org/josecaceresatencora/logistics/pkg/logistics"
	"bitbucket.org/josecaceresatencora/logistics/pkg/math"
	log "github.com/sirupsen/logrus"
)

type (
	QuoteShippingResponse struct {
		Distance float64 `json:"distance_amount"`
		Weight   float64 `json:"weight_amount"`
		Size     float64 `json:"size_amount"`
		Total    float64 `json:"total"`
	}
)

func QuoteShipping(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fromLat, err := strconv.ParseFloat(r.URL.Query().Get("from_lat"), 64)
	if err != nil {
		log.Error(err)
	}

	fromLng, err := strconv.ParseFloat(r.URL.Query().Get("from_lng"), 64)
	if err != nil {
		log.Error(err)
	}

	toLat, err := strconv.ParseFloat(r.URL.Query().Get("to_lat"), 64)
	if err != nil {
		log.Error(err)
	}

	toLng, err := strconv.ParseFloat(r.URL.Query().Get("to_lng"), 64)
	if err != nil {
		log.Error(err)
	}

	weigth, err := strconv.ParseFloat(r.URL.Query().Get("weigth"), 64)
	if err != nil {
		log.Error(err)
	}

	width, err := strconv.ParseFloat(r.URL.Query().Get("width"), 64)
	if err != nil {
		log.Error(err)
	}

	height, err := strconv.ParseFloat(r.URL.Query().Get("height"), 64)
	if err != nil {
		log.Error(err)
	}

	length, err := strconv.ParseFloat(r.URL.Query().Get("length"), 64)
	if err != nil {
		log.Error(err)
	}

	shipping := sales.Shipping{
		From: geo.Location{
			Lat: fromLat,
			Lng: fromLng,
		},
		To: geo.Location{
			Lat: toLat,
			Lng: toLng,
		},
		Weight: sales.Weight(weigth),
		Dimension: sales.Dimension{
			Width:  width,
			Height: height,
			Length: length,
		},
	}

	brief, err := logistics.App().Queries.QuoteShipping(ctx, shipping)
	if err != nil {
		Error(w, r, http.StatusInternalServerError, err)
	}

	Success(w, r, http.StatusOK, QuoteShippingResponse{
		Distance: math.Fixed(float64(brief.Distance), 2),
		Weight:   math.Fixed(float64(brief.Weight), 2),
		Size:     math.Fixed(float64(brief.Size), 2),
		Total:    math.Fixed(float64(brief.Total()), 2),
	})
}
