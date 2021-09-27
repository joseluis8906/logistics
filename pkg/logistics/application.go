package logistics

import (
	"context"

	"bitbucket.org/josecaceresatencora/logistics/internal/sales"
)

var (
	app = Application{
		Commands: Commands{},
		Queries:  Queries{},
	}
)

func SetApp(anApp Application) {
	app = anApp
}

func App() *Application {
	return &app
}

type (
	Application struct {
		Commands Commands
		Queries  Queries
	}

	Commands struct {
		PlaceShipping   func(context.Context) error
		RouteShipping   func(context.Context) error
		DeliverShipping func(context.Context) error
		ReverseShipping func(context.Context) error
	}

	Queries struct {
		QuoteShipping func(context.Context, sales.Shipping) (sales.QuoteBrief, error)
		TrackShipping func(context.Context) error
	}
)
