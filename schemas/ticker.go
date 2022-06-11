package schemas

import "github.com/graphql-go/graphql"

var TickerSchema = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Ticker",
		Fields: graphql.Fields{
			"high": &graphql.Field{
				Type: graphql.String,
			},
			"low": &graphql.Field{
				Type: graphql.String,
			},
			"vol": &graphql.Field{
				Type: graphql.String,
			},
			"last": &graphql.Field{
				Type: graphql.String,
			},
			"buy": &graphql.Field{
				Type: graphql.String,
			},
			"sell": &graphql.Field{
				Type: graphql.String,
			},
			"date": &graphql.Field{
				Type: graphql.Int,
			},
			"open": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
